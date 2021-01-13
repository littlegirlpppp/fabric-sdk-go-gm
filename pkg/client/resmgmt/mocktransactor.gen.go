// Code generated by counterfeiter. DO NOT EDIT.
package resmgmt

import (
	"sync"

	"github.com/littlegirlpppp/fabric-sdk-go-gm/pkg/common/providers/fab"
)

type MockTransactor struct {
	CreateTransactionStub        func(request fab.TransactionRequest) (*fab.Transaction, error)
	createTransactionMutex       sync.RWMutex
	createTransactionArgsForCall []struct {
		request fab.TransactionRequest
	}
	createTransactionReturns struct {
		result1 *fab.Transaction
		result2 error
	}
	createTransactionReturnsOnCall map[int]struct {
		result1 *fab.Transaction
		result2 error
	}
	SendTransactionStub        func(tx *fab.Transaction) (*fab.TransactionResponse, error)
	sendTransactionMutex       sync.RWMutex
	sendTransactionArgsForCall []struct {
		tx *fab.Transaction
	}
	sendTransactionReturns struct {
		result1 *fab.TransactionResponse
		result2 error
	}
	sendTransactionReturnsOnCall map[int]struct {
		result1 *fab.TransactionResponse
		result2 error
	}
	CreateTransactionHeaderStub        func(opts ...fab.TxnHeaderOpt) (fab.TransactionHeader, error)
	createTransactionHeaderMutex       sync.RWMutex
	createTransactionHeaderArgsForCall []struct {
		opts []fab.TxnHeaderOpt
	}
	createTransactionHeaderReturns struct {
		result1 fab.TransactionHeader
		result2 error
	}
	createTransactionHeaderReturnsOnCall map[int]struct {
		result1 fab.TransactionHeader
		result2 error
	}
	SendTransactionProposalStub        func(*fab.TransactionProposal, []fab.ProposalProcessor) ([]*fab.TransactionProposalResponse, error)
	sendTransactionProposalMutex       sync.RWMutex
	sendTransactionProposalArgsForCall []struct {
		arg1 *fab.TransactionProposal
		arg2 []fab.ProposalProcessor
	}
	sendTransactionProposalReturns struct {
		result1 []*fab.TransactionProposalResponse
		result2 error
	}
	sendTransactionProposalReturnsOnCall map[int]struct {
		result1 []*fab.TransactionProposalResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *MockTransactor) CreateTransaction(request fab.TransactionRequest) (*fab.Transaction, error) {
	fake.createTransactionMutex.Lock()
	ret, specificReturn := fake.createTransactionReturnsOnCall[len(fake.createTransactionArgsForCall)]
	fake.createTransactionArgsForCall = append(fake.createTransactionArgsForCall, struct {
		request fab.TransactionRequest
	}{request})
	fake.recordInvocation("CreateTransaction", []interface{}{request})
	fake.createTransactionMutex.Unlock()
	if fake.CreateTransactionStub != nil {
		return fake.CreateTransactionStub(request)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createTransactionReturns.result1, fake.createTransactionReturns.result2
}

func (fake *MockTransactor) CreateTransactionCallCount() int {
	fake.createTransactionMutex.RLock()
	defer fake.createTransactionMutex.RUnlock()
	return len(fake.createTransactionArgsForCall)
}

func (fake *MockTransactor) CreateTransactionArgsForCall(i int) fab.TransactionRequest {
	fake.createTransactionMutex.RLock()
	defer fake.createTransactionMutex.RUnlock()
	return fake.createTransactionArgsForCall[i].request
}

func (fake *MockTransactor) CreateTransactionReturns(result1 *fab.Transaction, result2 error) {
	fake.CreateTransactionStub = nil
	fake.createTransactionReturns = struct {
		result1 *fab.Transaction
		result2 error
	}{result1, result2}
}

func (fake *MockTransactor) CreateTransactionReturnsOnCall(i int, result1 *fab.Transaction, result2 error) {
	fake.CreateTransactionStub = nil
	if fake.createTransactionReturnsOnCall == nil {
		fake.createTransactionReturnsOnCall = make(map[int]struct {
			result1 *fab.Transaction
			result2 error
		})
	}
	fake.createTransactionReturnsOnCall[i] = struct {
		result1 *fab.Transaction
		result2 error
	}{result1, result2}
}

func (fake *MockTransactor) SendTransaction(tx *fab.Transaction) (*fab.TransactionResponse, error) {
	fake.sendTransactionMutex.Lock()
	ret, specificReturn := fake.sendTransactionReturnsOnCall[len(fake.sendTransactionArgsForCall)]
	fake.sendTransactionArgsForCall = append(fake.sendTransactionArgsForCall, struct {
		tx *fab.Transaction
	}{tx})
	fake.recordInvocation("SendTransaction", []interface{}{tx})
	fake.sendTransactionMutex.Unlock()
	if fake.SendTransactionStub != nil {
		return fake.SendTransactionStub(tx)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.sendTransactionReturns.result1, fake.sendTransactionReturns.result2
}

func (fake *MockTransactor) SendTransactionCallCount() int {
	fake.sendTransactionMutex.RLock()
	defer fake.sendTransactionMutex.RUnlock()
	return len(fake.sendTransactionArgsForCall)
}

func (fake *MockTransactor) SendTransactionArgsForCall(i int) *fab.Transaction {
	fake.sendTransactionMutex.RLock()
	defer fake.sendTransactionMutex.RUnlock()
	return fake.sendTransactionArgsForCall[i].tx
}

func (fake *MockTransactor) SendTransactionReturns(result1 *fab.TransactionResponse, result2 error) {
	fake.SendTransactionStub = nil
	fake.sendTransactionReturns = struct {
		result1 *fab.TransactionResponse
		result2 error
	}{result1, result2}
}

func (fake *MockTransactor) SendTransactionReturnsOnCall(i int, result1 *fab.TransactionResponse, result2 error) {
	fake.SendTransactionStub = nil
	if fake.sendTransactionReturnsOnCall == nil {
		fake.sendTransactionReturnsOnCall = make(map[int]struct {
			result1 *fab.TransactionResponse
			result2 error
		})
	}
	fake.sendTransactionReturnsOnCall[i] = struct {
		result1 *fab.TransactionResponse
		result2 error
	}{result1, result2}
}

func (fake *MockTransactor) CreateTransactionHeader(opts ...fab.TxnHeaderOpt) (fab.TransactionHeader, error) {
	fake.createTransactionHeaderMutex.Lock()
	ret, specificReturn := fake.createTransactionHeaderReturnsOnCall[len(fake.createTransactionHeaderArgsForCall)]
	fake.createTransactionHeaderArgsForCall = append(fake.createTransactionHeaderArgsForCall, struct {
		opts []fab.TxnHeaderOpt
	}{opts})
	fake.recordInvocation("CreateTransactionHeader", []interface{}{opts})
	fake.createTransactionHeaderMutex.Unlock()
	if fake.CreateTransactionHeaderStub != nil {
		return fake.CreateTransactionHeaderStub(opts...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createTransactionHeaderReturns.result1, fake.createTransactionHeaderReturns.result2
}

func (fake *MockTransactor) CreateTransactionHeaderCallCount() int {
	fake.createTransactionHeaderMutex.RLock()
	defer fake.createTransactionHeaderMutex.RUnlock()
	return len(fake.createTransactionHeaderArgsForCall)
}

func (fake *MockTransactor) CreateTransactionHeaderArgsForCall(i int) []fab.TxnHeaderOpt {
	fake.createTransactionHeaderMutex.RLock()
	defer fake.createTransactionHeaderMutex.RUnlock()
	return fake.createTransactionHeaderArgsForCall[i].opts
}

func (fake *MockTransactor) CreateTransactionHeaderReturns(result1 fab.TransactionHeader, result2 error) {
	fake.CreateTransactionHeaderStub = nil
	fake.createTransactionHeaderReturns = struct {
		result1 fab.TransactionHeader
		result2 error
	}{result1, result2}
}

func (fake *MockTransactor) CreateTransactionHeaderReturnsOnCall(i int, result1 fab.TransactionHeader, result2 error) {
	fake.CreateTransactionHeaderStub = nil
	if fake.createTransactionHeaderReturnsOnCall == nil {
		fake.createTransactionHeaderReturnsOnCall = make(map[int]struct {
			result1 fab.TransactionHeader
			result2 error
		})
	}
	fake.createTransactionHeaderReturnsOnCall[i] = struct {
		result1 fab.TransactionHeader
		result2 error
	}{result1, result2}
}

func (fake *MockTransactor) SendTransactionProposal(arg1 *fab.TransactionProposal, arg2 []fab.ProposalProcessor) ([]*fab.TransactionProposalResponse, error) {
	var arg2Copy []fab.ProposalProcessor
	if arg2 != nil {
		arg2Copy = make([]fab.ProposalProcessor, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.sendTransactionProposalMutex.Lock()
	ret, specificReturn := fake.sendTransactionProposalReturnsOnCall[len(fake.sendTransactionProposalArgsForCall)]
	fake.sendTransactionProposalArgsForCall = append(fake.sendTransactionProposalArgsForCall, struct {
		arg1 *fab.TransactionProposal
		arg2 []fab.ProposalProcessor
	}{arg1, arg2Copy})
	fake.recordInvocation("SendTransactionProposal", []interface{}{arg1, arg2Copy})
	fake.sendTransactionProposalMutex.Unlock()
	if fake.SendTransactionProposalStub != nil {
		return fake.SendTransactionProposalStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.sendTransactionProposalReturns.result1, fake.sendTransactionProposalReturns.result2
}

func (fake *MockTransactor) SendTransactionProposalCallCount() int {
	fake.sendTransactionProposalMutex.RLock()
	defer fake.sendTransactionProposalMutex.RUnlock()
	return len(fake.sendTransactionProposalArgsForCall)
}

func (fake *MockTransactor) SendTransactionProposalArgsForCall(i int) (*fab.TransactionProposal, []fab.ProposalProcessor) {
	fake.sendTransactionProposalMutex.RLock()
	defer fake.sendTransactionProposalMutex.RUnlock()
	return fake.sendTransactionProposalArgsForCall[i].arg1, fake.sendTransactionProposalArgsForCall[i].arg2
}

func (fake *MockTransactor) SendTransactionProposalReturns(result1 []*fab.TransactionProposalResponse, result2 error) {
	fake.SendTransactionProposalStub = nil
	fake.sendTransactionProposalReturns = struct {
		result1 []*fab.TransactionProposalResponse
		result2 error
	}{result1, result2}
}

func (fake *MockTransactor) SendTransactionProposalReturnsOnCall(i int, result1 []*fab.TransactionProposalResponse, result2 error) {
	fake.SendTransactionProposalStub = nil
	if fake.sendTransactionProposalReturnsOnCall == nil {
		fake.sendTransactionProposalReturnsOnCall = make(map[int]struct {
			result1 []*fab.TransactionProposalResponse
			result2 error
		})
	}
	fake.sendTransactionProposalReturnsOnCall[i] = struct {
		result1 []*fab.TransactionProposalResponse
		result2 error
	}{result1, result2}
}

func (fake *MockTransactor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createTransactionMutex.RLock()
	defer fake.createTransactionMutex.RUnlock()
	fake.sendTransactionMutex.RLock()
	defer fake.sendTransactionMutex.RUnlock()
	fake.createTransactionHeaderMutex.RLock()
	defer fake.createTransactionHeaderMutex.RUnlock()
	fake.sendTransactionProposalMutex.RLock()
	defer fake.sendTransactionProposalMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *MockTransactor) recordInvocation(key string, args []interface{}) {
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

var _ fab.Transactor = new(MockTransactor)
