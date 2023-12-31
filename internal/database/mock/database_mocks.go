// Code generated by MockGen. DO NOT EDIT.
// Source: internal/database/database.go
//
// Generated by this command:
//
//	mockgen -source=internal/database/database.go -destination=internal/database/mock/database_mocks.go -package=database_mock ComputerLayer StorageLayer
//
// Package database_mock is a generated GoMock package.
package database_mock

import (
	context "context"
	compute "quickquery/internal/database/compute"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockComputerLayer is a mock of ComputerLayer interface.
type MockComputerLayer struct {
	ctrl     *gomock.Controller
	recorder *MockComputerLayerMockRecorder
}

// MockComputerLayerMockRecorder is the mock recorder for MockComputerLayer.
type MockComputerLayerMockRecorder struct {
	mock *MockComputerLayer
}

// NewMockComputerLayer creates a new mock instance.
func NewMockComputerLayer(ctrl *gomock.Controller) *MockComputerLayer {
	mock := &MockComputerLayer{ctrl: ctrl}
	mock.recorder = &MockComputerLayerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComputerLayer) EXPECT() *MockComputerLayerMockRecorder {
	return m.recorder
}

// HandleQuery mocks base method.
func (m *MockComputerLayer) HandleQuery(ctx context.Context, queryStr string) (compute.Query, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleQuery", ctx, queryStr)
	ret0, _ := ret[0].(compute.Query)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleQuery indicates an expected call of HandleQuery.
func (mr *MockComputerLayerMockRecorder) HandleQuery(ctx, queryStr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleQuery", reflect.TypeOf((*MockComputerLayer)(nil).HandleQuery), ctx, queryStr)
}

// MockStorageLayer is a mock of StorageLayer interface.
type MockStorageLayer struct {
	ctrl     *gomock.Controller
	recorder *MockStorageLayerMockRecorder
}

// MockStorageLayerMockRecorder is the mock recorder for MockStorageLayer.
type MockStorageLayerMockRecorder struct {
	mock *MockStorageLayer
}

// NewMockStorageLayer creates a new mock instance.
func NewMockStorageLayer(ctrl *gomock.Controller) *MockStorageLayer {
	mock := &MockStorageLayer{ctrl: ctrl}
	mock.recorder = &MockStorageLayerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageLayer) EXPECT() *MockStorageLayerMockRecorder {
	return m.recorder
}

// Del mocks base method.
func (m *MockStorageLayer) Del(ctx context.Context, key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Del", ctx, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Del indicates an expected call of Del.
func (mr *MockStorageLayerMockRecorder) Del(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockStorageLayer)(nil).Del), ctx, key)
}

// Get mocks base method.
func (m *MockStorageLayer) Get(ctx context.Context, key string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockStorageLayerMockRecorder) Get(ctx, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStorageLayer)(nil).Get), ctx, key)
}

// Set mocks base method.
func (m *MockStorageLayer) Set(ctx context.Context, key, value string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockStorageLayerMockRecorder) Set(ctx, key, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockStorageLayer)(nil).Set), ctx, key, value)
}
