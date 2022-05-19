// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: flow.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/apps/devflow/flow/pb"
	grpc1 "google.golang.org/grpc"
)

var dependencies = []string{
	"grpc-client@erda.apps.devflow.flow",
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
	clientsType           = reflect.TypeOf((*Client)(nil)).Elem()
	flowServiceClientType = reflect.TypeOf((*pb.FlowServiceClient)(nil)).Elem()
	flowServiceServerType = reflect.TypeOf((*pb.FlowServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.apps.devflow.flow-client":
		return p.client
	case "erda.apps.devflow.flow.FlowService":
		return &flowServiceWrapper{client: p.client.FlowService(), opts: opts}
	case "erda.apps.devflow.flow.FlowService.client":
		return p.client.FlowService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case flowServiceClientType:
		return p.client.FlowService()
	case flowServiceServerType:
		return &flowServiceWrapper{client: p.client.FlowService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.apps.devflow.flow-client", &servicehub.Spec{
		Services: []string{
			"erda.apps.devflow.flow.FlowService",
			"erda.apps.devflow.flow.FlowService.client",
			"erda.apps.devflow.flow-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			flowServiceClientType,
			// server types
			flowServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
