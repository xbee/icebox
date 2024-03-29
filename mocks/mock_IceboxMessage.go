// Code generated by MockGen. DO NOT EDIT.
// Source: handler_test.go

// Package mocks is a generated GoMock package.
package mocks

import (
	protos "github.com/conseweb/icebox/protos"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIceboxMsgIntf is a mock of IceboxMsgIntf interface
type MockIceboxMsgIntf struct {
	ctrl     *gomock.Controller
	recorder *MockIceboxMsgIntfMockRecorder
}

// MockIceboxMsgIntfMockRecorder is the mock recorder for MockIceboxMsgIntf
type MockIceboxMsgIntfMockRecorder struct {
	mock *MockIceboxMsgIntf
}

// NewMockIceboxMsgIntf creates a new mock instance
func NewMockIceboxMsgIntf(ctrl *gomock.Controller) *MockIceboxMsgIntf {
	mock := &MockIceboxMsgIntf{ctrl: ctrl}
	mock.recorder = &MockIceboxMsgIntfMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIceboxMsgIntf) EXPECT() *MockIceboxMsgIntfMockRecorder {
	return m.recorder
}

// Reset mocks base method
func (m *MockIceboxMsgIntf) Reset() {
	m.ctrl.Call(m, "Reset")
}

// Reset indicates an expected call of Reset
func (mr *MockIceboxMsgIntfMockRecorder) Reset() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockIceboxMsgIntf)(nil).Reset))
}

// String mocks base method
func (m *MockIceboxMsgIntf) String() string {
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String
func (mr *MockIceboxMsgIntfMockRecorder) String() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockIceboxMsgIntf)(nil).String))
}

// ProtoMessage mocks base method
func (m *MockIceboxMsgIntf) ProtoMessage() {
	m.ctrl.Call(m, "ProtoMessage")
}

// ProtoMessage indicates an expected call of ProtoMessage
func (mr *MockIceboxMsgIntfMockRecorder) ProtoMessage() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProtoMessage", reflect.TypeOf((*MockIceboxMsgIntf)(nil).ProtoMessage))
}

// Descriptor mocks base method
func (m *MockIceboxMsgIntf) Descriptor() ([]byte, []int) {
	ret := m.ctrl.Call(m, "Descriptor")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].([]int)
	return ret0, ret1
}

// Descriptor indicates an expected call of Descriptor
func (mr *MockIceboxMsgIntfMockRecorder) Descriptor() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Descriptor", reflect.TypeOf((*MockIceboxMsgIntf)(nil).Descriptor))
}

// GetVersion mocks base method
func (m *MockIceboxMsgIntf) GetVersion() uint32 {
	ret := m.ctrl.Call(m, "GetVersion")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// GetVersion indicates an expected call of GetVersion
func (mr *MockIceboxMsgIntfMockRecorder) GetVersion() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersion", reflect.TypeOf((*MockIceboxMsgIntf)(nil).GetVersion))
}

// GetType mocks base method
func (m *MockIceboxMsgIntf) GetType() protos.IceboxMessage_Type {
	ret := m.ctrl.Call(m, "GetType")
	ret0, _ := ret[0].(protos.IceboxMessage_Type)
	return ret0
}

// GetType indicates an expected call of GetType
func (mr *MockIceboxMsgIntfMockRecorder) GetType() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetType", reflect.TypeOf((*MockIceboxMsgIntf)(nil).GetType))
}

// GetSessionId mocks base method
func (m *MockIceboxMsgIntf) GetSessionId() uint32 {
	ret := m.ctrl.Call(m, "GetSessionId")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// GetSessionId indicates an expected call of GetSessionId
func (mr *MockIceboxMsgIntfMockRecorder) GetSessionId() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionId", reflect.TypeOf((*MockIceboxMsgIntf)(nil).GetSessionId))
}

// GetPayload mocks base method
func (m *MockIceboxMsgIntf) GetPayload() []byte {
	ret := m.ctrl.Call(m, "GetPayload")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetPayload indicates an expected call of GetPayload
func (mr *MockIceboxMsgIntfMockRecorder) GetPayload() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPayload", reflect.TypeOf((*MockIceboxMsgIntf)(nil).GetPayload))
}

// GetSignature mocks base method
func (m *MockIceboxMsgIntf) GetSignature() []byte {
	ret := m.ctrl.Call(m, "GetSignature")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetSignature indicates an expected call of GetSignature
func (mr *MockIceboxMsgIntfMockRecorder) GetSignature() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSignature", reflect.TypeOf((*MockIceboxMsgIntf)(nil).GetSignature))
}
