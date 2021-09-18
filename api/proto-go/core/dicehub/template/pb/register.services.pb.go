// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Sources: template.proto

package pb

import (
	reflect "reflect"

	transport "github.com/erda-project/erda-infra/pkg/transport"
)

// RegisterTemplateServiceImp template.proto
func RegisterTemplateServiceImp(regester transport.Register, srv TemplateServiceServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterTemplateServiceHandler(regester, TemplateServiceHandler(srv), _ops.HTTP...)
	RegisterTemplateServiceServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.core.dicehub.template.TemplateService",
	)
}

var (
	templateServiceClientType  = reflect.TypeOf((*TemplateServiceClient)(nil)).Elem()
	templateServiceServerType  = reflect.TypeOf((*TemplateServiceServer)(nil)).Elem()
	templateServiceHandlerType = reflect.TypeOf((*TemplateServiceHandler)(nil)).Elem()
)

// TemplateServiceClientType .
func TemplateServiceClientType() reflect.Type { return templateServiceClientType }

// TemplateServiceServerType .
func TemplateServiceServerType() reflect.Type { return templateServiceServerType }

// TemplateServiceHandlerType .
func TemplateServiceHandlerType() reflect.Type { return templateServiceHandlerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		templateServiceClientType,
		// server types
		templateServiceServerType,
		// handler types
		templateServiceHandlerType,
	}
}
