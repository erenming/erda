// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: apm_alert.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/msp/apm/alert/pb"
	grpc1 "google.golang.org/grpc"
)

var dependencies = []string{
	"grpc-client@erda.msp.apm.alert",
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
	clientsType            = reflect.TypeOf((*Client)(nil)).Elem()
	alertServiceClientType = reflect.TypeOf((*pb.AlertServiceClient)(nil)).Elem()
	alertServiceServerType = reflect.TypeOf((*pb.AlertServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.msp.apm.alert-client":
		return p.client
	case "erda.msp.apm.alert.AlertService":
		return &alertServiceWrapper{client: p.client.AlertService(), opts: opts}
	case "erda.msp.apm.alert.AlertService.client":
		return p.client.AlertService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case alertServiceClientType:
		return p.client.AlertService()
	case alertServiceServerType:
		return &alertServiceWrapper{client: p.client.AlertService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.msp.apm.alert-client", &servicehub.Spec{
		Services: []string{
			"erda.msp.apm.alert.AlertService",
			"erda.msp.apm.alert.AlertService.client",
			"erda.msp.apm.alert-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			alertServiceClientType,
			// server types
			alertServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
