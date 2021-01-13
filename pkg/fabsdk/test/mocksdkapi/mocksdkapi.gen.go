// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/littlegirlpppp/fabric-sdk-go-gm/pkg/fabsdk/api (interfaces: CoreProviderFactory,MSPProviderFactory,ServiceProviderFactory)

// Package mocksdkapi is a generated GoMock package.
package mocksdkapi

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	options "github.com/littlegirlpppp/fabric-sdk-go-gm/pkg/common/options"
	core "github.com/littlegirlpppp/fabric-sdk-go-gm/pkg/common/providers/core"
	fab "github.com/littlegirlpppp/fabric-sdk-go-gm/pkg/common/providers/fab"
	msp "github.com/littlegirlpppp/fabric-sdk-go-gm/pkg/common/providers/msp"
)

// MockCoreProviderFactory is a mock of CoreProviderFactory interface
type MockCoreProviderFactory struct {
	ctrl     *gomock.Controller
	recorder *MockCoreProviderFactoryMockRecorder
}

// MockCoreProviderFactoryMockRecorder is the mock recorder for MockCoreProviderFactory
type MockCoreProviderFactoryMockRecorder struct {
	mock *MockCoreProviderFactory
}

// NewMockCoreProviderFactory creates a new mock instance
func NewMockCoreProviderFactory(ctrl *gomock.Controller) *MockCoreProviderFactory {
	mock := &MockCoreProviderFactory{ctrl: ctrl}
	mock.recorder = &MockCoreProviderFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCoreProviderFactory) EXPECT() *MockCoreProviderFactoryMockRecorder {
	return m.recorder
}

// CreateCryptoSuiteProvider mocks base method
func (m *MockCoreProviderFactory) CreateCryptoSuiteProvider(arg0 core.CryptoSuiteConfig) (core.CryptoSuite, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCryptoSuiteProvider", arg0)
	ret0, _ := ret[0].(core.CryptoSuite)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCryptoSuiteProvider indicates an expected call of CreateCryptoSuiteProvider
func (mr *MockCoreProviderFactoryMockRecorder) CreateCryptoSuiteProvider(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCryptoSuiteProvider", reflect.TypeOf((*MockCoreProviderFactory)(nil).CreateCryptoSuiteProvider), arg0)
}

// CreateInfraProvider mocks base method
func (m *MockCoreProviderFactory) CreateInfraProvider(arg0 fab.EndpointConfig) (fab.InfraProvider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInfraProvider", arg0)
	ret0, _ := ret[0].(fab.InfraProvider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateInfraProvider indicates an expected call of CreateInfraProvider
func (mr *MockCoreProviderFactoryMockRecorder) CreateInfraProvider(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInfraProvider", reflect.TypeOf((*MockCoreProviderFactory)(nil).CreateInfraProvider), arg0)
}

// CreateSigningManager mocks base method
func (m *MockCoreProviderFactory) CreateSigningManager(arg0 core.CryptoSuite) (core.SigningManager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSigningManager", arg0)
	ret0, _ := ret[0].(core.SigningManager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSigningManager indicates an expected call of CreateSigningManager
func (mr *MockCoreProviderFactoryMockRecorder) CreateSigningManager(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSigningManager", reflect.TypeOf((*MockCoreProviderFactory)(nil).CreateSigningManager), arg0)
}

// MockMSPProviderFactory is a mock of MSPProviderFactory interface
type MockMSPProviderFactory struct {
	ctrl     *gomock.Controller
	recorder *MockMSPProviderFactoryMockRecorder
}

// MockMSPProviderFactoryMockRecorder is the mock recorder for MockMSPProviderFactory
type MockMSPProviderFactoryMockRecorder struct {
	mock *MockMSPProviderFactory
}

// NewMockMSPProviderFactory creates a new mock instance
func NewMockMSPProviderFactory(ctrl *gomock.Controller) *MockMSPProviderFactory {
	mock := &MockMSPProviderFactory{ctrl: ctrl}
	mock.recorder = &MockMSPProviderFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMSPProviderFactory) EXPECT() *MockMSPProviderFactoryMockRecorder {
	return m.recorder
}

// CreateIdentityManagerProvider mocks base method
func (m *MockMSPProviderFactory) CreateIdentityManagerProvider(arg0 fab.EndpointConfig, arg1 core.CryptoSuite, arg2 msp.UserStore) (msp.IdentityManagerProvider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateIdentityManagerProvider", arg0, arg1, arg2)
	ret0, _ := ret[0].(msp.IdentityManagerProvider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateIdentityManagerProvider indicates an expected call of CreateIdentityManagerProvider
func (mr *MockMSPProviderFactoryMockRecorder) CreateIdentityManagerProvider(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIdentityManagerProvider", reflect.TypeOf((*MockMSPProviderFactory)(nil).CreateIdentityManagerProvider), arg0, arg1, arg2)
}

// CreateUserStore mocks base method
func (m *MockMSPProviderFactory) CreateUserStore(arg0 msp.IdentityConfig) (msp.UserStore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserStore", arg0)
	ret0, _ := ret[0].(msp.UserStore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUserStore indicates an expected call of CreateUserStore
func (mr *MockMSPProviderFactoryMockRecorder) CreateUserStore(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserStore", reflect.TypeOf((*MockMSPProviderFactory)(nil).CreateUserStore), arg0)
}

// MockServiceProviderFactory is a mock of ServiceProviderFactory interface
type MockServiceProviderFactory struct {
	ctrl     *gomock.Controller
	recorder *MockServiceProviderFactoryMockRecorder
}

// MockServiceProviderFactoryMockRecorder is the mock recorder for MockServiceProviderFactory
type MockServiceProviderFactoryMockRecorder struct {
	mock *MockServiceProviderFactory
}

// NewMockServiceProviderFactory creates a new mock instance
func NewMockServiceProviderFactory(ctrl *gomock.Controller) *MockServiceProviderFactory {
	mock := &MockServiceProviderFactory{ctrl: ctrl}
	mock.recorder = &MockServiceProviderFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServiceProviderFactory) EXPECT() *MockServiceProviderFactoryMockRecorder {
	return m.recorder
}

// CreateChannelProvider mocks base method
func (m *MockServiceProviderFactory) CreateChannelProvider(arg0 fab.EndpointConfig, arg1 ...options.Opt) (fab.ChannelProvider, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateChannelProvider", varargs...)
	ret0, _ := ret[0].(fab.ChannelProvider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateChannelProvider indicates an expected call of CreateChannelProvider
func (mr *MockServiceProviderFactoryMockRecorder) CreateChannelProvider(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChannelProvider", reflect.TypeOf((*MockServiceProviderFactory)(nil).CreateChannelProvider), varargs...)
}

// CreateLocalDiscoveryProvider mocks base method
func (m *MockServiceProviderFactory) CreateLocalDiscoveryProvider(arg0 fab.EndpointConfig) (fab.LocalDiscoveryProvider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLocalDiscoveryProvider", arg0)
	ret0, _ := ret[0].(fab.LocalDiscoveryProvider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLocalDiscoveryProvider indicates an expected call of CreateLocalDiscoveryProvider
func (mr *MockServiceProviderFactoryMockRecorder) CreateLocalDiscoveryProvider(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLocalDiscoveryProvider", reflect.TypeOf((*MockServiceProviderFactory)(nil).CreateLocalDiscoveryProvider), arg0)
}
