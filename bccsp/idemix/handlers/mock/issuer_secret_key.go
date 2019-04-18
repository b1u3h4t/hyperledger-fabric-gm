// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/dddengyunjie/hyperledger-fabric-gm/bccsp/idemix/handlers"
)

type IssuerSecretKey struct {
	BytesStub        func() ([]byte, error)
	bytesMutex       sync.RWMutex
	bytesArgsForCall []struct{}
	bytesReturns     struct {
		result1 []byte
		result2 error
	}
	bytesReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	PublicStub        func() handlers.IssuerPublicKey
	publicMutex       sync.RWMutex
	publicArgsForCall []struct{}
	publicReturns     struct {
		result1 handlers.IssuerPublicKey
	}
	publicReturnsOnCall map[int]struct {
		result1 handlers.IssuerPublicKey
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *IssuerSecretKey) Bytes() ([]byte, error) {
	fake.bytesMutex.Lock()
	ret, specificReturn := fake.bytesReturnsOnCall[len(fake.bytesArgsForCall)]
	fake.bytesArgsForCall = append(fake.bytesArgsForCall, struct{}{})
	fake.recordInvocation("Bytes", []interface{}{})
	fake.bytesMutex.Unlock()
	if fake.BytesStub != nil {
		return fake.BytesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.bytesReturns.result1, fake.bytesReturns.result2
}

func (fake *IssuerSecretKey) BytesCallCount() int {
	fake.bytesMutex.RLock()
	defer fake.bytesMutex.RUnlock()
	return len(fake.bytesArgsForCall)
}

func (fake *IssuerSecretKey) BytesReturns(result1 []byte, result2 error) {
	fake.BytesStub = nil
	fake.bytesReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *IssuerSecretKey) BytesReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.BytesStub = nil
	if fake.bytesReturnsOnCall == nil {
		fake.bytesReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.bytesReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *IssuerSecretKey) Public() handlers.IssuerPublicKey {
	fake.publicMutex.Lock()
	ret, specificReturn := fake.publicReturnsOnCall[len(fake.publicArgsForCall)]
	fake.publicArgsForCall = append(fake.publicArgsForCall, struct{}{})
	fake.recordInvocation("Public", []interface{}{})
	fake.publicMutex.Unlock()
	if fake.PublicStub != nil {
		return fake.PublicStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.publicReturns.result1
}

func (fake *IssuerSecretKey) PublicCallCount() int {
	fake.publicMutex.RLock()
	defer fake.publicMutex.RUnlock()
	return len(fake.publicArgsForCall)
}

func (fake *IssuerSecretKey) PublicReturns(result1 handlers.IssuerPublicKey) {
	fake.PublicStub = nil
	fake.publicReturns = struct {
		result1 handlers.IssuerPublicKey
	}{result1}
}

func (fake *IssuerSecretKey) PublicReturnsOnCall(i int, result1 handlers.IssuerPublicKey) {
	fake.PublicStub = nil
	if fake.publicReturnsOnCall == nil {
		fake.publicReturnsOnCall = make(map[int]struct {
			result1 handlers.IssuerPublicKey
		})
	}
	fake.publicReturnsOnCall[i] = struct {
		result1 handlers.IssuerPublicKey
	}{result1}
}

func (fake *IssuerSecretKey) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.bytesMutex.RLock()
	defer fake.bytesMutex.RUnlock()
	fake.publicMutex.RLock()
	defer fake.publicMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *IssuerSecretKey) recordInvocation(key string, args []interface{}) {
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

var _ handlers.IssuerSecretKey = new(IssuerSecretKey)
