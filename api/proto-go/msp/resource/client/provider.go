// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: resource.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/msp/resource/pb"
	grpc1 "google.golang.org/grpc"
)

var dependencies = []string{
	"grpc-client@erda.msp.resource",
	"grpc-client",
}

// +provider
type provider struct {
	client Client
}

func (p *provider) Init(ctx servicehub.Context) error {
	var conn grpc.ClientConnInterface
	for _, dep := range dependencies {
		c, ok := ctx.Service(dep).(grpc.ClientConnInterface)
		if ok {
			conn = c
			break
		}
	}
	if conn == nil {
		return fmt.Errorf("not found connector in (%s)", strings.Join(dependencies, ", "))
	}
	p.client = New(conn)
	return nil
}

var (
	clientsType               = reflect.TypeOf((*Client)(nil)).Elem()
	resourceServiceClientType = reflect.TypeOf((*pb.ResourceServiceClient)(nil)).Elem()
	resourceServiceServerType = reflect.TypeOf((*pb.ResourceServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.msp.resource-client":
		return p.client
	case "erda.msp.resource.ResourceService":
		return &resourceServiceWrapper{client: p.client.ResourceService(), opts: opts}
	case "erda.msp.resource.ResourceService.client":
		return p.client.ResourceService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case resourceServiceClientType:
		return p.client.ResourceService()
	case resourceServiceServerType:
		return &resourceServiceWrapper{client: p.client.ResourceService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.msp.resource-client", &servicehub.Spec{
		Services: []string{
			"erda.msp.resource.ResourceService",
			"erda.msp.resource-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			resourceServiceClientType,
			// server types
			resourceServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
