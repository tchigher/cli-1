// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	v7 "code.cloudfoundry.org/cli/command/v7"
	"code.cloudfoundry.org/cli/types"
)

type FakeDeleteLabelActor struct {
	UpdateApplicationLabelsByApplicationNameStub        func(string, string, map[string]types.NullString) (v7action.Warnings, error)
	updateApplicationLabelsByApplicationNameMutex       sync.RWMutex
	updateApplicationLabelsByApplicationNameArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 map[string]types.NullString
	}
	updateApplicationLabelsByApplicationNameReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	updateApplicationLabelsByApplicationNameReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	UpdateOrganizationLabelsByOrganizationNameStub        func(string, map[string]types.NullString) (v7action.Warnings, error)
	updateOrganizationLabelsByOrganizationNameMutex       sync.RWMutex
	updateOrganizationLabelsByOrganizationNameArgsForCall []struct {
		arg1 string
		arg2 map[string]types.NullString
	}
	updateOrganizationLabelsByOrganizationNameReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	updateOrganizationLabelsByOrganizationNameReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	UpdateSpaceLabelsBySpaceNameStub        func(string, string, map[string]types.NullString) (v7action.Warnings, error)
	updateSpaceLabelsBySpaceNameMutex       sync.RWMutex
	updateSpaceLabelsBySpaceNameArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 map[string]types.NullString
	}
	updateSpaceLabelsBySpaceNameReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	updateSpaceLabelsBySpaceNameReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDeleteLabelActor) UpdateApplicationLabelsByApplicationName(arg1 string, arg2 string, arg3 map[string]types.NullString) (v7action.Warnings, error) {
	fake.updateApplicationLabelsByApplicationNameMutex.Lock()
	ret, specificReturn := fake.updateApplicationLabelsByApplicationNameReturnsOnCall[len(fake.updateApplicationLabelsByApplicationNameArgsForCall)]
	fake.updateApplicationLabelsByApplicationNameArgsForCall = append(fake.updateApplicationLabelsByApplicationNameArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 map[string]types.NullString
	}{arg1, arg2, arg3})
	fake.recordInvocation("UpdateApplicationLabelsByApplicationName", []interface{}{arg1, arg2, arg3})
	fake.updateApplicationLabelsByApplicationNameMutex.Unlock()
	if fake.UpdateApplicationLabelsByApplicationNameStub != nil {
		return fake.UpdateApplicationLabelsByApplicationNameStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateApplicationLabelsByApplicationNameReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDeleteLabelActor) UpdateApplicationLabelsByApplicationNameCallCount() int {
	fake.updateApplicationLabelsByApplicationNameMutex.RLock()
	defer fake.updateApplicationLabelsByApplicationNameMutex.RUnlock()
	return len(fake.updateApplicationLabelsByApplicationNameArgsForCall)
}

func (fake *FakeDeleteLabelActor) UpdateApplicationLabelsByApplicationNameCalls(stub func(string, string, map[string]types.NullString) (v7action.Warnings, error)) {
	fake.updateApplicationLabelsByApplicationNameMutex.Lock()
	defer fake.updateApplicationLabelsByApplicationNameMutex.Unlock()
	fake.UpdateApplicationLabelsByApplicationNameStub = stub
}

func (fake *FakeDeleteLabelActor) UpdateApplicationLabelsByApplicationNameArgsForCall(i int) (string, string, map[string]types.NullString) {
	fake.updateApplicationLabelsByApplicationNameMutex.RLock()
	defer fake.updateApplicationLabelsByApplicationNameMutex.RUnlock()
	argsForCall := fake.updateApplicationLabelsByApplicationNameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeDeleteLabelActor) UpdateApplicationLabelsByApplicationNameReturns(result1 v7action.Warnings, result2 error) {
	fake.updateApplicationLabelsByApplicationNameMutex.Lock()
	defer fake.updateApplicationLabelsByApplicationNameMutex.Unlock()
	fake.UpdateApplicationLabelsByApplicationNameStub = nil
	fake.updateApplicationLabelsByApplicationNameReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeDeleteLabelActor) UpdateApplicationLabelsByApplicationNameReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.updateApplicationLabelsByApplicationNameMutex.Lock()
	defer fake.updateApplicationLabelsByApplicationNameMutex.Unlock()
	fake.UpdateApplicationLabelsByApplicationNameStub = nil
	if fake.updateApplicationLabelsByApplicationNameReturnsOnCall == nil {
		fake.updateApplicationLabelsByApplicationNameReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.updateApplicationLabelsByApplicationNameReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeDeleteLabelActor) UpdateOrganizationLabelsByOrganizationName(arg1 string, arg2 map[string]types.NullString) (v7action.Warnings, error) {
	fake.updateOrganizationLabelsByOrganizationNameMutex.Lock()
	ret, specificReturn := fake.updateOrganizationLabelsByOrganizationNameReturnsOnCall[len(fake.updateOrganizationLabelsByOrganizationNameArgsForCall)]
	fake.updateOrganizationLabelsByOrganizationNameArgsForCall = append(fake.updateOrganizationLabelsByOrganizationNameArgsForCall, struct {
		arg1 string
		arg2 map[string]types.NullString
	}{arg1, arg2})
	fake.recordInvocation("UpdateOrganizationLabelsByOrganizationName", []interface{}{arg1, arg2})
	fake.updateOrganizationLabelsByOrganizationNameMutex.Unlock()
	if fake.UpdateOrganizationLabelsByOrganizationNameStub != nil {
		return fake.UpdateOrganizationLabelsByOrganizationNameStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateOrganizationLabelsByOrganizationNameReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDeleteLabelActor) UpdateOrganizationLabelsByOrganizationNameCallCount() int {
	fake.updateOrganizationLabelsByOrganizationNameMutex.RLock()
	defer fake.updateOrganizationLabelsByOrganizationNameMutex.RUnlock()
	return len(fake.updateOrganizationLabelsByOrganizationNameArgsForCall)
}

func (fake *FakeDeleteLabelActor) UpdateOrganizationLabelsByOrganizationNameCalls(stub func(string, map[string]types.NullString) (v7action.Warnings, error)) {
	fake.updateOrganizationLabelsByOrganizationNameMutex.Lock()
	defer fake.updateOrganizationLabelsByOrganizationNameMutex.Unlock()
	fake.UpdateOrganizationLabelsByOrganizationNameStub = stub
}

func (fake *FakeDeleteLabelActor) UpdateOrganizationLabelsByOrganizationNameArgsForCall(i int) (string, map[string]types.NullString) {
	fake.updateOrganizationLabelsByOrganizationNameMutex.RLock()
	defer fake.updateOrganizationLabelsByOrganizationNameMutex.RUnlock()
	argsForCall := fake.updateOrganizationLabelsByOrganizationNameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeDeleteLabelActor) UpdateOrganizationLabelsByOrganizationNameReturns(result1 v7action.Warnings, result2 error) {
	fake.updateOrganizationLabelsByOrganizationNameMutex.Lock()
	defer fake.updateOrganizationLabelsByOrganizationNameMutex.Unlock()
	fake.UpdateOrganizationLabelsByOrganizationNameStub = nil
	fake.updateOrganizationLabelsByOrganizationNameReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeDeleteLabelActor) UpdateOrganizationLabelsByOrganizationNameReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.updateOrganizationLabelsByOrganizationNameMutex.Lock()
	defer fake.updateOrganizationLabelsByOrganizationNameMutex.Unlock()
	fake.UpdateOrganizationLabelsByOrganizationNameStub = nil
	if fake.updateOrganizationLabelsByOrganizationNameReturnsOnCall == nil {
		fake.updateOrganizationLabelsByOrganizationNameReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.updateOrganizationLabelsByOrganizationNameReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeDeleteLabelActor) UpdateSpaceLabelsBySpaceName(arg1 string, arg2 string, arg3 map[string]types.NullString) (v7action.Warnings, error) {
	fake.updateSpaceLabelsBySpaceNameMutex.Lock()
	ret, specificReturn := fake.updateSpaceLabelsBySpaceNameReturnsOnCall[len(fake.updateSpaceLabelsBySpaceNameArgsForCall)]
	fake.updateSpaceLabelsBySpaceNameArgsForCall = append(fake.updateSpaceLabelsBySpaceNameArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 map[string]types.NullString
	}{arg1, arg2, arg3})
	fake.recordInvocation("UpdateSpaceLabelsBySpaceName", []interface{}{arg1, arg2, arg3})
	fake.updateSpaceLabelsBySpaceNameMutex.Unlock()
	if fake.UpdateSpaceLabelsBySpaceNameStub != nil {
		return fake.UpdateSpaceLabelsBySpaceNameStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateSpaceLabelsBySpaceNameReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDeleteLabelActor) UpdateSpaceLabelsBySpaceNameCallCount() int {
	fake.updateSpaceLabelsBySpaceNameMutex.RLock()
	defer fake.updateSpaceLabelsBySpaceNameMutex.RUnlock()
	return len(fake.updateSpaceLabelsBySpaceNameArgsForCall)
}

func (fake *FakeDeleteLabelActor) UpdateSpaceLabelsBySpaceNameCalls(stub func(string, string, map[string]types.NullString) (v7action.Warnings, error)) {
	fake.updateSpaceLabelsBySpaceNameMutex.Lock()
	defer fake.updateSpaceLabelsBySpaceNameMutex.Unlock()
	fake.UpdateSpaceLabelsBySpaceNameStub = stub
}

func (fake *FakeDeleteLabelActor) UpdateSpaceLabelsBySpaceNameArgsForCall(i int) (string, string, map[string]types.NullString) {
	fake.updateSpaceLabelsBySpaceNameMutex.RLock()
	defer fake.updateSpaceLabelsBySpaceNameMutex.RUnlock()
	argsForCall := fake.updateSpaceLabelsBySpaceNameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeDeleteLabelActor) UpdateSpaceLabelsBySpaceNameReturns(result1 v7action.Warnings, result2 error) {
	fake.updateSpaceLabelsBySpaceNameMutex.Lock()
	defer fake.updateSpaceLabelsBySpaceNameMutex.Unlock()
	fake.UpdateSpaceLabelsBySpaceNameStub = nil
	fake.updateSpaceLabelsBySpaceNameReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeDeleteLabelActor) UpdateSpaceLabelsBySpaceNameReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.updateSpaceLabelsBySpaceNameMutex.Lock()
	defer fake.updateSpaceLabelsBySpaceNameMutex.Unlock()
	fake.UpdateSpaceLabelsBySpaceNameStub = nil
	if fake.updateSpaceLabelsBySpaceNameReturnsOnCall == nil {
		fake.updateSpaceLabelsBySpaceNameReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.updateSpaceLabelsBySpaceNameReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeDeleteLabelActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.updateApplicationLabelsByApplicationNameMutex.RLock()
	defer fake.updateApplicationLabelsByApplicationNameMutex.RUnlock()
	fake.updateOrganizationLabelsByOrganizationNameMutex.RLock()
	defer fake.updateOrganizationLabelsByOrganizationNameMutex.RUnlock()
	fake.updateSpaceLabelsBySpaceNameMutex.RLock()
	defer fake.updateSpaceLabelsBySpaceNameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDeleteLabelActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.DeleteLabelActor = new(FakeDeleteLabelActor)
