// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: ingester/v1/ingester.proto

package ingesterv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v11 "github.com/grafana/fire/pkg/gen/ingester/v1"
	v1 "github.com/grafana/fire/pkg/gen/push/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// IngesterServiceName is the fully-qualified name of the IngesterService service.
	IngesterServiceName = "ingester.v1.IngesterService"
)

// IngesterServiceClient is a client for the ingester.v1.IngesterService service.
type IngesterServiceClient interface {
	Push(context.Context, *connect_go.Request[v1.PushRequest]) (*connect_go.Response[v1.PushResponse], error)
	LabelValues(context.Context, *connect_go.Request[v11.LabelValuesRequest]) (*connect_go.Response[v11.LabelValuesResponse], error)
	ProfileTypes(context.Context, *connect_go.Request[v11.ProfileTypesRequest]) (*connect_go.Response[v11.ProfileTypesResponse], error)
	Series(context.Context, *connect_go.Request[v11.SeriesRequest]) (*connect_go.Response[v11.SeriesResponse], error)
	Flush(context.Context, *connect_go.Request[v11.FlushRequest]) (*connect_go.Response[v11.FlushResponse], error)
	// Todo(ctovena) we might want to batch stream profiles & symbolization instead of sending them all at once.
	// but this requires to ensure we have correct timestamp and labels ordering.
	SelectProfiles(context.Context, *connect_go.Request[v11.SelectProfilesRequest]) (*connect_go.Response[v11.SelectProfilesResponse], error)
}

// NewIngesterServiceClient constructs a client for the ingester.v1.IngesterService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewIngesterServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) IngesterServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &ingesterServiceClient{
		push: connect_go.NewClient[v1.PushRequest, v1.PushResponse](
			httpClient,
			baseURL+"/ingester.v1.IngesterService/Push",
			opts...,
		),
		labelValues: connect_go.NewClient[v11.LabelValuesRequest, v11.LabelValuesResponse](
			httpClient,
			baseURL+"/ingester.v1.IngesterService/LabelValues",
			opts...,
		),
		profileTypes: connect_go.NewClient[v11.ProfileTypesRequest, v11.ProfileTypesResponse](
			httpClient,
			baseURL+"/ingester.v1.IngesterService/ProfileTypes",
			opts...,
		),
		series: connect_go.NewClient[v11.SeriesRequest, v11.SeriesResponse](
			httpClient,
			baseURL+"/ingester.v1.IngesterService/Series",
			opts...,
		),
		flush: connect_go.NewClient[v11.FlushRequest, v11.FlushResponse](
			httpClient,
			baseURL+"/ingester.v1.IngesterService/Flush",
			opts...,
		),
		selectProfiles: connect_go.NewClient[v11.SelectProfilesRequest, v11.SelectProfilesResponse](
			httpClient,
			baseURL+"/ingester.v1.IngesterService/SelectProfiles",
			opts...,
		),
	}
}

// ingesterServiceClient implements IngesterServiceClient.
type ingesterServiceClient struct {
	push           *connect_go.Client[v1.PushRequest, v1.PushResponse]
	labelValues    *connect_go.Client[v11.LabelValuesRequest, v11.LabelValuesResponse]
	profileTypes   *connect_go.Client[v11.ProfileTypesRequest, v11.ProfileTypesResponse]
	series         *connect_go.Client[v11.SeriesRequest, v11.SeriesResponse]
	flush          *connect_go.Client[v11.FlushRequest, v11.FlushResponse]
	selectProfiles *connect_go.Client[v11.SelectProfilesRequest, v11.SelectProfilesResponse]
}

// Push calls ingester.v1.IngesterService.Push.
func (c *ingesterServiceClient) Push(ctx context.Context, req *connect_go.Request[v1.PushRequest]) (*connect_go.Response[v1.PushResponse], error) {
	return c.push.CallUnary(ctx, req)
}

// LabelValues calls ingester.v1.IngesterService.LabelValues.
func (c *ingesterServiceClient) LabelValues(ctx context.Context, req *connect_go.Request[v11.LabelValuesRequest]) (*connect_go.Response[v11.LabelValuesResponse], error) {
	return c.labelValues.CallUnary(ctx, req)
}

// ProfileTypes calls ingester.v1.IngesterService.ProfileTypes.
func (c *ingesterServiceClient) ProfileTypes(ctx context.Context, req *connect_go.Request[v11.ProfileTypesRequest]) (*connect_go.Response[v11.ProfileTypesResponse], error) {
	return c.profileTypes.CallUnary(ctx, req)
}

// Series calls ingester.v1.IngesterService.Series.
func (c *ingesterServiceClient) Series(ctx context.Context, req *connect_go.Request[v11.SeriesRequest]) (*connect_go.Response[v11.SeriesResponse], error) {
	return c.series.CallUnary(ctx, req)
}

// Flush calls ingester.v1.IngesterService.Flush.
func (c *ingesterServiceClient) Flush(ctx context.Context, req *connect_go.Request[v11.FlushRequest]) (*connect_go.Response[v11.FlushResponse], error) {
	return c.flush.CallUnary(ctx, req)
}

// SelectProfiles calls ingester.v1.IngesterService.SelectProfiles.
func (c *ingesterServiceClient) SelectProfiles(ctx context.Context, req *connect_go.Request[v11.SelectProfilesRequest]) (*connect_go.Response[v11.SelectProfilesResponse], error) {
	return c.selectProfiles.CallUnary(ctx, req)
}

// IngesterServiceHandler is an implementation of the ingester.v1.IngesterService service.
type IngesterServiceHandler interface {
	Push(context.Context, *connect_go.Request[v1.PushRequest]) (*connect_go.Response[v1.PushResponse], error)
	LabelValues(context.Context, *connect_go.Request[v11.LabelValuesRequest]) (*connect_go.Response[v11.LabelValuesResponse], error)
	ProfileTypes(context.Context, *connect_go.Request[v11.ProfileTypesRequest]) (*connect_go.Response[v11.ProfileTypesResponse], error)
	Series(context.Context, *connect_go.Request[v11.SeriesRequest]) (*connect_go.Response[v11.SeriesResponse], error)
	Flush(context.Context, *connect_go.Request[v11.FlushRequest]) (*connect_go.Response[v11.FlushResponse], error)
	// Todo(ctovena) we might want to batch stream profiles & symbolization instead of sending them all at once.
	// but this requires to ensure we have correct timestamp and labels ordering.
	SelectProfiles(context.Context, *connect_go.Request[v11.SelectProfilesRequest]) (*connect_go.Response[v11.SelectProfilesResponse], error)
}

// NewIngesterServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewIngesterServiceHandler(svc IngesterServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/ingester.v1.IngesterService/Push", connect_go.NewUnaryHandler(
		"/ingester.v1.IngesterService/Push",
		svc.Push,
		opts...,
	))
	mux.Handle("/ingester.v1.IngesterService/LabelValues", connect_go.NewUnaryHandler(
		"/ingester.v1.IngesterService/LabelValues",
		svc.LabelValues,
		opts...,
	))
	mux.Handle("/ingester.v1.IngesterService/ProfileTypes", connect_go.NewUnaryHandler(
		"/ingester.v1.IngesterService/ProfileTypes",
		svc.ProfileTypes,
		opts...,
	))
	mux.Handle("/ingester.v1.IngesterService/Series", connect_go.NewUnaryHandler(
		"/ingester.v1.IngesterService/Series",
		svc.Series,
		opts...,
	))
	mux.Handle("/ingester.v1.IngesterService/Flush", connect_go.NewUnaryHandler(
		"/ingester.v1.IngesterService/Flush",
		svc.Flush,
		opts...,
	))
	mux.Handle("/ingester.v1.IngesterService/SelectProfiles", connect_go.NewUnaryHandler(
		"/ingester.v1.IngesterService/SelectProfiles",
		svc.SelectProfiles,
		opts...,
	))
	return "/ingester.v1.IngesterService/", mux
}

// UnimplementedIngesterServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedIngesterServiceHandler struct{}

func (UnimplementedIngesterServiceHandler) Push(context.Context, *connect_go.Request[v1.PushRequest]) (*connect_go.Response[v1.PushResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ingester.v1.IngesterService.Push is not implemented"))
}

func (UnimplementedIngesterServiceHandler) LabelValues(context.Context, *connect_go.Request[v11.LabelValuesRequest]) (*connect_go.Response[v11.LabelValuesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ingester.v1.IngesterService.LabelValues is not implemented"))
}

func (UnimplementedIngesterServiceHandler) ProfileTypes(context.Context, *connect_go.Request[v11.ProfileTypesRequest]) (*connect_go.Response[v11.ProfileTypesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ingester.v1.IngesterService.ProfileTypes is not implemented"))
}

func (UnimplementedIngesterServiceHandler) Series(context.Context, *connect_go.Request[v11.SeriesRequest]) (*connect_go.Response[v11.SeriesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ingester.v1.IngesterService.Series is not implemented"))
}

func (UnimplementedIngesterServiceHandler) Flush(context.Context, *connect_go.Request[v11.FlushRequest]) (*connect_go.Response[v11.FlushResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ingester.v1.IngesterService.Flush is not implemented"))
}

func (UnimplementedIngesterServiceHandler) SelectProfiles(context.Context, *connect_go.Request[v11.SelectProfilesRequest]) (*connect_go.Response[v11.SelectProfilesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ingester.v1.IngesterService.SelectProfiles is not implemented"))
}
