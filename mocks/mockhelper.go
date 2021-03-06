// Code generated by MockGen. DO NOT EDIT.
// Source: powerpilot.visualstudio.com/PowerPilot/_git/pphelpers.git (interfaces: AuthLoggerHelper)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	pplogger "github.com/sukhajata/pplogger"
	reflect "reflect"
)

// MockAuthLoggerHelper is a mock of AuthLoggerHelper interface
type MockAuthLoggerHelper struct {
	ctrl     *gomock.Controller
	recorder *MockAuthLoggerHelperMockRecorder
}

// MockAuthLoggerHelperMockRecorder is the mock recorder for MockAuthLoggerHelper
type MockAuthLoggerHelperMockRecorder struct {
	mock *MockAuthLoggerHelper
}

// NewMockAuthLoggerHelper creates a new mock instance
func NewMockAuthLoggerHelper(ctrl *gomock.Controller) *MockAuthLoggerHelper {
	mock := &MockAuthLoggerHelper{ctrl: ctrl}
	mock.recorder = &MockAuthLoggerHelperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthLoggerHelper) EXPECT() *MockAuthLoggerHelperMockRecorder {
	return m.recorder
}

// CheckToken mocks base method
func (m *MockAuthLoggerHelper) CheckToken(arg0 string, arg1 []string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckToken", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckToken indicates an expected call of CheckToken
func (mr *MockAuthLoggerHelperMockRecorder) CheckToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckToken", reflect.TypeOf((*MockAuthLoggerHelper)(nil).CheckToken), arg0, arg1)
}

// HandleGrpcError mocks base method
func (m *MockAuthLoggerHelper) HandleGrpcError(arg0 error) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleGrpcError", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// HandleGrpcError indicates an expected call of HandleGrpcError
func (mr *MockAuthLoggerHelperMockRecorder) HandleGrpcError(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleGrpcError", reflect.TypeOf((*MockAuthLoggerHelper)(nil).HandleGrpcError), arg0)
}

// LogError mocks base method
func (m *MockAuthLoggerHelper) LogError(arg0, arg1 string, arg2 pplogger.ErrorMessage_Severity) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LogError", arg0, arg1, arg2)
}

// LogError indicates an expected call of LogError
func (mr *MockAuthLoggerHelperMockRecorder) LogError(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogError", reflect.TypeOf((*MockAuthLoggerHelper)(nil).LogError), arg0, arg1, arg2)
}
