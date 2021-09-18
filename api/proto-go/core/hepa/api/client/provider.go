// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: micro_api.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/hepa/api/pb"
	grpc1 "google.golang.org/grpc"
)

var dependencies = []string{
	"grpc-client@erda.core.hepa.api",
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
	clientsType          = reflect.TypeOf((*Client)(nil)).Elem()
	apiServiceClientType = reflect.TypeOf((*pb.ApiServiceClient)(nil)).Elem()
	apiServiceServerType = reflect.TypeOf((*pb.ApiServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.core.hepa.api-client":
		return p.client
	case "erda.core.hepa.api.ApiService":
		return &apiServiceWrapper{client: p.client.ApiService(), opts: opts}
	case "erda.core.hepa.api.ApiService.client":
		return p.client.ApiService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case apiServiceClientType:
		return p.client.ApiService()
	case apiServiceServerType:
		return &apiServiceWrapper{client: p.client.ApiService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.core.hepa.api-client", &servicehub.Spec{
		Services: []string{
			"erda.core.hepa.api.ApiService",
			"erda.core.hepa.api-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			apiServiceClientType,
			// server types
			apiServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
