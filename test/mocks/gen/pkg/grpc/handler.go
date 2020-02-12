// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/grpc/handler.go

// Package grpc is a generated GoMock package.
package grpc

import (
	gomock "github.com/golang/mock/gomock"
	grpc "github.com/matheuscscp/fd8-judge/pkg/grpc"
	grpc0 "google.golang.org/grpc"
	reflect "reflect"
)

// MockRegisterable is a mock of Registerable interface
type MockRegisterable struct {
	ctrl     *gomock.Controller
	recorder *MockRegisterableMockRecorder
}

// MockRegisterableMockRecorder is the mock recorder for MockRegisterable
type MockRegisterableMockRecorder struct {
	mock *MockRegisterable
}

// NewMockRegisterable creates a new mock instance
func NewMockRegisterable(ctrl *gomock.Controller) *MockRegisterable {
	mock := &MockRegisterable{ctrl: ctrl}
	mock.recorder = &MockRegisterableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRegisterable) EXPECT() *MockRegisterableMockRecorder {
	return m.recorder
}

// Register mocks base method
func (m *MockRegisterable) Register(server *grpc0.Server) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Register", server)
}

// Register indicates an expected call of Register
func (mr *MockRegisterableMockRecorder) Register(server interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockRegisterable)(nil).Register), server)
}

// GetGatewayRegisterFunc mocks base method
func (m *MockRegisterable) GetGatewayRegisterFunc() grpc.GatewayRegisterFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGatewayRegisterFunc")
	ret0, _ := ret[0].(grpc.GatewayRegisterFunc)
	return ret0
}

// GetGatewayRegisterFunc indicates an expected call of GetGatewayRegisterFunc
func (mr *MockRegisterableMockRecorder) GetGatewayRegisterFunc() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGatewayRegisterFunc", reflect.TypeOf((*MockRegisterable)(nil).GetGatewayRegisterFunc))
}