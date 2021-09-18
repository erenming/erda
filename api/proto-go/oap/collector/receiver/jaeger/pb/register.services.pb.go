// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Sources: jaeger.proto

package pb

import (
	reflect "reflect"

	transport "github.com/erda-project/erda-infra/pkg/transport"
)

// RegisterJaegerServiceImp jaeger.proto
func RegisterJaegerServiceImp(regester transport.Register, srv JaegerServiceServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterJaegerServiceHandler(regester, JaegerServiceHandler(srv), _ops.HTTP...)
	RegisterJaegerServiceServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.oap.collector.receiver.jaeger.JaegerService",
	)
}

var (
	jaegerServiceClientType  = reflect.TypeOf((*JaegerServiceClient)(nil)).Elem()
	jaegerServiceServerType  = reflect.TypeOf((*JaegerServiceServer)(nil)).Elem()
	jaegerServiceHandlerType = reflect.TypeOf((*JaegerServiceHandler)(nil)).Elem()
)

// JaegerServiceClientType .
func JaegerServiceClientType() reflect.Type { return jaegerServiceClientType }

// JaegerServiceServerType .
func JaegerServiceServerType() reflect.Type { return jaegerServiceServerType }

// JaegerServiceHandlerType .
func JaegerServiceHandlerType() reflect.Type { return jaegerServiceHandlerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		jaegerServiceClientType,
		// server types
		jaegerServiceServerType,
		// handler types
		jaegerServiceHandlerType,
	}
}
