// Code generated by counterfeiter. DO NOT EDIT.
package sbomfakes

import (
	"sync"

	"github.com/interlynk-io/sbomqs/pkg/sbom"
)

type FakeAuthor struct {
	EmailStub        func() string
	emailMutex       sync.RWMutex
	emailArgsForCall []struct {
	}
	emailReturns struct {
		result1 string
	}
	emailReturnsOnCall map[int]struct {
		result1 string
	}
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct {
	}
	nameReturns struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
	TypeStub        func() string
	typeMutex       sync.RWMutex
	typeArgsForCall []struct {
	}
	typeReturns struct {
		result1 string
	}
	typeReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAuthor) GetPhone() string {
	// Implement the method as needed
	return ""
}

func (fake *FakeAuthor) GetEmail() string {
	fake.emailMutex.Lock()
	ret, specificReturn := fake.emailReturnsOnCall[len(fake.emailArgsForCall)]
	fake.emailArgsForCall = append(fake.emailArgsForCall, struct {
	}{})
	stub := fake.EmailStub
	fakeReturns := fake.emailReturns
	fake.recordInvocation("Email", []interface{}{})
	fake.emailMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeAuthor) EmailCallCount() int {
	fake.emailMutex.RLock()
	defer fake.emailMutex.RUnlock()
	return len(fake.emailArgsForCall)
}

func (fake *FakeAuthor) EmailCalls(stub func() string) {
	fake.emailMutex.Lock()
	defer fake.emailMutex.Unlock()
	fake.EmailStub = stub
}

func (fake *FakeAuthor) EmailReturns(result1 string) {
	fake.emailMutex.Lock()
	defer fake.emailMutex.Unlock()
	fake.EmailStub = nil
	fake.emailReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeAuthor) EmailReturnsOnCall(i int, result1 string) {
	fake.emailMutex.Lock()
	defer fake.emailMutex.Unlock()
	fake.EmailStub = nil
	if fake.emailReturnsOnCall == nil {
		fake.emailReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.emailReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeAuthor) GetName() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct {
	}{})
	stub := fake.NameStub
	fakeReturns := fake.nameReturns
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeAuthor) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeAuthor) NameCalls(stub func() string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = stub
}

func (fake *FakeAuthor) NameReturns(result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeAuthor) NameReturnsOnCall(i int, result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeAuthor) GetType() string {
	fake.typeMutex.Lock()
	ret, specificReturn := fake.typeReturnsOnCall[len(fake.typeArgsForCall)]
	fake.typeArgsForCall = append(fake.typeArgsForCall, struct {
	}{})
	stub := fake.TypeStub
	fakeReturns := fake.typeReturns
	fake.recordInvocation("Type", []interface{}{})
	fake.typeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeAuthor) TypeCallCount() int {
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	return len(fake.typeArgsForCall)
}

func (fake *FakeAuthor) TypeCalls(stub func() string) {
	fake.typeMutex.Lock()
	defer fake.typeMutex.Unlock()
	fake.TypeStub = stub
}

func (fake *FakeAuthor) TypeReturns(result1 string) {
	fake.typeMutex.Lock()
	defer fake.typeMutex.Unlock()
	fake.TypeStub = nil
	fake.typeReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeAuthor) TypeReturnsOnCall(i int, result1 string) {
	fake.typeMutex.Lock()
	defer fake.typeMutex.Unlock()
	fake.TypeStub = nil
	if fake.typeReturnsOnCall == nil {
		fake.typeReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.typeReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeAuthor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.emailMutex.RLock()
	defer fake.emailMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAuthor) recordInvocation(key string, args []interface{}) {
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

var _ sbom.GetAuthor = new(FakeAuthor)
