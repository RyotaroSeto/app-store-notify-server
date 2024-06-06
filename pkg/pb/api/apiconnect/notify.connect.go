// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/notify.proto

package apiconnect

import (
	api "app-store-notify-server/pkg/pb/api"
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// NotifyServiceName is the fully-qualified name of the NotifyService service.
	NotifyServiceName = "api.NotifyService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// NotifyServiceNotifyProcedure is the fully-qualified name of the NotifyService's Notify RPC.
	NotifyServiceNotifyProcedure = "/api.NotifyService/Notify"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	notifyServiceServiceDescriptor      = api.File_api_notify_proto.Services().ByName("NotifyService")
	notifyServiceNotifyMethodDescriptor = notifyServiceServiceDescriptor.Methods().ByName("Notify")
)

// NotifyServiceClient is a client for the api.NotifyService service.
type NotifyServiceClient interface {
	Notify(context.Context, *connect.Request[api.NotifyRequest]) (*connect.Response[api.NotifyResponse], error)
}

// NewNotifyServiceClient constructs a client for the api.NotifyService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewNotifyServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) NotifyServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &notifyServiceClient{
		notify: connect.NewClient[api.NotifyRequest, api.NotifyResponse](
			httpClient,
			baseURL+NotifyServiceNotifyProcedure,
			connect.WithSchema(notifyServiceNotifyMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// notifyServiceClient implements NotifyServiceClient.
type notifyServiceClient struct {
	notify *connect.Client[api.NotifyRequest, api.NotifyResponse]
}

// Notify calls api.NotifyService.Notify.
func (c *notifyServiceClient) Notify(ctx context.Context, req *connect.Request[api.NotifyRequest]) (*connect.Response[api.NotifyResponse], error) {
	return c.notify.CallUnary(ctx, req)
}

// NotifyServiceHandler is an implementation of the api.NotifyService service.
type NotifyServiceHandler interface {
	Notify(context.Context, *connect.Request[api.NotifyRequest]) (*connect.Response[api.NotifyResponse], error)
}

// NewNotifyServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewNotifyServiceHandler(svc NotifyServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	notifyServiceNotifyHandler := connect.NewUnaryHandler(
		NotifyServiceNotifyProcedure,
		svc.Notify,
		connect.WithSchema(notifyServiceNotifyMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/api.NotifyService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case NotifyServiceNotifyProcedure:
			notifyServiceNotifyHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedNotifyServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedNotifyServiceHandler struct{}

func (UnimplementedNotifyServiceHandler) Notify(context.Context, *connect.Request[api.NotifyRequest]) (*connect.Response[api.NotifyResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("api.NotifyService.Notify is not implemented"))
}
