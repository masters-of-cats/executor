// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/executor/api"
)

type FakeClient struct {
	PingStub        func() error
	pingMutex       sync.RWMutex
	pingArgsForCall []struct{}
	pingReturns struct {
		result1 error
	}
	AllocateContainerStub        func(allocationGuid string, request api.ContainerAllocationRequest) (api.Container, error)
	allocateContainerMutex       sync.RWMutex
	allocateContainerArgsForCall []struct {
		allocationGuid string
		request        api.ContainerAllocationRequest
	}
	allocateContainerReturns struct {
		result1 api.Container
		result2 error
	}
	GetContainerStub        func(allocationGuid string) (api.Container, error)
	getContainerMutex       sync.RWMutex
	getContainerArgsForCall []struct {
		allocationGuid string
	}
	getContainerReturns struct {
		result1 api.Container
		result2 error
	}
	InitializeContainerStub        func(allocationGuid string, request api.ContainerInitializationRequest) (api.Container, error)
	initializeContainerMutex       sync.RWMutex
	initializeContainerArgsForCall []struct {
		allocationGuid string
		request        api.ContainerInitializationRequest
	}
	initializeContainerReturns struct {
		result1 api.Container
		result2 error
	}
	RunStub        func(allocationGuid string, request api.ContainerRunRequest) error
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		allocationGuid string
		request        api.ContainerRunRequest
	}
	runReturns struct {
		result1 error
	}
	DeleteContainerStub        func(allocationGuid string) error
	deleteContainerMutex       sync.RWMutex
	deleteContainerArgsForCall []struct {
		allocationGuid string
	}
	deleteContainerReturns struct {
		result1 error
	}
	ListContainersStub        func() ([]api.Container, error)
	listContainersMutex       sync.RWMutex
	listContainersArgsForCall []struct{}
	listContainersReturns struct {
		result1 []api.Container
		result2 error
	}
	RemainingResourcesStub        func() (api.ExecutorResources, error)
	remainingResourcesMutex       sync.RWMutex
	remainingResourcesArgsForCall []struct{}
	remainingResourcesReturns struct {
		result1 api.ExecutorResources
		result2 error
	}
	TotalResourcesStub        func() (api.ExecutorResources, error)
	totalResourcesMutex       sync.RWMutex
	totalResourcesArgsForCall []struct{}
	totalResourcesReturns struct {
		result1 api.ExecutorResources
		result2 error
	}
}

func (fake *FakeClient) Ping() error {
	fake.pingMutex.Lock()
	fake.pingArgsForCall = append(fake.pingArgsForCall, struct{}{})
	fake.pingMutex.Unlock()
	if fake.PingStub != nil {
		return fake.PingStub()
	} else {
		return fake.pingReturns.result1
	}
}

func (fake *FakeClient) PingCallCount() int {
	fake.pingMutex.RLock()
	defer fake.pingMutex.RUnlock()
	return len(fake.pingArgsForCall)
}

func (fake *FakeClient) PingReturns(result1 error) {
	fake.PingStub = nil
	fake.pingReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) AllocateContainer(allocationGuid string, request api.ContainerAllocationRequest) (api.Container, error) {
	fake.allocateContainerMutex.Lock()
	fake.allocateContainerArgsForCall = append(fake.allocateContainerArgsForCall, struct {
		allocationGuid string
		request        api.ContainerAllocationRequest
	}{allocationGuid, request})
	fake.allocateContainerMutex.Unlock()
	if fake.AllocateContainerStub != nil {
		return fake.AllocateContainerStub(allocationGuid, request)
	} else {
		return fake.allocateContainerReturns.result1, fake.allocateContainerReturns.result2
	}
}

func (fake *FakeClient) AllocateContainerCallCount() int {
	fake.allocateContainerMutex.RLock()
	defer fake.allocateContainerMutex.RUnlock()
	return len(fake.allocateContainerArgsForCall)
}

func (fake *FakeClient) AllocateContainerArgsForCall(i int) (string, api.ContainerAllocationRequest) {
	fake.allocateContainerMutex.RLock()
	defer fake.allocateContainerMutex.RUnlock()
	return fake.allocateContainerArgsForCall[i].allocationGuid, fake.allocateContainerArgsForCall[i].request
}

func (fake *FakeClient) AllocateContainerReturns(result1 api.Container, result2 error) {
	fake.AllocateContainerStub = nil
	fake.allocateContainerReturns = struct {
		result1 api.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetContainer(allocationGuid string) (api.Container, error) {
	fake.getContainerMutex.Lock()
	fake.getContainerArgsForCall = append(fake.getContainerArgsForCall, struct {
		allocationGuid string
	}{allocationGuid})
	fake.getContainerMutex.Unlock()
	if fake.GetContainerStub != nil {
		return fake.GetContainerStub(allocationGuid)
	} else {
		return fake.getContainerReturns.result1, fake.getContainerReturns.result2
	}
}

func (fake *FakeClient) GetContainerCallCount() int {
	fake.getContainerMutex.RLock()
	defer fake.getContainerMutex.RUnlock()
	return len(fake.getContainerArgsForCall)
}

func (fake *FakeClient) GetContainerArgsForCall(i int) string {
	fake.getContainerMutex.RLock()
	defer fake.getContainerMutex.RUnlock()
	return fake.getContainerArgsForCall[i].allocationGuid
}

func (fake *FakeClient) GetContainerReturns(result1 api.Container, result2 error) {
	fake.GetContainerStub = nil
	fake.getContainerReturns = struct {
		result1 api.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) InitializeContainer(allocationGuid string, request api.ContainerInitializationRequest) (api.Container, error) {
	fake.initializeContainerMutex.Lock()
	fake.initializeContainerArgsForCall = append(fake.initializeContainerArgsForCall, struct {
		allocationGuid string
		request        api.ContainerInitializationRequest
	}{allocationGuid, request})
	fake.initializeContainerMutex.Unlock()
	if fake.InitializeContainerStub != nil {
		return fake.InitializeContainerStub(allocationGuid, request)
	} else {
		return fake.initializeContainerReturns.result1, fake.initializeContainerReturns.result2
	}
}

func (fake *FakeClient) InitializeContainerCallCount() int {
	fake.initializeContainerMutex.RLock()
	defer fake.initializeContainerMutex.RUnlock()
	return len(fake.initializeContainerArgsForCall)
}

func (fake *FakeClient) InitializeContainerArgsForCall(i int) (string, api.ContainerInitializationRequest) {
	fake.initializeContainerMutex.RLock()
	defer fake.initializeContainerMutex.RUnlock()
	return fake.initializeContainerArgsForCall[i].allocationGuid, fake.initializeContainerArgsForCall[i].request
}

func (fake *FakeClient) InitializeContainerReturns(result1 api.Container, result2 error) {
	fake.InitializeContainerStub = nil
	fake.initializeContainerReturns = struct {
		result1 api.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Run(allocationGuid string, request api.ContainerRunRequest) error {
	fake.runMutex.Lock()
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		allocationGuid string
		request        api.ContainerRunRequest
	}{allocationGuid, request})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub(allocationGuid, request)
	} else {
		return fake.runReturns.result1
	}
}

func (fake *FakeClient) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeClient) RunArgsForCall(i int) (string, api.ContainerRunRequest) {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return fake.runArgsForCall[i].allocationGuid, fake.runArgsForCall[i].request
}

func (fake *FakeClient) RunReturns(result1 error) {
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) DeleteContainer(allocationGuid string) error {
	fake.deleteContainerMutex.Lock()
	fake.deleteContainerArgsForCall = append(fake.deleteContainerArgsForCall, struct {
		allocationGuid string
	}{allocationGuid})
	fake.deleteContainerMutex.Unlock()
	if fake.DeleteContainerStub != nil {
		return fake.DeleteContainerStub(allocationGuid)
	} else {
		return fake.deleteContainerReturns.result1
	}
}

func (fake *FakeClient) DeleteContainerCallCount() int {
	fake.deleteContainerMutex.RLock()
	defer fake.deleteContainerMutex.RUnlock()
	return len(fake.deleteContainerArgsForCall)
}

func (fake *FakeClient) DeleteContainerArgsForCall(i int) string {
	fake.deleteContainerMutex.RLock()
	defer fake.deleteContainerMutex.RUnlock()
	return fake.deleteContainerArgsForCall[i].allocationGuid
}

func (fake *FakeClient) DeleteContainerReturns(result1 error) {
	fake.DeleteContainerStub = nil
	fake.deleteContainerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) ListContainers() ([]api.Container, error) {
	fake.listContainersMutex.Lock()
	fake.listContainersArgsForCall = append(fake.listContainersArgsForCall, struct{}{})
	fake.listContainersMutex.Unlock()
	if fake.ListContainersStub != nil {
		return fake.ListContainersStub()
	} else {
		return fake.listContainersReturns.result1, fake.listContainersReturns.result2
	}
}

func (fake *FakeClient) ListContainersCallCount() int {
	fake.listContainersMutex.RLock()
	defer fake.listContainersMutex.RUnlock()
	return len(fake.listContainersArgsForCall)
}

func (fake *FakeClient) ListContainersReturns(result1 []api.Container, result2 error) {
	fake.ListContainersStub = nil
	fake.listContainersReturns = struct {
		result1 []api.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) RemainingResources() (api.ExecutorResources, error) {
	fake.remainingResourcesMutex.Lock()
	fake.remainingResourcesArgsForCall = append(fake.remainingResourcesArgsForCall, struct{}{})
	fake.remainingResourcesMutex.Unlock()
	if fake.RemainingResourcesStub != nil {
		return fake.RemainingResourcesStub()
	} else {
		return fake.remainingResourcesReturns.result1, fake.remainingResourcesReturns.result2
	}
}

func (fake *FakeClient) RemainingResourcesCallCount() int {
	fake.remainingResourcesMutex.RLock()
	defer fake.remainingResourcesMutex.RUnlock()
	return len(fake.remainingResourcesArgsForCall)
}

func (fake *FakeClient) RemainingResourcesReturns(result1 api.ExecutorResources, result2 error) {
	fake.RemainingResourcesStub = nil
	fake.remainingResourcesReturns = struct {
		result1 api.ExecutorResources
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) TotalResources() (api.ExecutorResources, error) {
	fake.totalResourcesMutex.Lock()
	fake.totalResourcesArgsForCall = append(fake.totalResourcesArgsForCall, struct{}{})
	fake.totalResourcesMutex.Unlock()
	if fake.TotalResourcesStub != nil {
		return fake.TotalResourcesStub()
	} else {
		return fake.totalResourcesReturns.result1, fake.totalResourcesReturns.result2
	}
}

func (fake *FakeClient) TotalResourcesCallCount() int {
	fake.totalResourcesMutex.RLock()
	defer fake.totalResourcesMutex.RUnlock()
	return len(fake.totalResourcesArgsForCall)
}

func (fake *FakeClient) TotalResourcesReturns(result1 api.ExecutorResources, result2 error) {
	fake.TotalResourcesStub = nil
	fake.totalResourcesReturns = struct {
		result1 api.ExecutorResources
		result2 error
	}{result1, result2}
}

var _ api.Client = new(FakeClient)