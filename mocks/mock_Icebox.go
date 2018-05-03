// Code generated by MockGen. DO NOT EDIT.
// Source: icebox.pb.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockIceboxClient is a mock of IceboxClient interface
type MockIceboxClient struct {
	ctrl     *gomock.Controller
	recorder *MockIceboxClientMockRecorder
}

// MockIceboxClientMockRecorder is the mock recorder for MockIceboxClient
type MockIceboxClientMockRecorder struct {
	mock *MockIceboxClient
}

// NewMockIceboxClient creates a new mock instance
func NewMockIceboxClient(ctrl *gomock.Controller) *MockIceboxClient {
	mock := &MockIceboxClient{ctrl: ctrl}
	mock.recorder = &MockIceboxClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIceboxClient) EXPECT() *MockIceboxClientMockRecorder {
	return m.recorder
}

// Chat mocks base method
func (m *MockIceboxClient) Chat(ctx context.Context, in *IceboxMessage, opts ...grpc.CallOption) (*IceboxMessage, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Chat", varargs...)
	ret0, _ := ret[0].(*IceboxMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Chat indicates an expected call of Chat
func (mr *MockIceboxClientMockRecorder) Chat(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Chat", reflect.TypeOf((*MockIceboxClient)(nil).Chat), varargs...)
}

// MockIceboxServer is a mock of IceboxServer interface
type MockIceboxServer struct {
	ctrl     *gomock.Controller
	recorder *MockIceboxServerMockRecorder
}

// MockIceboxServerMockRecorder is the mock recorder for MockIceboxServer
type MockIceboxServerMockRecorder struct {
	mock *MockIceboxServer
}

// NewMockIceboxServer creates a new mock instance
func NewMockIceboxServer(ctrl *gomock.Controller) *MockIceboxServer {
	mock := &MockIceboxServer{ctrl: ctrl}
	mock.recorder = &MockIceboxServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIceboxServer) EXPECT() *MockIceboxServerMockRecorder {
	return m.recorder
}

// Chat mocks base method
func (m *MockIceboxServer) Chat(arg0 context.Context, arg1 *IceboxMessage) (*IceboxMessage, error) {
	ret := m.ctrl.Call(m, "Chat", arg0, arg1)
	ret0, _ := ret[0].(*IceboxMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Chat indicates an expected call of Chat
func (mr *MockIceboxServerMockRecorder) Chat(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Chat", reflect.TypeOf((*MockIceboxServer)(nil).Chat), arg0, arg1)
}