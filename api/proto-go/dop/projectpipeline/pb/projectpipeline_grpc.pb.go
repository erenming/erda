// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// Source: projectpipeline.proto

package pb

import (
	context "context"

	transport "github.com/erda-project/erda-infra/pkg/transport"
	grpc1 "github.com/erda-project/erda-infra/pkg/transport/grpc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion5

// ProjectPipelineServiceClient is the client API for ProjectPipelineService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectPipelineServiceClient interface {
	Create(ctx context.Context, in *CreateProjectPipelineRequest, opts ...grpc.CallOption) (*CreateProjectPipelineResponse, error)
	ListApp(ctx context.Context, in *ListAppRequest, opts ...grpc.CallOption) (*ListAppResponse, error)
	ListPipelineYml(ctx context.Context, in *ListAppPipelineYmlRequest, opts ...grpc.CallOption) (*ListAppPipelineYmlResponse, error)
	CreateNamePreCheck(ctx context.Context, in *CreateProjectPipelineNamePreCheckRequest, opts ...grpc.CallOption) (*CreateProjectPipelineNamePreCheckResponse, error)
	CreateSourcePreCheck(ctx context.Context, in *CreateProjectPipelineSourcePreCheckRequest, opts ...grpc.CallOption) (*CreateProjectPipelineSourcePreCheckResponse, error)
	ListPipelineCategory(ctx context.Context, in *ListPipelineCategoryRequest, opts ...grpc.CallOption) (*ListPipelineCategoryResponse, error)
	Update(ctx context.Context, in *UpdateProjectPipelineRequest, opts ...grpc.CallOption) (*UpdateProjectPipelineResponse, error)
	Run(ctx context.Context, in *RunProjectPipelineRequest, opts ...grpc.CallOption) (*RunProjectPipelineResponse, error)
	Rerun(ctx context.Context, in *RerunProjectPipelineRequest, opts ...grpc.CallOption) (*RerunProjectPipelineResponse, error)
	RerunFailed(ctx context.Context, in *RerunFailedProjectPipelineRequest, opts ...grpc.CallOption) (*RerunFailedProjectPipelineResponse, error)
	Cancel(ctx context.Context, in *CancelProjectPipelineRequest, opts ...grpc.CallOption) (*CancelProjectPipelineResponse, error)
	OneClickCreate(ctx context.Context, in *OneClickCreateProjectPipelineRequest, opts ...grpc.CallOption) (*OneClickCreateProjectPipelineResponse, error)
	BatchCreateByGittarPushHook(ctx context.Context, in *GittarPushPayloadEvent, opts ...grpc.CallOption) (*BatchCreateProjectPipelineResponse, error)
}

type projectPipelineServiceClient struct {
	cc grpc1.ClientConnInterface
}

func NewProjectPipelineServiceClient(cc grpc1.ClientConnInterface) ProjectPipelineServiceClient {
	return &projectPipelineServiceClient{cc}
}

func (c *projectPipelineServiceClient) Create(ctx context.Context, in *CreateProjectPipelineRequest, opts ...grpc.CallOption) (*CreateProjectPipelineResponse, error) {
	out := new(CreateProjectPipelineResponse)
	err := c.cc.Invoke(ctx, "/erda.dop.projectpipeline.ProjectPipelineService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectPipelineServiceClient) ListApp(ctx context.Context, in *ListAppRequest, opts ...grpc.CallOption) (*ListAppResponse, error) {
	out := new(ListAppResponse)
	err := c.cc.Invoke(ctx, "/erda.dop.projectpipeline.ProjectPipelineService/ListApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectPipelineServiceClient) ListPipelineYml(ctx context.Context, in *ListAppPipelineYmlRequest, opts ...grpc.CallOption) (*ListAppPipelineYmlResponse, error) {
	out := new(ListAppPipelineYmlResponse)
	err := c.cc.Invoke(ctx, "/erda.dop.projectpipeline.ProjectPipelineService/ListPipelineYml", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectPipelineServiceClient) CreateNamePreCheck(ctx context.Context, in *CreateProjectPipelineNamePreCheckRequest, opts ...grpc.CallOption) (*CreateProjectPipelineNamePreCheckResponse, error) {
	out := new(CreateProjectPipelineNamePreCheckResponse)
	err := c.cc.Invoke(ctx, "/erda.dop.projectpipeline.ProjectPipelineService/CreateNamePreCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectPipelineServiceClient) CreateSourcePreCheck(ctx context.Context, in *CreateProjectPipelineSourcePreCheckRequest, opts ...grpc.CallOption) (*CreateProjectPipelineSourcePreCheckResponse, error) {
	out := new(CreateProjectPipelineSourcePreCheckResponse)
	err := c.cc.Invoke(ctx, "/erda.dop.projectpipeline.ProjectPipelineService/CreateSourcePreCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectPipelineServiceClient) ListPipelineCategory(ctx context.Context, in *ListPipelineCategoryRequest, opts ...grpc.CallOption) (*ListPipelineCategoryResponse, error) {
	out := new(ListPipelineCategoryResponse)
	err := c.cc.Invoke(ctx, "/erda.dop.projectpipeline.ProjectPipelineService/ListPipelineCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectPipelineServiceClient) Update(ctx context.Context, in *UpdateProjectPipelineRequest, opts ...grpc.CallOption) (*UpdateProjectPipelineResponse, error) {
	out := new(UpdateProjectPipelineResponse)
	err := c.cc.Invoke(ctx, "/erda.dop.projectpipeline.ProjectPipelineService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectPipelineServiceClient) Run(ctx context.Context, in *RunProjectPipelineRequest, opts ...grpc.CallOption) (*RunProjectPipelineResponse, error) {
	out := new(RunProjectPipelineResponse)
	err := c.cc.Invoke(ctx, "/erda.dop.projectpipeline.ProjectPipelineService/Run", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectPipelineServiceClient) Rerun(ctx context.Context, in *RerunProjectPipelineRequest, opts ...grpc.CallOption) (*RerunProjectPipelineResponse, error) {
	out := new(RerunProjectPipelineResponse)
	err := c.cc.Invoke(ctx, "/erda.dop.projectpipeline.ProjectPipelineService/Rerun", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectPipelineServiceClient) RerunFailed(ctx context.Context, in *RerunFailedProjectPipelineRequest, opts ...grpc.CallOption) (*RerunFailedProjectPipelineResponse, error) {
	out := new(RerunFailedProjectPipelineResponse)
	err := c.cc.Invoke(ctx, "/erda.dop.projectpipeline.ProjectPipelineService/RerunFailed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectPipelineServiceClient) Cancel(ctx context.Context, in *CancelProjectPipelineRequest, opts ...grpc.CallOption) (*CancelProjectPipelineResponse, error) {
	out := new(CancelProjectPipelineResponse)
	err := c.cc.Invoke(ctx, "/erda.dop.projectpipeline.ProjectPipelineService/Cancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectPipelineServiceClient) OneClickCreate(ctx context.Context, in *OneClickCreateProjectPipelineRequest, opts ...grpc.CallOption) (*OneClickCreateProjectPipelineResponse, error) {
	out := new(OneClickCreateProjectPipelineResponse)
	err := c.cc.Invoke(ctx, "/erda.dop.projectpipeline.ProjectPipelineService/OneClickCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectPipelineServiceClient) BatchCreateByGittarPushHook(ctx context.Context, in *GittarPushPayloadEvent, opts ...grpc.CallOption) (*BatchCreateProjectPipelineResponse, error) {
	out := new(BatchCreateProjectPipelineResponse)
	err := c.cc.Invoke(ctx, "/erda.dop.projectpipeline.ProjectPipelineService/BatchCreateByGittarPushHook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectPipelineServiceServer is the server API for ProjectPipelineService service.
// All implementations should embed UnimplementedProjectPipelineServiceServer
// for forward compatibility
type ProjectPipelineServiceServer interface {
	Create(context.Context, *CreateProjectPipelineRequest) (*CreateProjectPipelineResponse, error)
	ListApp(context.Context, *ListAppRequest) (*ListAppResponse, error)
	ListPipelineYml(context.Context, *ListAppPipelineYmlRequest) (*ListAppPipelineYmlResponse, error)
	CreateNamePreCheck(context.Context, *CreateProjectPipelineNamePreCheckRequest) (*CreateProjectPipelineNamePreCheckResponse, error)
	CreateSourcePreCheck(context.Context, *CreateProjectPipelineSourcePreCheckRequest) (*CreateProjectPipelineSourcePreCheckResponse, error)
	ListPipelineCategory(context.Context, *ListPipelineCategoryRequest) (*ListPipelineCategoryResponse, error)
	Update(context.Context, *UpdateProjectPipelineRequest) (*UpdateProjectPipelineResponse, error)
	Run(context.Context, *RunProjectPipelineRequest) (*RunProjectPipelineResponse, error)
	Rerun(context.Context, *RerunProjectPipelineRequest) (*RerunProjectPipelineResponse, error)
	RerunFailed(context.Context, *RerunFailedProjectPipelineRequest) (*RerunFailedProjectPipelineResponse, error)
	Cancel(context.Context, *CancelProjectPipelineRequest) (*CancelProjectPipelineResponse, error)
	OneClickCreate(context.Context, *OneClickCreateProjectPipelineRequest) (*OneClickCreateProjectPipelineResponse, error)
	BatchCreateByGittarPushHook(context.Context, *GittarPushPayloadEvent) (*BatchCreateProjectPipelineResponse, error)
}

// UnimplementedProjectPipelineServiceServer should be embedded to have forward compatible implementations.
type UnimplementedProjectPipelineServiceServer struct {
}

func (*UnimplementedProjectPipelineServiceServer) Create(context.Context, *CreateProjectPipelineRequest) (*CreateProjectPipelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedProjectPipelineServiceServer) ListApp(context.Context, *ListAppRequest) (*ListAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListApp not implemented")
}
func (*UnimplementedProjectPipelineServiceServer) ListPipelineYml(context.Context, *ListAppPipelineYmlRequest) (*ListAppPipelineYmlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPipelineYml not implemented")
}
func (*UnimplementedProjectPipelineServiceServer) CreateNamePreCheck(context.Context, *CreateProjectPipelineNamePreCheckRequest) (*CreateProjectPipelineNamePreCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNamePreCheck not implemented")
}
func (*UnimplementedProjectPipelineServiceServer) CreateSourcePreCheck(context.Context, *CreateProjectPipelineSourcePreCheckRequest) (*CreateProjectPipelineSourcePreCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSourcePreCheck not implemented")
}
func (*UnimplementedProjectPipelineServiceServer) ListPipelineCategory(context.Context, *ListPipelineCategoryRequest) (*ListPipelineCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPipelineCategory not implemented")
}
func (*UnimplementedProjectPipelineServiceServer) Update(context.Context, *UpdateProjectPipelineRequest) (*UpdateProjectPipelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedProjectPipelineServiceServer) Run(context.Context, *RunProjectPipelineRequest) (*RunProjectPipelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Run not implemented")
}
func (*UnimplementedProjectPipelineServiceServer) Rerun(context.Context, *RerunProjectPipelineRequest) (*RerunProjectPipelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rerun not implemented")
}
func (*UnimplementedProjectPipelineServiceServer) RerunFailed(context.Context, *RerunFailedProjectPipelineRequest) (*RerunFailedProjectPipelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RerunFailed not implemented")
}
func (*UnimplementedProjectPipelineServiceServer) Cancel(context.Context, *CancelProjectPipelineRequest) (*CancelProjectPipelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}
func (*UnimplementedProjectPipelineServiceServer) OneClickCreate(context.Context, *OneClickCreateProjectPipelineRequest) (*OneClickCreateProjectPipelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OneClickCreate not implemented")
}
func (*UnimplementedProjectPipelineServiceServer) BatchCreateByGittarPushHook(context.Context, *GittarPushPayloadEvent) (*BatchCreateProjectPipelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchCreateByGittarPushHook not implemented")
}

func RegisterProjectPipelineServiceServer(s grpc1.ServiceRegistrar, srv ProjectPipelineServiceServer, opts ...grpc1.HandleOption) {
	s.RegisterService(_get_ProjectPipelineService_serviceDesc(srv, opts...), srv)
}

var _ProjectPipelineService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "erda.dop.projectpipeline.ProjectPipelineService",
	HandlerType: (*ProjectPipelineServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "projectpipeline.proto",
}

func _get_ProjectPipelineService_serviceDesc(srv ProjectPipelineServiceServer, opts ...grpc1.HandleOption) *grpc.ServiceDesc {
	h := grpc1.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}

	_ProjectPipelineService_Create_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.Create(ctx, req.(*CreateProjectPipelineRequest))
	}
	var _ProjectPipelineService_Create_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ProjectPipelineService_Create_info = transport.NewServiceInfo("erda.dop.projectpipeline.ProjectPipelineService", "Create", srv)
		_ProjectPipelineService_Create_Handler = h.Interceptor(_ProjectPipelineService_Create_Handler)
	}

	_ProjectPipelineService_ListApp_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.ListApp(ctx, req.(*ListAppRequest))
	}
	var _ProjectPipelineService_ListApp_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ProjectPipelineService_ListApp_info = transport.NewServiceInfo("erda.dop.projectpipeline.ProjectPipelineService", "ListApp", srv)
		_ProjectPipelineService_ListApp_Handler = h.Interceptor(_ProjectPipelineService_ListApp_Handler)
	}

	_ProjectPipelineService_ListPipelineYml_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.ListPipelineYml(ctx, req.(*ListAppPipelineYmlRequest))
	}
	var _ProjectPipelineService_ListPipelineYml_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ProjectPipelineService_ListPipelineYml_info = transport.NewServiceInfo("erda.dop.projectpipeline.ProjectPipelineService", "ListPipelineYml", srv)
		_ProjectPipelineService_ListPipelineYml_Handler = h.Interceptor(_ProjectPipelineService_ListPipelineYml_Handler)
	}

	_ProjectPipelineService_CreateNamePreCheck_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.CreateNamePreCheck(ctx, req.(*CreateProjectPipelineNamePreCheckRequest))
	}
	var _ProjectPipelineService_CreateNamePreCheck_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ProjectPipelineService_CreateNamePreCheck_info = transport.NewServiceInfo("erda.dop.projectpipeline.ProjectPipelineService", "CreateNamePreCheck", srv)
		_ProjectPipelineService_CreateNamePreCheck_Handler = h.Interceptor(_ProjectPipelineService_CreateNamePreCheck_Handler)
	}

	_ProjectPipelineService_CreateSourcePreCheck_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.CreateSourcePreCheck(ctx, req.(*CreateProjectPipelineSourcePreCheckRequest))
	}
	var _ProjectPipelineService_CreateSourcePreCheck_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ProjectPipelineService_CreateSourcePreCheck_info = transport.NewServiceInfo("erda.dop.projectpipeline.ProjectPipelineService", "CreateSourcePreCheck", srv)
		_ProjectPipelineService_CreateSourcePreCheck_Handler = h.Interceptor(_ProjectPipelineService_CreateSourcePreCheck_Handler)
	}

	_ProjectPipelineService_ListPipelineCategory_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.ListPipelineCategory(ctx, req.(*ListPipelineCategoryRequest))
	}
	var _ProjectPipelineService_ListPipelineCategory_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ProjectPipelineService_ListPipelineCategory_info = transport.NewServiceInfo("erda.dop.projectpipeline.ProjectPipelineService", "ListPipelineCategory", srv)
		_ProjectPipelineService_ListPipelineCategory_Handler = h.Interceptor(_ProjectPipelineService_ListPipelineCategory_Handler)
	}

	_ProjectPipelineService_Update_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.Update(ctx, req.(*UpdateProjectPipelineRequest))
	}
	var _ProjectPipelineService_Update_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ProjectPipelineService_Update_info = transport.NewServiceInfo("erda.dop.projectpipeline.ProjectPipelineService", "Update", srv)
		_ProjectPipelineService_Update_Handler = h.Interceptor(_ProjectPipelineService_Update_Handler)
	}

	_ProjectPipelineService_Run_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.Run(ctx, req.(*RunProjectPipelineRequest))
	}
	var _ProjectPipelineService_Run_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ProjectPipelineService_Run_info = transport.NewServiceInfo("erda.dop.projectpipeline.ProjectPipelineService", "Run", srv)
		_ProjectPipelineService_Run_Handler = h.Interceptor(_ProjectPipelineService_Run_Handler)
	}

	_ProjectPipelineService_Rerun_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.Rerun(ctx, req.(*RerunProjectPipelineRequest))
	}
	var _ProjectPipelineService_Rerun_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ProjectPipelineService_Rerun_info = transport.NewServiceInfo("erda.dop.projectpipeline.ProjectPipelineService", "Rerun", srv)
		_ProjectPipelineService_Rerun_Handler = h.Interceptor(_ProjectPipelineService_Rerun_Handler)
	}

	_ProjectPipelineService_RerunFailed_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.RerunFailed(ctx, req.(*RerunFailedProjectPipelineRequest))
	}
	var _ProjectPipelineService_RerunFailed_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ProjectPipelineService_RerunFailed_info = transport.NewServiceInfo("erda.dop.projectpipeline.ProjectPipelineService", "RerunFailed", srv)
		_ProjectPipelineService_RerunFailed_Handler = h.Interceptor(_ProjectPipelineService_RerunFailed_Handler)
	}

	_ProjectPipelineService_Cancel_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.Cancel(ctx, req.(*CancelProjectPipelineRequest))
	}
	var _ProjectPipelineService_Cancel_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ProjectPipelineService_Cancel_info = transport.NewServiceInfo("erda.dop.projectpipeline.ProjectPipelineService", "Cancel", srv)
		_ProjectPipelineService_Cancel_Handler = h.Interceptor(_ProjectPipelineService_Cancel_Handler)
	}

	_ProjectPipelineService_OneClickCreate_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.OneClickCreate(ctx, req.(*OneClickCreateProjectPipelineRequest))
	}
	var _ProjectPipelineService_OneClickCreate_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ProjectPipelineService_OneClickCreate_info = transport.NewServiceInfo("erda.dop.projectpipeline.ProjectPipelineService", "OneClickCreate", srv)
		_ProjectPipelineService_OneClickCreate_Handler = h.Interceptor(_ProjectPipelineService_OneClickCreate_Handler)
	}

	_ProjectPipelineService_BatchCreateByGittarPushHook_Handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.BatchCreateByGittarPushHook(ctx, req.(*GittarPushPayloadEvent))
	}
	var _ProjectPipelineService_BatchCreateByGittarPushHook_info transport.ServiceInfo
	if h.Interceptor != nil {
		_ProjectPipelineService_BatchCreateByGittarPushHook_info = transport.NewServiceInfo("erda.dop.projectpipeline.ProjectPipelineService", "BatchCreateByGittarPushHook", srv)
		_ProjectPipelineService_BatchCreateByGittarPushHook_Handler = h.Interceptor(_ProjectPipelineService_BatchCreateByGittarPushHook_Handler)
	}

	var serviceDesc = _ProjectPipelineService_serviceDesc
	serviceDesc.Methods = []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(CreateProjectPipelineRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ProjectPipelineServiceServer).Create(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ProjectPipelineService_Create_info)
				}
				if interceptor == nil {
					return _ProjectPipelineService_Create_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.dop.projectpipeline.ProjectPipelineService/Create",
				}
				return interceptor(ctx, in, info, _ProjectPipelineService_Create_Handler)
			},
		},
		{
			MethodName: "ListApp",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(ListAppRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ProjectPipelineServiceServer).ListApp(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ProjectPipelineService_ListApp_info)
				}
				if interceptor == nil {
					return _ProjectPipelineService_ListApp_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.dop.projectpipeline.ProjectPipelineService/ListApp",
				}
				return interceptor(ctx, in, info, _ProjectPipelineService_ListApp_Handler)
			},
		},
		{
			MethodName: "ListPipelineYml",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(ListAppPipelineYmlRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ProjectPipelineServiceServer).ListPipelineYml(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ProjectPipelineService_ListPipelineYml_info)
				}
				if interceptor == nil {
					return _ProjectPipelineService_ListPipelineYml_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.dop.projectpipeline.ProjectPipelineService/ListPipelineYml",
				}
				return interceptor(ctx, in, info, _ProjectPipelineService_ListPipelineYml_Handler)
			},
		},
		{
			MethodName: "CreateNamePreCheck",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(CreateProjectPipelineNamePreCheckRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ProjectPipelineServiceServer).CreateNamePreCheck(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ProjectPipelineService_CreateNamePreCheck_info)
				}
				if interceptor == nil {
					return _ProjectPipelineService_CreateNamePreCheck_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.dop.projectpipeline.ProjectPipelineService/CreateNamePreCheck",
				}
				return interceptor(ctx, in, info, _ProjectPipelineService_CreateNamePreCheck_Handler)
			},
		},
		{
			MethodName: "CreateSourcePreCheck",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(CreateProjectPipelineSourcePreCheckRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ProjectPipelineServiceServer).CreateSourcePreCheck(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ProjectPipelineService_CreateSourcePreCheck_info)
				}
				if interceptor == nil {
					return _ProjectPipelineService_CreateSourcePreCheck_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.dop.projectpipeline.ProjectPipelineService/CreateSourcePreCheck",
				}
				return interceptor(ctx, in, info, _ProjectPipelineService_CreateSourcePreCheck_Handler)
			},
		},
		{
			MethodName: "ListPipelineCategory",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(ListPipelineCategoryRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ProjectPipelineServiceServer).ListPipelineCategory(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ProjectPipelineService_ListPipelineCategory_info)
				}
				if interceptor == nil {
					return _ProjectPipelineService_ListPipelineCategory_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.dop.projectpipeline.ProjectPipelineService/ListPipelineCategory",
				}
				return interceptor(ctx, in, info, _ProjectPipelineService_ListPipelineCategory_Handler)
			},
		},
		{
			MethodName: "Update",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(UpdateProjectPipelineRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ProjectPipelineServiceServer).Update(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ProjectPipelineService_Update_info)
				}
				if interceptor == nil {
					return _ProjectPipelineService_Update_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.dop.projectpipeline.ProjectPipelineService/Update",
				}
				return interceptor(ctx, in, info, _ProjectPipelineService_Update_Handler)
			},
		},
		{
			MethodName: "Run",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(RunProjectPipelineRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ProjectPipelineServiceServer).Run(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ProjectPipelineService_Run_info)
				}
				if interceptor == nil {
					return _ProjectPipelineService_Run_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.dop.projectpipeline.ProjectPipelineService/Run",
				}
				return interceptor(ctx, in, info, _ProjectPipelineService_Run_Handler)
			},
		},
		{
			MethodName: "Rerun",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(RerunProjectPipelineRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ProjectPipelineServiceServer).Rerun(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ProjectPipelineService_Rerun_info)
				}
				if interceptor == nil {
					return _ProjectPipelineService_Rerun_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.dop.projectpipeline.ProjectPipelineService/Rerun",
				}
				return interceptor(ctx, in, info, _ProjectPipelineService_Rerun_Handler)
			},
		},
		{
			MethodName: "RerunFailed",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(RerunFailedProjectPipelineRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ProjectPipelineServiceServer).RerunFailed(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ProjectPipelineService_RerunFailed_info)
				}
				if interceptor == nil {
					return _ProjectPipelineService_RerunFailed_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.dop.projectpipeline.ProjectPipelineService/RerunFailed",
				}
				return interceptor(ctx, in, info, _ProjectPipelineService_RerunFailed_Handler)
			},
		},
		{
			MethodName: "Cancel",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(CancelProjectPipelineRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ProjectPipelineServiceServer).Cancel(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ProjectPipelineService_Cancel_info)
				}
				if interceptor == nil {
					return _ProjectPipelineService_Cancel_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.dop.projectpipeline.ProjectPipelineService/Cancel",
				}
				return interceptor(ctx, in, info, _ProjectPipelineService_Cancel_Handler)
			},
		},
		{
			MethodName: "OneClickCreate",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(OneClickCreateProjectPipelineRequest)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ProjectPipelineServiceServer).OneClickCreate(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ProjectPipelineService_OneClickCreate_info)
				}
				if interceptor == nil {
					return _ProjectPipelineService_OneClickCreate_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.dop.projectpipeline.ProjectPipelineService/OneClickCreate",
				}
				return interceptor(ctx, in, info, _ProjectPipelineService_OneClickCreate_Handler)
			},
		},
		{
			MethodName: "BatchCreateByGittarPushHook",
			Handler: func(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
				in := new(GittarPushPayloadEvent)
				if err := dec(in); err != nil {
					return nil, err
				}
				if interceptor == nil && h.Interceptor == nil {
					return srv.(ProjectPipelineServiceServer).BatchCreateByGittarPushHook(ctx, in)
				}
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, _ProjectPipelineService_BatchCreateByGittarPushHook_info)
				}
				if interceptor == nil {
					return _ProjectPipelineService_BatchCreateByGittarPushHook_Handler(ctx, in)
				}
				info := &grpc.UnaryServerInfo{
					Server:     srv,
					FullMethod: "/erda.dop.projectpipeline.ProjectPipelineService/BatchCreateByGittarPushHook",
				}
				return interceptor(ctx, in, info, _ProjectPipelineService_BatchCreateByGittarPushHook_Handler)
			},
		},
	}
	return &serviceDesc
}
