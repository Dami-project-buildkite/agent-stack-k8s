// Code generated by MockGen. DO NOT EDIT.
// Source: scheduler.go

// Package scheduler_test is a generated GoMock package.
package scheduler_test

import (
	context "context"
	reflect "reflect"

	monitor "github.com/buildkite/agent-stack-k8s/monitor"
	gomock "github.com/golang/mock/gomock"
)

// MockJobHandler is a mock of JobHandler interface.
type MockJobHandler struct {
	ctrl     *gomock.Controller
	recorder *MockJobHandlerMockRecorder
}

// MockJobHandlerMockRecorder is the mock recorder for MockJobHandler.
type MockJobHandlerMockRecorder struct {
	mock *MockJobHandler
}

// NewMockJobHandler creates a new mock instance.
func NewMockJobHandler(ctrl *gomock.Controller) *MockJobHandler {
	mock := &MockJobHandler{ctrl: ctrl}
	mock.recorder = &MockJobHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJobHandler) EXPECT() *MockJobHandlerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockJobHandler) Create(arg0 context.Context, arg1 *monitor.Job) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockJobHandlerMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockJobHandler)(nil).Create), arg0, arg1)
}
