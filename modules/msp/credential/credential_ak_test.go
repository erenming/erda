// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/erda-project/erda-proto-go/core/services/authentication/credentials/accesskey/pb (interfaces: AccessKeyServiceServer)

// Package exporter is a generated GoMock package.
package credential

import (
	context "context"
	reflect "reflect"

	pb "github.com/erda-project/erda-proto-go/core/services/authentication/credentials/accesskey/pb"
	gomock "github.com/golang/mock/gomock"
)

// MockAccessKeyServiceServer is a mock of AccessKeyServiceServer interface.
type MockAccessKeyServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockAccessKeyServiceServerMockRecorder
}

// MockAccessKeyServiceServerMockRecorder is the mock recorder for MockAccessKeyServiceServer.
type MockAccessKeyServiceServerMockRecorder struct {
	mock *MockAccessKeyServiceServer
}

// NewMockAccessKeyServiceServer creates a new mock instance.
func NewMockAccessKeyServiceServer(ctrl *gomock.Controller) *MockAccessKeyServiceServer {
	mock := &MockAccessKeyServiceServer{ctrl: ctrl}
	mock.recorder = &MockAccessKeyServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccessKeyServiceServer) EXPECT() *MockAccessKeyServiceServerMockRecorder {
	return m.recorder
}

// CreateAccessKey mocks base method.
func (m *MockAccessKeyServiceServer) CreateAccessKey(arg0 context.Context, arg1 *pb.CreateAccessKeyRequest) (*pb.CreateAccessKeyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccessKey", arg0, arg1)
	ret0, _ := ret[0].(*pb.CreateAccessKeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccessKey indicates an expected call of CreateAccessKey.
func (mr *MockAccessKeyServiceServerMockRecorder) CreateAccessKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccessKey", reflect.TypeOf((*MockAccessKeyServiceServer)(nil).CreateAccessKey), arg0, arg1)
}

// DeleteAccessKey mocks base method.
func (m *MockAccessKeyServiceServer) DeleteAccessKey(arg0 context.Context, arg1 *pb.DeleteAccessKeyRequest) (*pb.DeleteAccessKeyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccessKey", arg0, arg1)
	ret0, _ := ret[0].(*pb.DeleteAccessKeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAccessKey indicates an expected call of DeleteAccessKey.
func (mr *MockAccessKeyServiceServerMockRecorder) DeleteAccessKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccessKey", reflect.TypeOf((*MockAccessKeyServiceServer)(nil).DeleteAccessKey), arg0, arg1)
}

// GetAccessKey mocks base method.
func (m *MockAccessKeyServiceServer) GetAccessKey(arg0 context.Context, arg1 *pb.GetAccessKeyRequest) (*pb.GetAccessKeyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccessKey", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetAccessKeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessKey indicates an expected call of GetAccessKey.
func (mr *MockAccessKeyServiceServerMockRecorder) GetAccessKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessKey", reflect.TypeOf((*MockAccessKeyServiceServer)(nil).GetAccessKey), arg0, arg1)
}

// QueryAccessKeys mocks base method.
func (m *MockAccessKeyServiceServer) QueryAccessKeys(arg0 context.Context, arg1 *pb.QueryAccessKeysRequest) (*pb.QueryAccessKeysResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryAccessKeys", arg0, arg1)
	ret0, _ := ret[0].(*pb.QueryAccessKeysResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryAccessKeys indicates an expected call of QueryAccessKeys.
func (mr *MockAccessKeyServiceServerMockRecorder) QueryAccessKeys(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryAccessKeys", reflect.TypeOf((*MockAccessKeyServiceServer)(nil).QueryAccessKeys), arg0, arg1)
}

// UpdateAccessKey mocks base method.
func (m *MockAccessKeyServiceServer) UpdateAccessKey(arg0 context.Context, arg1 *pb.UpdateAccessKeyRequest) (*pb.UpdateAccessKeyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccessKey", arg0, arg1)
	ret0, _ := ret[0].(*pb.UpdateAccessKeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAccessKey indicates an expected call of UpdateAccessKey.
func (mr *MockAccessKeyServiceServerMockRecorder) UpdateAccessKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccessKey", reflect.TypeOf((*MockAccessKeyServiceServer)(nil).UpdateAccessKey), arg0, arg1)
}
