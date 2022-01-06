// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: configcenter.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/msp/configcenter/pb"
	grpc1 "google.golang.org/grpc"
)

var dependencies = []string{
	"grpc-client@erda.msp.configcenter",
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
	clientsType                   = reflect.TypeOf((*Client)(nil)).Elem()
	configCenterServiceClientType = reflect.TypeOf((*pb.ConfigCenterServiceClient)(nil)).Elem()
	configCenterServiceServerType = reflect.TypeOf((*pb.ConfigCenterServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.msp.configcenter-client":
		return p.client
	case "erda.msp.configcenter.ConfigCenterService":
		return &configCenterServiceWrapper{client: p.client.ConfigCenterService(), opts: opts}
	case "erda.msp.configcenter.ConfigCenterService.client":
		return p.client.ConfigCenterService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case configCenterServiceClientType:
		return p.client.ConfigCenterService()
	case configCenterServiceServerType:
		return &configCenterServiceWrapper{client: p.client.ConfigCenterService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.msp.configcenter-client", &servicehub.Spec{
		Services: []string{
			"erda.msp.configcenter.ConfigCenterService",
			"erda.msp.configcenter.ConfigCenterService.client",
			"erda.msp.configcenter-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			configCenterServiceClientType,
			// server types
			configCenterServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
