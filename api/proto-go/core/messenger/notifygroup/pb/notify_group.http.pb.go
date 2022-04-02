// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: notify_group.proto

package pb

import (
	context "context"
	http1 "net/http"
	strconv "strconv"
	strings "strings"

	transport "github.com/erda-project/erda-infra/pkg/transport"
	http "github.com/erda-project/erda-infra/pkg/transport/http"
	httprule "github.com/erda-project/erda-infra/pkg/transport/http/httprule"
	runtime "github.com/erda-project/erda-infra/pkg/transport/http/runtime"
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/transport/http" package it is being compiled against.
const _ = http.SupportPackageIsVersion1

// NotifyGroupServiceHandler is the server API for NotifyGroupService service.
type NotifyGroupServiceHandler interface {
	// POST /api/notify-groups
	CreateNotifyGroup(context.Context, *CreateNotifyGroupRequest) (*CreateNotifyGroupResponse, error)
	// GET /api/notify-groups
	QueryNotifyGroup(context.Context, *QueryNotifyGroupRequest) (*QueryNotifyGroupResponse, error)
	// GET /api/notify-groups/{groupID}
	GetNotifyGroup(context.Context, *GetNotifyGroupRequest) (*GetNotifyGroupResponse, error)
	// PUT /api/notify-groups/{groupID}
	UpdateNotifyGroup(context.Context, *UpdateNotifyGroupRequest) (*UpdateNotifyGroupResponse, error)
	// GET /api/notify-groups/{groupID}/detail
	GetNotifyGroupDetail(context.Context, *GetNotifyGroupDetailRequest) (*GetNotifyGroupDetailResponse, error)
	// DELETE /api/notify-groups/{groupID}
	DeleteNotifyGroup(context.Context, *DeleteNotifyGroupRequest) (*DeleteNotifyGroupResponse, error)
	// GET /api/notify-groups/actions/batch-get
	BatchGetNotifyGroup(context.Context, *BatchGetNotifyGroupRequest) (*BatchGetNotifyGroupResponse, error)
}

// RegisterNotifyGroupServiceHandler register NotifyGroupServiceHandler to http.Router.
func RegisterNotifyGroupServiceHandler(r http.Router, srv NotifyGroupServiceHandler, opts ...http.HandleOption) {
	h := http.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}
	encodeFunc := func(fn func(http1.ResponseWriter, *http1.Request) (interface{}, error)) http.HandlerFunc {
		handler := func(w http1.ResponseWriter, r *http1.Request) {
			out, err := fn(w, r)
			if err != nil {
				h.Error(w, r, err)
				return
			}
			if err := h.Encode(w, r, out); err != nil {
				h.Error(w, r, err)
			}
		}
		if h.HTTPInterceptor != nil {
			handler = h.HTTPInterceptor(handler)
		}
		return handler
	}

	add_CreateNotifyGroup := func(method, path string, fn func(context.Context, *CreateNotifyGroupRequest) (*CreateNotifyGroupResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*CreateNotifyGroupRequest))
		}
		var CreateNotifyGroup_info transport.ServiceInfo
		if h.Interceptor != nil {
			CreateNotifyGroup_info = transport.NewServiceInfo("erda.core.messenger.notifygroup.NotifyGroupService", "CreateNotifyGroup", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, CreateNotifyGroup_info)
				}
				r = r.WithContext(ctx)
				var in CreateNotifyGroupRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_QueryNotifyGroup := func(method, path string, fn func(context.Context, *QueryNotifyGroupRequest) (*QueryNotifyGroupResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*QueryNotifyGroupRequest))
		}
		var QueryNotifyGroup_info transport.ServiceInfo
		if h.Interceptor != nil {
			QueryNotifyGroup_info = transport.NewServiceInfo("erda.core.messenger.notifygroup.NotifyGroupService", "QueryNotifyGroup", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, QueryNotifyGroup_info)
				}
				r = r.WithContext(ctx)
				var in QueryNotifyGroupRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_GetNotifyGroup := func(method, path string, fn func(context.Context, *GetNotifyGroupRequest) (*GetNotifyGroupResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetNotifyGroupRequest))
		}
		var GetNotifyGroup_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetNotifyGroup_info = transport.NewServiceInfo("erda.core.messenger.notifygroup.NotifyGroupService", "GetNotifyGroup", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetNotifyGroup_info)
				}
				r = r.WithContext(ctx)
				var in GetNotifyGroupRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "groupID":
							val, err := strconv.ParseInt(val, 10, 64)
							if err != nil {
								return nil, err
							}
							in.GroupID = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_UpdateNotifyGroup := func(method, path string, fn func(context.Context, *UpdateNotifyGroupRequest) (*UpdateNotifyGroupResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*UpdateNotifyGroupRequest))
		}
		var UpdateNotifyGroup_info transport.ServiceInfo
		if h.Interceptor != nil {
			UpdateNotifyGroup_info = transport.NewServiceInfo("erda.core.messenger.notifygroup.NotifyGroupService", "UpdateNotifyGroup", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, UpdateNotifyGroup_info)
				}
				r = r.WithContext(ctx)
				var in UpdateNotifyGroupRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "groupID":
							val, err := strconv.ParseInt(val, 10, 64)
							if err != nil {
								return nil, err
							}
							in.GroupID = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_GetNotifyGroupDetail := func(method, path string, fn func(context.Context, *GetNotifyGroupDetailRequest) (*GetNotifyGroupDetailResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*GetNotifyGroupDetailRequest))
		}
		var GetNotifyGroupDetail_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetNotifyGroupDetail_info = transport.NewServiceInfo("erda.core.messenger.notifygroup.NotifyGroupService", "GetNotifyGroupDetail", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetNotifyGroupDetail_info)
				}
				r = r.WithContext(ctx)
				var in GetNotifyGroupDetailRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "groupID":
							val, err := strconv.ParseInt(val, 10, 64)
							if err != nil {
								return nil, err
							}
							in.GroupID = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_DeleteNotifyGroup := func(method, path string, fn func(context.Context, *DeleteNotifyGroupRequest) (*DeleteNotifyGroupResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*DeleteNotifyGroupRequest))
		}
		var DeleteNotifyGroup_info transport.ServiceInfo
		if h.Interceptor != nil {
			DeleteNotifyGroup_info = transport.NewServiceInfo("erda.core.messenger.notifygroup.NotifyGroupService", "DeleteNotifyGroup", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, DeleteNotifyGroup_info)
				}
				r = r.WithContext(ctx)
				var in DeleteNotifyGroupRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "groupID":
							val, err := strconv.ParseInt(val, 10, 64)
							if err != nil {
								return nil, err
							}
							in.GroupID = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_BatchGetNotifyGroup := func(method, path string, fn func(context.Context, *BatchGetNotifyGroupRequest) (*BatchGetNotifyGroupResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*BatchGetNotifyGroupRequest))
		}
		var BatchGetNotifyGroup_info transport.ServiceInfo
		if h.Interceptor != nil {
			BatchGetNotifyGroup_info = transport.NewServiceInfo("erda.core.messenger.notifygroup.NotifyGroupService", "BatchGetNotifyGroup", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, BatchGetNotifyGroup_info)
				}
				r = r.WithContext(ctx)
				var in BatchGetNotifyGroupRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_CreateNotifyGroup("POST", "/api/notify-groups", srv.CreateNotifyGroup)
	add_QueryNotifyGroup("GET", "/api/notify-groups", srv.QueryNotifyGroup)
	add_GetNotifyGroup("GET", "/api/notify-groups/{groupID}", srv.GetNotifyGroup)
	add_UpdateNotifyGroup("PUT", "/api/notify-groups/{groupID}", srv.UpdateNotifyGroup)
	add_GetNotifyGroupDetail("GET", "/api/notify-groups/{groupID}/detail", srv.GetNotifyGroupDetail)
	add_DeleteNotifyGroup("DELETE", "/api/notify-groups/{groupID}", srv.DeleteNotifyGroup)
	add_BatchGetNotifyGroup("GET", "/api/notify-groups/actions/batch-get", srv.BatchGetNotifyGroup)
}
