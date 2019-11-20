// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/services/program.go

// Package services is a generated GoMock package.
package services

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	exec "os/exec"
	reflect "reflect"
)

// MockProgramService is a mock of ProgramService interface
type MockProgramService struct {
	ctrl     *gomock.Controller
	recorder *MockProgramServiceMockRecorder
}

// MockProgramServiceMockRecorder is the mock recorder for MockProgramService
type MockProgramServiceMockRecorder struct {
	mock *MockProgramService
}

// NewMockProgramService creates a new mock instance
func NewMockProgramService(ctrl *gomock.Controller) *MockProgramService {
	mock := &MockProgramService{ctrl: ctrl}
	mock.recorder = &MockProgramServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProgramService) EXPECT() *MockProgramServiceMockRecorder {
	return m.recorder
}

// Compile mocks base method
func (m *MockProgramService) Compile(ctx context.Context, sourceRelativePath, binaryRelativePath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Compile", ctx, sourceRelativePath, binaryRelativePath)
	ret0, _ := ret[0].(error)
	return ret0
}

// Compile indicates an expected call of Compile
func (mr *MockProgramServiceMockRecorder) Compile(ctx, sourceRelativePath, binaryRelativePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Compile", reflect.TypeOf((*MockProgramService)(nil).Compile), ctx, sourceRelativePath, binaryRelativePath)
}

// GetExecutionCommand mocks base method
func (m *MockProgramService) GetExecutionCommand(ctx context.Context, sourceRelativePath, binaryRelativePath string) *exec.Cmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExecutionCommand", ctx, sourceRelativePath, binaryRelativePath)
	ret0, _ := ret[0].(*exec.Cmd)
	return ret0
}

// GetExecutionCommand indicates an expected call of GetExecutionCommand
func (mr *MockProgramServiceMockRecorder) GetExecutionCommand(ctx, sourceRelativePath, binaryRelativePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExecutionCommand", reflect.TypeOf((*MockProgramService)(nil).GetExecutionCommand), ctx, sourceRelativePath, binaryRelativePath)
}

// GetSourceFileExtension mocks base method
func (m *MockProgramService) GetSourceFileExtension() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSourceFileExtension")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSourceFileExtension indicates an expected call of GetSourceFileExtension
func (mr *MockProgramServiceMockRecorder) GetSourceFileExtension() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSourceFileExtension", reflect.TypeOf((*MockProgramService)(nil).GetSourceFileExtension))
}

// GetBinaryFileExtension mocks base method
func (m *MockProgramService) GetBinaryFileExtension() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBinaryFileExtension")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetBinaryFileExtension indicates an expected call of GetBinaryFileExtension
func (mr *MockProgramServiceMockRecorder) GetBinaryFileExtension() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBinaryFileExtension", reflect.TypeOf((*MockProgramService)(nil).GetBinaryFileExtension))
}

// MockprogramServiceRuntime is a mock of programServiceRuntime interface
type MockprogramServiceRuntime struct {
	ctrl     *gomock.Controller
	recorder *MockprogramServiceRuntimeMockRecorder
}

// MockprogramServiceRuntimeMockRecorder is the mock recorder for MockprogramServiceRuntime
type MockprogramServiceRuntimeMockRecorder struct {
	mock *MockprogramServiceRuntime
}

// NewMockprogramServiceRuntime creates a new mock instance
func NewMockprogramServiceRuntime(ctrl *gomock.Controller) *MockprogramServiceRuntime {
	mock := &MockprogramServiceRuntime{ctrl: ctrl}
	mock.recorder = &MockprogramServiceRuntimeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockprogramServiceRuntime) EXPECT() *MockprogramServiceRuntimeMockRecorder {
	return m.recorder
}

// Run mocks base method
func (m *MockprogramServiceRuntime) Run(cmd *exec.Cmd) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", cmd)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
func (mr *MockprogramServiceRuntimeMockRecorder) Run(cmd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockprogramServiceRuntime)(nil).Run), cmd)
}
