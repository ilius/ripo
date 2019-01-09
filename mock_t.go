// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ilius/ripo (interfaces: SmallT)

// Package mock_ripo is a generated GoMock package.
package ripo

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSmallT is a mock of SmallT interface
type MockSmallT struct {
	ctrl     *gomock.Controller
	recorder *MockSmallTMockRecorder
}

// MockSmallTMockRecorder is the mock recorder for MockSmallT
type MockSmallTMockRecorder struct {
	mock *MockSmallT
}

// NewMockSmallT creates a new mock instance
func NewMockSmallT(ctrl *gomock.Controller) *MockSmallT {
	mock := &MockSmallT{ctrl: ctrl}
	mock.recorder = &MockSmallTMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSmallT) EXPECT() *MockSmallTMockRecorder {
	return m.recorder
}

// Fatalf mocks base method
func (m *MockSmallT) Fatalf(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatalf", varargs...)
}

// Fatalf indicates an expected call of Fatalf
func (mr *MockSmallTMockRecorder) Fatalf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatalf", reflect.TypeOf((*MockSmallT)(nil).Fatalf), varargs...)
}

// Helper mocks base method
func (m *MockSmallT) Helper() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Helper")
}

// Helper indicates an expected call of Helper
func (mr *MockSmallTMockRecorder) Helper() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Helper", reflect.TypeOf((*MockSmallT)(nil).Helper))
}