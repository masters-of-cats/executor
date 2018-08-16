package containermetrics

import (
	"os"
	"sync/atomic"
	"syscall"
	"time"

	"code.cloudfoundry.org/clock"
	loggingclient "code.cloudfoundry.org/diego-logging-client"
	"code.cloudfoundry.org/executor"
	"code.cloudfoundry.org/lager"
	"github.com/cloudfoundry/sonde-go/events"
)

var megabytesToBytes int = 1024 * 1024

type StatsReporter struct {
	logger lager.Logger

	interval       time.Duration
	clock          clock.Clock
	executorClient executor.Client
	metrics        atomic.Value

	metronClient          loggingclient.IngressClient
	enableContainerProxy  bool
	proxyMemoryAllocation float64
	setCpuWeight          bool
}

type cpuInfo struct {
	timeSpentInCPU time.Duration
	timeOfSample   time.Time
}

func NewStatsReporter(logger lager.Logger,
	interval time.Duration,
	clock clock.Clock,
	enableContainerProxy bool,
	additionalMemoryMB int,
	executorClient executor.Client,
	metronClient loggingclient.IngressClient,
	setCpuWeight bool,
) *StatsReporter {
	return &StatsReporter{
		logger: logger,

		interval:              interval,
		clock:                 clock,
		executorClient:        executorClient,
		metronClient:          metronClient,
		enableContainerProxy:  enableContainerProxy,
		proxyMemoryAllocation: float64(additionalMemoryMB * megabytesToBytes),
		setCpuWeight:          setCpuWeight,
	}
}

func (reporter *StatsReporter) Run(signals <-chan os.Signal, ready chan<- struct{}) error {
	logger := reporter.logger.Session("container-metrics-reporter")

	ticker := reporter.clock.NewTicker(reporter.interval)
	defer ticker.Stop()

	close(ready)

	cpuInfos := make(map[string]*cpuInfo)
	for {
		select {
		case signal := <-signals:
			logger.Info("signalled", lager.Data{"signal": signal.String()})
			return nil

		case now := <-ticker.C():
			cpuInfos = reporter.emitContainerMetrics(logger, cpuInfos, now)
		}
	}
}

func (reporter *StatsReporter) Metrics() map[string]*CachedContainerMetrics {
	if v := reporter.metrics.Load(); v != nil {
		return v.(map[string]*CachedContainerMetrics)
	}
	return nil
}

func (reporter *StatsReporter) emitContainerMetrics(logger lager.Logger, previousCPUInfos map[string]*cpuInfo, now time.Time) map[string]*cpuInfo {
	logger = logger.Session("tick")

	startTime := reporter.clock.Now()

	logger.Debug("started")
	defer func() {
		logger.Debug("done", lager.Data{
			"took": reporter.clock.Now().Sub(startTime).String(),
		})
	}()

	metrics, err := reporter.executorClient.GetBulkMetrics(logger)
	if err != nil {
		logger.Error("failed-to-get-all-metrics", err)
		return previousCPUInfos
	}

	logger.Debug("emitting", lager.Data{
		"total-containers": len(metrics),
		"get-metrics-took": reporter.clock.Now().Sub(startTime).String(),
	})

	containers, err := reporter.executorClient.ListContainers(logger)
	if err != nil {
		logger.Error("failed-to-fetch-containers", err)
		return previousCPUInfos
	}

	newCPUInfos := make(map[string]*cpuInfo)
	repMetricsMap := make(map[string]*CachedContainerMetrics)

	for _, container := range containers {
		guid := container.Guid
		metric := metrics[guid]

		previousCPUInfo := previousCPUInfos[guid]

		if reporter.enableContainerProxy && container.EnableContainerProxy {
			metric.MemoryUsageInBytes = uint64(float64(metric.MemoryUsageInBytes) * reporter.scaleMemory(container))
			metric.MemoryLimitInBytes = uint64(float64(metric.MemoryLimitInBytes) - reporter.proxyMemoryAllocation)
		}

		repMetrics, cpu := reporter.calculateAndSendMetrics(logger, metric.MetricsConfig, metric.ContainerMetrics, previousCPUInfo, now, container.Resource.MemoryMB)
		if cpu != nil {
			newCPUInfos[guid] = cpu
		}

		if repMetrics != nil {
			repMetricsMap[guid] = repMetrics
		}
	}

	reporter.metrics.Store(repMetricsMap)
	return newCPUInfos
}

func (reporter *StatsReporter) calculateCPUPercentWeighted(cpuPercent float64, containerMemoryMB int) (float64, error) {
	if !reporter.setCpuWeight {
		return cpuPercent, nil
	}

	cellMemoryB, err := sysTotalMemory()
	if err != nil {
		return 0, err
	}
	cellMemoryMB := cellMemoryB / 1024 / 1024
	return cpuPercent * float64(cellMemoryMB) / float64(containerMemoryMB), nil
}

func (reporter *StatsReporter) calculateAndSendMetrics(
	logger lager.Logger,
	metricsConfig executor.MetricsConfig,
	containerMetrics executor.ContainerMetrics,
	previousInfo *cpuInfo,
	now time.Time,
	containerMemoryMB int,
) (*CachedContainerMetrics, *cpuInfo) {
	currentInfo, cpuPercent := calculateInfo(containerMetrics, previousInfo, now)

	cpuPercentWeighted, err := reporter.calculateCPUPercentWeighted(cpuPercent, containerMemoryMB)
	if err != nil {
		logger.Error("get-cell-memory", err)
	}

	if metricsConfig.Guid != "" {
		instanceIndex := int32(metricsConfig.Index)
		err := reporter.metronClient.SendAppMetrics(&events.ContainerMetric{
			ApplicationId:         &metricsConfig.Guid,
			InstanceIndex:         &instanceIndex,
			CpuPercentage:         &cpuPercent,
			CpuPercentageWeighted: &cpuPercentWeighted,
			MemoryBytes:           &containerMetrics.MemoryUsageInBytes,
			DiskBytes:             &containerMetrics.DiskUsageInBytes,
			MemoryBytesQuota:      &containerMetrics.MemoryLimitInBytes,
			DiskBytesQuota:        &containerMetrics.DiskLimitInBytes,
		})

		if err != nil {
			logger.Error("failed-to-send-container-metrics", err, lager.Data{
				"metrics_guid":  metricsConfig.Guid,
				"metrics_index": metricsConfig.Index,
			})
		}
	}

	return &CachedContainerMetrics{
		MetricGUID:               metricsConfig.Guid,
		CPUUsageFraction:         cpuPercent / 100,
		CPUUsageFractionWeighted: cpuPercentWeighted / 100,
		DiskUsageBytes:           containerMetrics.DiskUsageInBytes,
		DiskQuotaBytes:           containerMetrics.DiskLimitInBytes,
		MemoryUsageBytes:         containerMetrics.MemoryUsageInBytes,
		MemoryQuotaBytes:         containerMetrics.MemoryLimitInBytes,
	}, &currentInfo
}

func computeCPUWeightRatio() (float64, error) {
	cellMemoryB, err := sysTotalMemory()
	if err != nil {
		return 0, err
	}

	return 1.0 / (float64(cellMemoryB) / float64(1024*1024)), nil
}

func sysTotalMemory() (uint64, error) {
	in := &syscall.Sysinfo_t{}
	err := syscall.Sysinfo(in)
	if err != nil {
		return 0, err
	}
	// If this is a 32-bit system, then these fields are
	// uint32 instead of uint64.
	// So we always convert to uint64 to match signature.
	return uint64(in.Totalram) * uint64(in.Unit), nil
}

func calculateInfo(containerMetrics executor.ContainerMetrics, previousInfo *cpuInfo, now time.Time) (cpuInfo, float64) {
	currentInfo := cpuInfo{
		timeSpentInCPU: containerMetrics.TimeSpentInCPU,
		timeOfSample:   now,
	}

	var cpuPercent float64
	if previousInfo == nil {
		cpuPercent = 0.0
	} else {
		cpuPercent = computeCPUPercent(
			previousInfo.timeSpentInCPU,
			currentInfo.timeSpentInCPU,
			previousInfo.timeOfSample,
			currentInfo.timeOfSample,
		)
	}
	return currentInfo, cpuPercent
}

// scale from (0 - 100) * runtime.NumCPU()
func computeCPUPercent(timeSpentA, timeSpentB time.Duration, sampleTimeA, sampleTimeB time.Time) float64 {
	// divide change in time spent in CPU over time between samples.
	// result is out of 100 * runtime.NumCPU()
	//
	// don't worry about overflowing int64. it's like, 30 years.
	return float64((timeSpentB-timeSpentA)*100) / float64(sampleTimeB.UnixNano()-sampleTimeA.UnixNano())
}

func (reporter *StatsReporter) scaleMemory(container executor.Container) float64 {
	memFloat := float64(container.MemoryLimit)
	return (memFloat - reporter.proxyMemoryAllocation) / memFloat
}
