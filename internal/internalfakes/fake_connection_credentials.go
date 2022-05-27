// Code generated by counterfeiter. DO NOT EDIT.
package internalfakes

import (
	"github.com/rabbitmq/messaging-topology-operator/rabbitmqclient"
	"sync"
)

type FakeConnectionCredentials struct {
	DataStub        func(string) ([]byte, bool)
	dataMutex       sync.RWMutex
	dataArgsForCall []struct {
		arg1 string
	}
	dataReturns struct {
		result1 []byte
		result2 bool
	}
	dataReturnsOnCall map[int]struct {
		result1 []byte
		result2 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConnectionCredentials) Data(arg1 string) ([]byte, bool) {
	fake.dataMutex.Lock()
	ret, specificReturn := fake.dataReturnsOnCall[len(fake.dataArgsForCall)]
	fake.dataArgsForCall = append(fake.dataArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.DataStub
	fakeReturns := fake.dataReturns
	fake.recordInvocation("Data", []interface{}{arg1})
	fake.dataMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeConnectionCredentials) DataCallCount() int {
	fake.dataMutex.RLock()
	defer fake.dataMutex.RUnlock()
	return len(fake.dataArgsForCall)
}

func (fake *FakeConnectionCredentials) DataCalls(stub func(string) ([]byte, bool)) {
	fake.dataMutex.Lock()
	defer fake.dataMutex.Unlock()
	fake.DataStub = stub
}

func (fake *FakeConnectionCredentials) DataArgsForCall(i int) string {
	fake.dataMutex.RLock()
	defer fake.dataMutex.RUnlock()
	argsForCall := fake.dataArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeConnectionCredentials) DataReturns(result1 []byte, result2 bool) {
	fake.dataMutex.Lock()
	defer fake.dataMutex.Unlock()
	fake.DataStub = nil
	fake.dataReturns = struct {
		result1 []byte
		result2 bool
	}{result1, result2}
}

func (fake *FakeConnectionCredentials) DataReturnsOnCall(i int, result1 []byte, result2 bool) {
	fake.dataMutex.Lock()
	defer fake.dataMutex.Unlock()
	fake.DataStub = nil
	if fake.dataReturnsOnCall == nil {
		fake.dataReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 bool
		})
	}
	fake.dataReturnsOnCall[i] = struct {
		result1 []byte
		result2 bool
	}{result1, result2}
}

func (fake *FakeConnectionCredentials) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.dataMutex.RLock()
	defer fake.dataMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeConnectionCredentials) recordInvocation(key string, args []interface{}) {
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

var _ rabbitmqclient.ConnectionCredentials = new(FakeConnectionCredentials)
