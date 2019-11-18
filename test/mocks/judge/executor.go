// Code generated by MockGen. DO NOT EDIT.
// Source: judge/executor.go

// Package judge is a generated GoMock package.
package judge

import (
	gomock "github.com/golang/mock/gomock"
)

// MockExecutorRuntime is a mock of ExecutorRuntime interface
type MockExecutorRuntime struct {
	ctrl     *gomock.Controller
	recorder *MockExecutorRuntimeMockRecorder
}

// MockExecutorRuntimeMockRecorder is the mock recorder for MockExecutorRuntime
type MockExecutorRuntimeMockRecorder struct {
	mock *MockExecutorRuntime
}

// NewMockExecutorRuntime creates a new mock instance
func NewMockExecutorRuntime(ctrl *gomock.Controller) *MockExecutorRuntime {
	mock := &MockExecutorRuntime{ctrl: ctrl}
	mock.recorder = &MockExecutorRuntimeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockExecutorRuntime) EXPECT() *MockExecutorRuntimeMockRecorder {
	return m.recorder
}