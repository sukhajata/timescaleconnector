// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/sukhajata/timescaleconnector/timescale (interfaces: DBClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	ppuplink "github.com/sukhajata/ppmessage/ppuplink"
	reflect "reflect"
)

// MockDBClient is a mock of DBClient interface
type MockDBClient struct {
	ctrl     *gomock.Controller
	recorder *MockDBClientMockRecorder
}

// MockDBClientMockRecorder is the mock recorder for MockDBClient
type MockDBClientMockRecorder struct {
	mock *MockDBClient
}

// NewMockDBClient creates a new mock instance
func NewMockDBClient(ctrl *gomock.Controller) *MockDBClient {
	mock := &MockDBClient{ctrl: ctrl}
	mock.recorder = &MockDBClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDBClient) EXPECT() *MockDBClientMockRecorder {
	return m.recorder
}

// InsertAlarmMsg mocks base method
func (m *MockDBClient) InsertAlarmMsg(arg0 ppuplink.AlarmMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertAlarmMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertAlarmMsg indicates an expected call of InsertAlarmMsg
func (mr *MockDBClientMockRecorder) InsertAlarmMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertAlarmMsg", reflect.TypeOf((*MockDBClient)(nil).InsertAlarmMsg), arg0)
}

// InsertCircuitEnergyMsg mocks base method
func (m *MockDBClient) InsertCircuitEnergyMsg(arg0 ppuplink.CircuitEnergyMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertCircuitEnergyMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertCircuitEnergyMsg indicates an expected call of InsertCircuitEnergyMsg
func (mr *MockDBClientMockRecorder) InsertCircuitEnergyMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertCircuitEnergyMsg", reflect.TypeOf((*MockDBClient)(nil).InsertCircuitEnergyMsg), arg0)
}

// InsertCircuitLoadMessage mocks base method
func (m *MockDBClient) InsertCircuitLoadMessage(arg0 ppuplink.CircuitLoadMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertCircuitLoadMessage", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertCircuitLoadMessage indicates an expected call of InsertCircuitLoadMessage
func (mr *MockDBClientMockRecorder) InsertCircuitLoadMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertCircuitLoadMessage", reflect.TypeOf((*MockDBClient)(nil).InsertCircuitLoadMessage), arg0)
}

// InsertEnergyMsg mocks base method
func (m *MockDBClient) InsertEnergyMsg(arg0 ppuplink.EnergyMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertEnergyMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertEnergyMsg indicates an expected call of InsertEnergyMsg
func (mr *MockDBClientMockRecorder) InsertEnergyMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertEnergyMsg", reflect.TypeOf((*MockDBClient)(nil).InsertEnergyMsg), arg0)
}

// InsertGatewayMsg mocks base method
func (m *MockDBClient) InsertGatewayMsg(arg0 ppuplink.GatewayMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertGatewayMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertGatewayMsg indicates an expected call of InsertGatewayMsg
func (mr *MockDBClientMockRecorder) InsertGatewayMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertGatewayMsg", reflect.TypeOf((*MockDBClient)(nil).InsertGatewayMsg), arg0)
}

// InsertGeoscanMsg mocks base method
func (m *MockDBClient) InsertGeoscanMsg(arg0 ppuplink.GeoscanMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertGeoscanMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertGeoscanMsg indicates an expected call of InsertGeoscanMsg
func (mr *MockDBClientMockRecorder) InsertGeoscanMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertGeoscanMsg", reflect.TypeOf((*MockDBClient)(nil).InsertGeoscanMsg), arg0)
}

// InsertHVAlarmMsg mocks base method
func (m *MockDBClient) InsertHVAlarmMsg(arg0 ppuplink.HVAlarmMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertHVAlarmMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertHVAlarmMsg indicates an expected call of InsertHVAlarmMsg
func (mr *MockDBClientMockRecorder) InsertHVAlarmMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertHVAlarmMsg", reflect.TypeOf((*MockDBClient)(nil).InsertHVAlarmMsg), arg0)
}

// InsertHarmonicsLowerMsg mocks base method
func (m *MockDBClient) InsertHarmonicsLowerMsg(arg0 ppuplink.HarmonicsLowerMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertHarmonicsLowerMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertHarmonicsLowerMsg indicates an expected call of InsertHarmonicsLowerMsg
func (mr *MockDBClientMockRecorder) InsertHarmonicsLowerMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertHarmonicsLowerMsg", reflect.TypeOf((*MockDBClient)(nil).InsertHarmonicsLowerMsg), arg0)
}

// InsertHarmonicsUpperMsg mocks base method
func (m *MockDBClient) InsertHarmonicsUpperMsg(arg0 ppuplink.HarmonicsUpperMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertHarmonicsUpperMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertHarmonicsUpperMsg indicates an expected call of InsertHarmonicsUpperMsg
func (mr *MockDBClientMockRecorder) InsertHarmonicsUpperMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertHarmonicsUpperMsg", reflect.TypeOf((*MockDBClient)(nil).InsertHarmonicsUpperMsg), arg0)
}

// InsertInstMsg mocks base method
func (m *MockDBClient) InsertInstMsg(arg0 ppuplink.InstMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertInstMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertInstMsg indicates an expected call of InsertInstMsg
func (mr *MockDBClientMockRecorder) InsertInstMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertInstMsg", reflect.TypeOf((*MockDBClient)(nil).InsertInstMsg), arg0)
}

// InsertInstP2PMsg mocks base method
func (m *MockDBClient) InsertInstP2PMsg(arg0 ppuplink.InstP2PMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertInstP2PMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertInstP2PMsg indicates an expected call of InsertInstP2PMsg
func (mr *MockDBClientMockRecorder) InsertInstP2PMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertInstP2PMsg", reflect.TypeOf((*MockDBClient)(nil).InsertInstP2PMsg), arg0)
}

// InsertPQEventsMsg mocks base method
func (m *MockDBClient) InsertPQEventsMsg(arg0 ppuplink.PQEventsMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertPQEventsMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertPQEventsMsg indicates an expected call of InsertPQEventsMsg
func (mr *MockDBClientMockRecorder) InsertPQEventsMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertPQEventsMsg", reflect.TypeOf((*MockDBClient)(nil).InsertPQEventsMsg), arg0)
}

// InsertPQMsg mocks base method
func (m *MockDBClient) InsertPQMsg(arg0 ppuplink.PQMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertPQMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertPQMsg indicates an expected call of InsertPQMsg
func (mr *MockDBClientMockRecorder) InsertPQMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertPQMsg", reflect.TypeOf((*MockDBClient)(nil).InsertPQMsg), arg0)
}

// InsertProcessedMsg mocks base method
func (m *MockDBClient) InsertProcessedMsg(arg0 ppuplink.ProcessedMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertProcessedMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertProcessedMsg indicates an expected call of InsertProcessedMsg
func (mr *MockDBClientMockRecorder) InsertProcessedMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertProcessedMsg", reflect.TypeOf((*MockDBClient)(nil).InsertProcessedMsg), arg0)
}

// InsertResendResponseMsg mocks base method
func (m *MockDBClient) InsertResendResponseMsg(arg0 ppuplink.ResendResponseMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertResendResponseMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertResendResponseMsg indicates an expected call of InsertResendResponseMsg
func (mr *MockDBClientMockRecorder) InsertResendResponseMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertResendResponseMsg", reflect.TypeOf((*MockDBClient)(nil).InsertResendResponseMsg), arg0)
}

// InsertS11PQ mocks base method
func (m *MockDBClient) InsertS11PQ(arg0 ppuplink.S11PQMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertS11PQ", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertS11PQ indicates an expected call of InsertS11PQ
func (mr *MockDBClientMockRecorder) InsertS11PQ(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertS11PQ", reflect.TypeOf((*MockDBClient)(nil).InsertS11PQ), arg0)
}

// InsertUplinkMsg mocks base method
func (m *MockDBClient) InsertUplinkMsg(arg0 ppuplink.UplinkMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUplinkMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertUplinkMsg indicates an expected call of InsertUplinkMsg
func (mr *MockDBClientMockRecorder) InsertUplinkMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUplinkMsg", reflect.TypeOf((*MockDBClient)(nil).InsertUplinkMsg), arg0)
}

// InsertVoltageStats mocks base method
func (m *MockDBClient) InsertVoltageStats(arg0 ppuplink.VoltageStatsMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertVoltageStats", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertVoltageStats indicates an expected call of InsertVoltageStats
func (mr *MockDBClientMockRecorder) InsertVoltageStats(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertVoltageStats", reflect.TypeOf((*MockDBClient)(nil).InsertVoltageStats), arg0)
}

// Listen mocks base method
func (m *MockDBClient) Listen(arg0 <-chan interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Listen", arg0)
}

// Listen indicates an expected call of Listen
func (mr *MockDBClientMockRecorder) Listen(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Listen", reflect.TypeOf((*MockDBClient)(nil).Listen), arg0)
}
