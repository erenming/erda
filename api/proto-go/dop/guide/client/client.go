// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: guide.proto

package client

import (
	context "context"

	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/dop/guide/pb"
	grpc1 "google.golang.org/grpc"
)

// Client provide all service clients.
type Client interface {
	// GuideService guide.proto
	GuideService() pb.GuideServiceClient
}

// New create client
func New(cc grpc.ClientConnInterface) Client {
	return &serviceClients{
		guideService: pb.NewGuideServiceClient(cc),
	}
}

type serviceClients struct {
	guideService pb.GuideServiceClient
}

func (c *serviceClients) GuideService() pb.GuideServiceClient {
	return c.guideService
}

type guideServiceWrapper struct {
	client pb.GuideServiceClient
	opts   []grpc1.CallOption
}

func (s *guideServiceWrapper) CreateGuideByGittarPushHook(ctx context.Context, req *pb.GittarPushPayloadEvent) (*pb.CreateGuideResponse, error) {
	return s.client.CreateGuideByGittarPushHook(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *guideServiceWrapper) ListGuide(ctx context.Context, req *pb.ListGuideRequest) (*pb.ListGuideResponse, error) {
	return s.client.ListGuide(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *guideServiceWrapper) ProcessGuide(ctx context.Context, req *pb.ProcessGuideRequest) (*pb.ProcessGuideResponse, error) {
	return s.client.ProcessGuide(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *guideServiceWrapper) DeleteGuideByGittarPushHook(ctx context.Context, req *pb.GittarPushPayloadEvent) (*pb.DeleteGuideResponse, error) {
	return s.client.DeleteGuideByGittarPushHook(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}

func (s *guideServiceWrapper) CancelGuide(ctx context.Context, req *pb.CancelGuideRequest) (*pb.CancelGuideResponse, error) {
	return s.client.CancelGuide(ctx, req, append(grpc.CallOptionFromContext(ctx), s.opts...)...)
}
