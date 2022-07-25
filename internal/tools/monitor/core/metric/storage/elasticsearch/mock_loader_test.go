// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package elasticsearch

import (
	"context"
	"reflect"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/olivere/elastic"

	"github.com/erda-project/erda/internal/tools/monitor/core/storekit/elasticsearch/index"
	"github.com/erda-project/erda/internal/tools/monitor/core/storekit/elasticsearch/index/loader"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// AllIndices mocks base method.
func (m *MockInterface) AllIndices() *loader.IndexGroup {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllIndices")
	ret0, _ := ret[0].(*loader.IndexGroup)
	return ret0
}

// AllIndices indicates an expected call of AllIndices.
func (mr *MockInterfaceMockRecorder) AllIndices() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllIndices", reflect.TypeOf((*MockInterface)(nil).AllIndices))
}

// Client mocks base method.
func (m *MockInterface) Client() *elastic.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Client")
	ret0, _ := ret[0].(*elastic.Client)
	return ret0
}

// Client indicates an expected call of Client.
func (mr *MockInterfaceMockRecorder) Client() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockInterface)(nil).Client))
}

// IndexGroup mocks base method.
func (m *MockInterface) IndexGroup(path ...string) *loader.IndexGroup {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range path {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IndexGroup", varargs...)
	ret0, _ := ret[0].(*loader.IndexGroup)
	return ret0
}

// IndexGroup indicates an expected call of IndexGroup.
func (mr *MockInterfaceMockRecorder) IndexGroup(path ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IndexGroup", reflect.TypeOf((*MockInterface)(nil).IndexGroup), path...)
}

// Indices mocks base method.
func (m *MockInterface) Indices(ctx context.Context, start, end int64, path ...loader.KeyPath) []string {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, start, end}
	for _, a := range path {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Indices", varargs...)
	ret0, _ := ret[0].([]string)
	return ret0
}

// Indices indicates an expected call of Indices.
func (mr *MockInterfaceMockRecorder) Indices(ctx, start, end interface{}, path ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, start, end}, path...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Indices", reflect.TypeOf((*MockInterface)(nil).Indices), varargs...)
}

// Keys mocks base method.
func (m *MockInterface) Keys(path ...string) []string {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range path {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Keys", varargs...)
	ret0, _ := ret[0].([]string)
	return ret0
}

// Keys indicates an expected call of Keys.
func (mr *MockInterfaceMockRecorder) Keys(path ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockInterface)(nil).Keys), path...)
}

// LoadMode mocks base method.
func (m *MockInterface) LoadMode() loader.LoadMode {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadMode")
	ret0, _ := ret[0].(loader.LoadMode)
	return ret0
}

// LoadMode indicates an expected call of LoadMode.
func (mr *MockInterfaceMockRecorder) LoadMode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadMode", reflect.TypeOf((*MockInterface)(nil).LoadMode))
}

// Match mocks base method.
func (m *MockInterface) Match(idx string) *index.MatchResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Match", idx)
	ret0, _ := ret[0].(*index.MatchResult)
	return ret0
}

// Match indicates an expected call of Match.
func (mr *MockInterfaceMockRecorder) Match(index interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Match", reflect.TypeOf((*MockInterface)(nil).Match), index)
}

// Prefixes mocks base method.
func (m *MockInterface) Prefixes() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prefixes")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Prefixes indicates an expected call of Prefixes.
func (mr *MockInterfaceMockRecorder) Prefixes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prefixes", reflect.TypeOf((*MockInterface)(nil).Prefixes))
}

// QueryIndexTimeRange mocks base method.
func (m *MockInterface) QueryIndexTimeRange() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryIndexTimeRange")
	ret0, _ := ret[0].(bool)
	return ret0
}

// QueryIndexTimeRange indicates an expected call of QueryIndexTimeRange.
func (mr *MockInterfaceMockRecorder) QueryIndexTimeRange() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryIndexTimeRange", reflect.TypeOf((*MockInterface)(nil).QueryIndexTimeRange))
}

// ReloadIndices mocks base method.
func (m *MockInterface) ReloadIndices() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadIndices")
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadIndices indicates an expected call of ReloadIndices.
func (mr *MockInterfaceMockRecorder) ReloadIndices() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadIndices", reflect.TypeOf((*MockInterface)(nil).ReloadIndices))
}

// RequestTimeout mocks base method.
func (m *MockInterface) RequestTimeout() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestTimeout")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// RequestTimeout indicates an expected call of RequestTimeout.
func (mr *MockInterfaceMockRecorder) RequestTimeout() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestTimeout", reflect.TypeOf((*MockInterface)(nil).RequestTimeout))
}

// URLs mocks base method.
func (m *MockInterface) URLs() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "URLs")
	ret0, _ := ret[0].(string)
	return ret0
}

// URLs indicates an expected call of URLs.
func (mr *MockInterfaceMockRecorder) URLs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "URLs", reflect.TypeOf((*MockInterface)(nil).URLs))
}

// WaitAndGetIndices mocks base method.
func (m *MockInterface) WaitAndGetIndices(ctx context.Context) *loader.IndexGroup {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitAndGetIndices", ctx)
	ret0, _ := ret[0].(*loader.IndexGroup)
	return ret0
}

// WaitAndGetIndices indicates an expected call of WaitAndGetIndices.
func (mr *MockInterfaceMockRecorder) WaitAndGetIndices(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitAndGetIndices", reflect.TypeOf((*MockInterface)(nil).WaitAndGetIndices), ctx)
}

// WatchLoadEvent mocks base method.
func (m *MockInterface) WatchLoadEvent(arg0 func(*loader.IndexGroup)) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WatchLoadEvent", arg0)
}

// WatchLoadEvent indicates an expected call of WatchLoadEvent.
func (mr *MockInterfaceMockRecorder) WatchLoadEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchLoadEvent", reflect.TypeOf((*MockInterface)(nil).WatchLoadEvent), arg0)
}
