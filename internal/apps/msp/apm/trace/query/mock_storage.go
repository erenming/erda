// Code generated by MockGen. DO NOT EDIT.
// Source: ../storage/storage.go

// Package query is a generated GoMock package.
package query

import (
	"context"
	"reflect"

	"github.com/erda-project/erda/internal/apps/msp/apm/trace/storage"
	"github.com/erda-project/erda/internal/tools/monitor/core/storekit"

	"github.com/golang/mock/gomock"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockStorage) Count(ctx context.Context, traceId string) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, traceId)
	ret0, _ := ret[0].(int64)
	return ret0
}

// Count indicates an expected call of Count.
func (mr *MockStorageMockRecorder) Count(ctx, traceId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockStorage)(nil).Count), ctx, traceId)
}

// Iterator mocks base method.
func (m *MockStorage) Iterator(ctx context.Context, sel *storage.Selector) (storekit.Iterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Iterator", ctx, sel)
	ret0, _ := ret[0].(storekit.Iterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Iterator indicates an expected call of Iterator.
func (mr *MockStorageMockRecorder) Iterator(ctx, sel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iterator", reflect.TypeOf((*MockStorage)(nil).Iterator), ctx, sel)
}

// NewWriter mocks base method.
func (m *MockStorage) NewWriter(ctx context.Context) (storekit.BatchWriter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewWriter", ctx)
	ret0, _ := ret[0].(storekit.BatchWriter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewWriter indicates an expected call of NewWriter.
func (mr *MockStorageMockRecorder) NewWriter(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewWriter", reflect.TypeOf((*MockStorage)(nil).NewWriter), ctx)
}
