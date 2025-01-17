// Code generated by MockGen. DO NOT EDIT.
// Source: current.go

// Package coremock is a generated GoMock package.
package coremock

import (
	rsa "crypto/rsa"
	kubernetes "github.com/caos/orbos/pkg/kubernetes"
	operator "github.com/caos/zitadel/operator"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDatabaseCurrent is a mock of DatabaseCurrent interface
type MockDatabaseCurrent struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseCurrentMockRecorder
}

// MockDatabaseCurrentMockRecorder is the mock recorder for MockDatabaseCurrent
type MockDatabaseCurrentMockRecorder struct {
	mock *MockDatabaseCurrent
}

// NewMockDatabaseCurrent creates a new mock instance
func NewMockDatabaseCurrent(ctrl *gomock.Controller) *MockDatabaseCurrent {
	mock := &MockDatabaseCurrent{ctrl: ctrl}
	mock.recorder = &MockDatabaseCurrentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDatabaseCurrent) EXPECT() *MockDatabaseCurrentMockRecorder {
	return m.recorder
}

// GetURL mocks base method
func (m *MockDatabaseCurrent) GetURL() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetURL")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetURL indicates an expected call of GetURL
func (mr *MockDatabaseCurrentMockRecorder) GetURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetURL", reflect.TypeOf((*MockDatabaseCurrent)(nil).GetURL))
}

// GetPort mocks base method
func (m *MockDatabaseCurrent) GetPort() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPort")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPort indicates an expected call of GetPort
func (mr *MockDatabaseCurrentMockRecorder) GetPort() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPort", reflect.TypeOf((*MockDatabaseCurrent)(nil).GetPort))
}

// GetReadyQuery mocks base method
func (m *MockDatabaseCurrent) GetReadyQuery() operator.EnsureFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReadyQuery")
	ret0, _ := ret[0].(operator.EnsureFunc)
	return ret0
}

// GetReadyQuery indicates an expected call of GetReadyQuery
func (mr *MockDatabaseCurrentMockRecorder) GetReadyQuery() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReadyQuery", reflect.TypeOf((*MockDatabaseCurrent)(nil).GetReadyQuery))
}

// GetCertificateKey mocks base method
func (m *MockDatabaseCurrent) GetCertificateKey() *rsa.PrivateKey {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCertificateKey")
	ret0, _ := ret[0].(*rsa.PrivateKey)
	return ret0
}

// GetCertificateKey indicates an expected call of GetCertificateKey
func (mr *MockDatabaseCurrentMockRecorder) GetCertificateKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCertificateKey", reflect.TypeOf((*MockDatabaseCurrent)(nil).GetCertificateKey))
}

// SetCertificateKey mocks base method
func (m *MockDatabaseCurrent) SetCertificateKey(arg0 *rsa.PrivateKey) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCertificateKey", arg0)
}

// SetCertificateKey indicates an expected call of SetCertificateKey
func (mr *MockDatabaseCurrentMockRecorder) SetCertificateKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCertificateKey", reflect.TypeOf((*MockDatabaseCurrent)(nil).SetCertificateKey), arg0)
}

// GetCertificate mocks base method
func (m *MockDatabaseCurrent) GetCertificate() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCertificate")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetCertificate indicates an expected call of GetCertificate
func (mr *MockDatabaseCurrentMockRecorder) GetCertificate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCertificate", reflect.TypeOf((*MockDatabaseCurrent)(nil).GetCertificate))
}

// SetCertificate mocks base method
func (m *MockDatabaseCurrent) SetCertificate(arg0 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCertificate", arg0)
}

// SetCertificate indicates an expected call of SetCertificate
func (mr *MockDatabaseCurrentMockRecorder) SetCertificate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCertificate", reflect.TypeOf((*MockDatabaseCurrent)(nil).SetCertificate), arg0)
}

// GetAddUserFunc mocks base method
func (m *MockDatabaseCurrent) GetAddUserFunc() func(string) (operator.QueryFunc, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddUserFunc")
	ret0, _ := ret[0].(func(string) (operator.QueryFunc, error))
	return ret0
}

// GetAddUserFunc indicates an expected call of GetAddUserFunc
func (mr *MockDatabaseCurrentMockRecorder) GetAddUserFunc() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddUserFunc", reflect.TypeOf((*MockDatabaseCurrent)(nil).GetAddUserFunc))
}

// GetDeleteUserFunc mocks base method
func (m *MockDatabaseCurrent) GetDeleteUserFunc() func(string) (operator.DestroyFunc, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeleteUserFunc")
	ret0, _ := ret[0].(func(string) (operator.DestroyFunc, error))
	return ret0
}

// GetDeleteUserFunc indicates an expected call of GetDeleteUserFunc
func (mr *MockDatabaseCurrentMockRecorder) GetDeleteUserFunc() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeleteUserFunc", reflect.TypeOf((*MockDatabaseCurrent)(nil).GetDeleteUserFunc))
}

// GetListUsersFunc mocks base method
func (m *MockDatabaseCurrent) GetListUsersFunc() func(kubernetes.ClientInt) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListUsersFunc")
	ret0, _ := ret[0].(func(kubernetes.ClientInt) ([]string, error))
	return ret0
}

// GetListUsersFunc indicates an expected call of GetListUsersFunc
func (mr *MockDatabaseCurrentMockRecorder) GetListUsersFunc() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListUsersFunc", reflect.TypeOf((*MockDatabaseCurrent)(nil).GetListUsersFunc))
}

// GetListDatabasesFunc mocks base method
func (m *MockDatabaseCurrent) GetListDatabasesFunc() func(kubernetes.ClientInt) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListDatabasesFunc")
	ret0, _ := ret[0].(func(kubernetes.ClientInt) ([]string, error))
	return ret0
}

// GetListDatabasesFunc indicates an expected call of GetListDatabasesFunc
func (mr *MockDatabaseCurrentMockRecorder) GetListDatabasesFunc() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListDatabasesFunc", reflect.TypeOf((*MockDatabaseCurrent)(nil).GetListDatabasesFunc))
}
