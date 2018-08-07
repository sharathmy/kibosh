// Code generated by counterfeiter. DO NOT EDIT.
package brokerfakes

import (
	"sync"

	"github.com/cf-platform-eng/kibosh/pkg/broker"
	my_helm "github.com/cf-platform-eng/kibosh/pkg/helm"
	"github.com/cf-platform-eng/kibosh/pkg/k8s"
)

type FakeHelmClientFactory struct {
	HelmClientStub        func(cluster k8s.Cluster) my_helm.MyHelmClient
	helmClientMutex       sync.RWMutex
	helmClientArgsForCall []struct {
		cluster k8s.Cluster
	}
	helmClientReturns struct {
		result1 my_helm.MyHelmClient
	}
	helmClientReturnsOnCall map[int]struct {
		result1 my_helm.MyHelmClient
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeHelmClientFactory) HelmClient(cluster k8s.Cluster) my_helm.MyHelmClient {
	fake.helmClientMutex.Lock()
	ret, specificReturn := fake.helmClientReturnsOnCall[len(fake.helmClientArgsForCall)]
	fake.helmClientArgsForCall = append(fake.helmClientArgsForCall, struct {
		cluster k8s.Cluster
	}{cluster})
	fake.recordInvocation("HelmClient", []interface{}{cluster})
	fake.helmClientMutex.Unlock()
	if fake.HelmClientStub != nil {
		return fake.HelmClientStub(cluster)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.helmClientReturns.result1
}

func (fake *FakeHelmClientFactory) HelmClientCallCount() int {
	fake.helmClientMutex.RLock()
	defer fake.helmClientMutex.RUnlock()
	return len(fake.helmClientArgsForCall)
}

func (fake *FakeHelmClientFactory) HelmClientArgsForCall(i int) k8s.Cluster {
	fake.helmClientMutex.RLock()
	defer fake.helmClientMutex.RUnlock()
	return fake.helmClientArgsForCall[i].cluster
}

func (fake *FakeHelmClientFactory) HelmClientReturns(result1 my_helm.MyHelmClient) {
	fake.HelmClientStub = nil
	fake.helmClientReturns = struct {
		result1 my_helm.MyHelmClient
	}{result1}
}

func (fake *FakeHelmClientFactory) HelmClientReturnsOnCall(i int, result1 my_helm.MyHelmClient) {
	fake.HelmClientStub = nil
	if fake.helmClientReturnsOnCall == nil {
		fake.helmClientReturnsOnCall = make(map[int]struct {
			result1 my_helm.MyHelmClient
		})
	}
	fake.helmClientReturnsOnCall[i] = struct {
		result1 my_helm.MyHelmClient
	}{result1}
}

func (fake *FakeHelmClientFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.helmClientMutex.RLock()
	defer fake.helmClientMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeHelmClientFactory) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ broker.HelmClientFactory = new(FakeHelmClientFactory)
