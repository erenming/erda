// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: service.proto

package client

import (
	context "context"

	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/msp/apm/service/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// ApmServiceService service.proto
	ApmServiceService() pb.ApmServiceServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		apmServiceService: pb.NewApmServiceServiceClient(cc),
	}
}

type serviceClients struct {
	apmServiceService pb.ApmServiceServiceClient
}

func (c *serviceClients) ApmServiceService() pb.ApmServiceServiceClient {
	return c.apmServiceService
}

type apmServiceServiceWrapper struct {
	client pb.ApmServiceServiceClient
	opts   []grpc1.CallOption
}

func (s *apmServiceServiceWrapper) GetServices(ctx context.Context, req *pb.GetServicesRequest) (*pb.GetServicesResponse, error) {
	return s.client.GetServices(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *apmServiceServiceWrapper) GetServiceAnalyzerOverview(ctx context.Context, req *pb.GetServiceAnalyzerOverviewRequest) (*pb.GetServiceAnalyzerOverviewResponse, error) {
	return s.client.GetServiceAnalyzerOverview(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *apmServiceServiceWrapper) GetServiceCount(ctx context.Context, req *pb.GetServiceCountRequest) (*pb.GetServiceCountResponse, error) {
	return s.client.GetServiceCount(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}
