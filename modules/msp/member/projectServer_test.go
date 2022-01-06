// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/erda-project/erda-proto-go/msp/tenant/project/pb (interfaces: ProjectServiceServer)

// Package member is a generated GoMock package.
package member

import (
	context "context"
	reflect "reflect"

	pb "github.com/erda-project/erda-proto-go/msp/tenant/project/pb"
	gomock "github.com/golang/mock/gomock"
)

// MockProjectServiceServer is a mock of ProjectServiceServer interface.
type MockProjectServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockProjectServiceServerMockRecorder
}

// MockProjectServiceServerMockRecorder is the mock recorder for MockProjectServiceServer.
type MockProjectServiceServerMockRecorder struct {
	mock *MockProjectServiceServer
}

// NewMockProjectServiceServer creates a new mock instance.
func NewMockProjectServiceServer(ctrl *gomock.Controller) *MockProjectServiceServer {
	mock := &MockProjectServiceServer{ctrl: ctrl}
	mock.recorder = &MockProjectServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectServiceServer) EXPECT() *MockProjectServiceServerMockRecorder {
	return m.recorder
}

// CreateProject mocks base method.
func (m *MockProjectServiceServer) CreateProject(arg0 context.Context, arg1 *pb.CreateProjectRequest) (*pb.CreateProjectResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProject", arg0, arg1)
	ret0, _ := ret[0].(*pb.CreateProjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProject indicates an expected call of CreateProject.
func (mr *MockProjectServiceServerMockRecorder) CreateProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockProjectServiceServer)(nil).CreateProject), arg0, arg1)
}

// DeleteProject mocks base method.
func (m *MockProjectServiceServer) DeleteProject(arg0 context.Context, arg1 *pb.DeleteProjectRequest) (*pb.DeleteProjectResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProject", arg0, arg1)
	ret0, _ := ret[0].(*pb.DeleteProjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProject indicates an expected call of DeleteProject.
func (mr *MockProjectServiceServerMockRecorder) DeleteProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProject", reflect.TypeOf((*MockProjectServiceServer)(nil).DeleteProject), arg0, arg1)
}

// GetProject mocks base method.
func (m *MockProjectServiceServer) GetProject(arg0 context.Context, arg1 *pb.GetProjectRequest) (*pb.GetProjectResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProject", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetProjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProject indicates an expected call of GetProject.
func (mr *MockProjectServiceServerMockRecorder) GetProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProject", reflect.TypeOf((*MockProjectServiceServer)(nil).GetProject), arg0, arg1)
}

// GetProjectOverview mocks base method.
func (m *MockProjectServiceServer) GetProjectOverview(arg0 context.Context, arg1 *pb.GetProjectOverviewRequest) (*pb.GetProjectOverviewResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectOverview", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetProjectOverviewResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjectOverview indicates an expected call of GetProjectOverview.
func (mr *MockProjectServiceServerMockRecorder) GetProjectOverview(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectOverview", reflect.TypeOf((*MockProjectServiceServer)(nil).GetProjectOverview), arg0, arg1)
}

// GetProjectStatistics mocks base method.
func (m *MockProjectServiceServer) GetProjectStatistics(arg0 context.Context, arg1 *pb.GetProjectStatisticsRequest) (*pb.GetProjectStatisticsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectStatistics", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetProjectStatisticsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjectStatistics indicates an expected call of GetProjectStatistics.
func (mr *MockProjectServiceServerMockRecorder) GetProjectStatistics(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectStatistics", reflect.TypeOf((*MockProjectServiceServer)(nil).GetProjectStatistics), arg0, arg1)
}

// GetProjects mocks base method.
func (m *MockProjectServiceServer) GetProjects(arg0 context.Context, arg1 *pb.GetProjectsRequest) (*pb.GetProjectsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjects", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetProjectsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjects indicates an expected call of GetProjects.
func (mr *MockProjectServiceServerMockRecorder) GetProjects(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjects", reflect.TypeOf((*MockProjectServiceServer)(nil).GetProjects), arg0, arg1)
}

// GetProjectsTenantsIDs mocks base method.
func (m *MockProjectServiceServer) GetProjectsTenantsIDs(arg0 context.Context, arg1 *pb.GetProjectsTenantsIDsRequest) (*pb.GetProjectsTenantsIDsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectsTenantsIDs", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetProjectsTenantsIDsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjectsTenantsIDs indicates an expected call of GetProjectsTenantsIDs.
func (mr *MockProjectServiceServerMockRecorder) GetProjectsTenantsIDs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectsTenantsIDs", reflect.TypeOf((*MockProjectServiceServer)(nil).GetProjectsTenantsIDs), arg0, arg1)
}

// UpdateProject mocks base method.
func (m *MockProjectServiceServer) UpdateProject(arg0 context.Context, arg1 *pb.UpdateProjectRequest) (*pb.UpdateProjectResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProject", arg0, arg1)
	ret0, _ := ret[0].(*pb.UpdateProjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProject indicates an expected call of UpdateProject.
func (mr *MockProjectServiceServerMockRecorder) UpdateProject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProject", reflect.TypeOf((*MockProjectServiceServer)(nil).UpdateProject), arg0, arg1)
}
