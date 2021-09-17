// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/iptables-logger/rotatablesink"
)

type DestinationFileInfo struct {
	FileExistsStub        func(string) (bool, error)
	fileExistsMutex       sync.RWMutex
	fileExistsArgsForCall []struct {
		arg1 string
	}
	fileExistsReturns struct {
		result1 bool
		result2 error
	}
	fileExistsReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	FileInodeStub        func(string) (uint64, error)
	fileInodeMutex       sync.RWMutex
	fileInodeArgsForCall []struct {
		arg1 string
	}
	fileInodeReturns struct {
		result1 uint64
		result2 error
	}
	fileInodeReturnsOnCall map[int]struct {
		result1 uint64
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *DestinationFileInfo) FileExists(arg1 string) (bool, error) {
	fake.fileExistsMutex.Lock()
	ret, specificReturn := fake.fileExistsReturnsOnCall[len(fake.fileExistsArgsForCall)]
	fake.fileExistsArgsForCall = append(fake.fileExistsArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.FileExistsStub
	fakeReturns := fake.fileExistsReturns
	fake.recordInvocation("FileExists", []interface{}{arg1})
	fake.fileExistsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *DestinationFileInfo) FileExistsCallCount() int {
	fake.fileExistsMutex.RLock()
	defer fake.fileExistsMutex.RUnlock()
	return len(fake.fileExistsArgsForCall)
}

func (fake *DestinationFileInfo) FileExistsCalls(stub func(string) (bool, error)) {
	fake.fileExistsMutex.Lock()
	defer fake.fileExistsMutex.Unlock()
	fake.FileExistsStub = stub
}

func (fake *DestinationFileInfo) FileExistsArgsForCall(i int) string {
	fake.fileExistsMutex.RLock()
	defer fake.fileExistsMutex.RUnlock()
	argsForCall := fake.fileExistsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *DestinationFileInfo) FileExistsReturns(result1 bool, result2 error) {
	fake.fileExistsMutex.Lock()
	defer fake.fileExistsMutex.Unlock()
	fake.FileExistsStub = nil
	fake.fileExistsReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *DestinationFileInfo) FileExistsReturnsOnCall(i int, result1 bool, result2 error) {
	fake.fileExistsMutex.Lock()
	defer fake.fileExistsMutex.Unlock()
	fake.FileExistsStub = nil
	if fake.fileExistsReturnsOnCall == nil {
		fake.fileExistsReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.fileExistsReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *DestinationFileInfo) FileInode(arg1 string) (uint64, error) {
	fake.fileInodeMutex.Lock()
	ret, specificReturn := fake.fileInodeReturnsOnCall[len(fake.fileInodeArgsForCall)]
	fake.fileInodeArgsForCall = append(fake.fileInodeArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.FileInodeStub
	fakeReturns := fake.fileInodeReturns
	fake.recordInvocation("FileInode", []interface{}{arg1})
	fake.fileInodeMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *DestinationFileInfo) FileInodeCallCount() int {
	fake.fileInodeMutex.RLock()
	defer fake.fileInodeMutex.RUnlock()
	return len(fake.fileInodeArgsForCall)
}

func (fake *DestinationFileInfo) FileInodeCalls(stub func(string) (uint64, error)) {
	fake.fileInodeMutex.Lock()
	defer fake.fileInodeMutex.Unlock()
	fake.FileInodeStub = stub
}

func (fake *DestinationFileInfo) FileInodeArgsForCall(i int) string {
	fake.fileInodeMutex.RLock()
	defer fake.fileInodeMutex.RUnlock()
	argsForCall := fake.fileInodeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *DestinationFileInfo) FileInodeReturns(result1 uint64, result2 error) {
	fake.fileInodeMutex.Lock()
	defer fake.fileInodeMutex.Unlock()
	fake.FileInodeStub = nil
	fake.fileInodeReturns = struct {
		result1 uint64
		result2 error
	}{result1, result2}
}

func (fake *DestinationFileInfo) FileInodeReturnsOnCall(i int, result1 uint64, result2 error) {
	fake.fileInodeMutex.Lock()
	defer fake.fileInodeMutex.Unlock()
	fake.FileInodeStub = nil
	if fake.fileInodeReturnsOnCall == nil {
		fake.fileInodeReturnsOnCall = make(map[int]struct {
			result1 uint64
			result2 error
		})
	}
	fake.fileInodeReturnsOnCall[i] = struct {
		result1 uint64
		result2 error
	}{result1, result2}
}

func (fake *DestinationFileInfo) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.fileExistsMutex.RLock()
	defer fake.fileExistsMutex.RUnlock()
	fake.fileInodeMutex.RLock()
	defer fake.fileInodeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *DestinationFileInfo) recordInvocation(key string, args []interface{}) {
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

var _ rotatablesink.DestinationFileInfo = new(DestinationFileInfo)