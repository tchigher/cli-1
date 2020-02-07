// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	v7 "code.cloudfoundry.org/cli/command/v7"
	"code.cloudfoundry.org/cli/resources"
)

type FakeOrgQuotaActor struct {
	GetOrganizationQuotaByNameStub        func(string) (resources.OrganizationQuota, v7action.Warnings, error)
	getOrganizationQuotaByNameMutex       sync.RWMutex
	getOrganizationQuotaByNameArgsForCall []struct {
		arg1 string
	}
	getOrganizationQuotaByNameReturns struct {
		result1 resources.OrganizationQuota
		result2 v7action.Warnings
		result3 error
	}
	getOrganizationQuotaByNameReturnsOnCall map[int]struct {
		result1 resources.OrganizationQuota
		result2 v7action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeOrgQuotaActor) GetOrganizationQuotaByName(arg1 string) (resources.OrganizationQuota, v7action.Warnings, error) {
	fake.getOrganizationQuotaByNameMutex.Lock()
	ret, specificReturn := fake.getOrganizationQuotaByNameReturnsOnCall[len(fake.getOrganizationQuotaByNameArgsForCall)]
	fake.getOrganizationQuotaByNameArgsForCall = append(fake.getOrganizationQuotaByNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetOrganizationQuotaByName", []interface{}{arg1})
	fake.getOrganizationQuotaByNameMutex.Unlock()
	if fake.GetOrganizationQuotaByNameStub != nil {
		return fake.GetOrganizationQuotaByNameStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getOrganizationQuotaByNameReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeOrgQuotaActor) GetOrganizationQuotaByNameCallCount() int {
	fake.getOrganizationQuotaByNameMutex.RLock()
	defer fake.getOrganizationQuotaByNameMutex.RUnlock()
	return len(fake.getOrganizationQuotaByNameArgsForCall)
}

func (fake *FakeOrgQuotaActor) GetOrganizationQuotaByNameCalls(stub func(string) (resources.OrganizationQuota, v7action.Warnings, error)) {
	fake.getOrganizationQuotaByNameMutex.Lock()
	defer fake.getOrganizationQuotaByNameMutex.Unlock()
	fake.GetOrganizationQuotaByNameStub = stub
}

func (fake *FakeOrgQuotaActor) GetOrganizationQuotaByNameArgsForCall(i int) string {
	fake.getOrganizationQuotaByNameMutex.RLock()
	defer fake.getOrganizationQuotaByNameMutex.RUnlock()
	argsForCall := fake.getOrganizationQuotaByNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeOrgQuotaActor) GetOrganizationQuotaByNameReturns(result1 resources.OrganizationQuota, result2 v7action.Warnings, result3 error) {
	fake.getOrganizationQuotaByNameMutex.Lock()
	defer fake.getOrganizationQuotaByNameMutex.Unlock()
	fake.GetOrganizationQuotaByNameStub = nil
	fake.getOrganizationQuotaByNameReturns = struct {
		result1 resources.OrganizationQuota
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeOrgQuotaActor) GetOrganizationQuotaByNameReturnsOnCall(i int, result1 resources.OrganizationQuota, result2 v7action.Warnings, result3 error) {
	fake.getOrganizationQuotaByNameMutex.Lock()
	defer fake.getOrganizationQuotaByNameMutex.Unlock()
	fake.GetOrganizationQuotaByNameStub = nil
	if fake.getOrganizationQuotaByNameReturnsOnCall == nil {
		fake.getOrganizationQuotaByNameReturnsOnCall = make(map[int]struct {
			result1 resources.OrganizationQuota
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getOrganizationQuotaByNameReturnsOnCall[i] = struct {
		result1 resources.OrganizationQuota
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeOrgQuotaActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getOrganizationQuotaByNameMutex.RLock()
	defer fake.getOrganizationQuotaByNameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeOrgQuotaActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.OrgQuotaActor = new(FakeOrgQuotaActor)
