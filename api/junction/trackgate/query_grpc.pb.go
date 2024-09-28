// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: junction/trackgate/query.proto

package trackgate

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Query_Params_FullMethodName               = "/junction.trackgate.Query/Params"
	Query_GetExtTrackStation_FullMethodName   = "/junction.trackgate.Query/GetExtTrackStation"
	Query_ListExtTrackStations_FullMethodName = "/junction.trackgate.Query/ListExtTrackStations"
	Query_RetrieveSchemaKey_FullMethodName    = "/junction.trackgate.Query/RetrieveSchemaKey"
	Query_ListSchemas_FullMethodName          = "/junction.trackgate.Query/ListSchemas"
	Query_GetTrackEngagement_FullMethodName   = "/junction.trackgate.Query/GetTrackEngagement"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of GetExtTrackStation items.
	GetExtTrackStation(ctx context.Context, in *QueryGetExtTrackStationRequest, opts ...grpc.CallOption) (*QueryGetExtTrackStationResponse, error)
	// Queries a list of ListExtTrackStations items.
	ListExtTrackStations(ctx context.Context, in *QueryListExtTrackStationsRequest, opts ...grpc.CallOption) (*QueryListExtTrackStationsResponse, error)
	// Queries a list of RetrieveSchemaKey items.
	RetrieveSchemaKey(ctx context.Context, in *QueryRetrieveSchemaKeyRequest, opts ...grpc.CallOption) (*QueryRetrieveSchemaKeyResponse, error)
	// Queries a list of ListSchemas items.
	ListSchemas(ctx context.Context, in *QueryListSchemasRequest, opts ...grpc.CallOption) (*QueryListSchemasResponse, error)
	// Queries a list of GetTrackEngagement items.
	GetTrackEngagement(ctx context.Context, in *QueryGetTrackEngagementRequest, opts ...grpc.CallOption) (*QueryGetTrackEngagementResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetExtTrackStation(ctx context.Context, in *QueryGetExtTrackStationRequest, opts ...grpc.CallOption) (*QueryGetExtTrackStationResponse, error) {
	out := new(QueryGetExtTrackStationResponse)
	err := c.cc.Invoke(ctx, Query_GetExtTrackStation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ListExtTrackStations(ctx context.Context, in *QueryListExtTrackStationsRequest, opts ...grpc.CallOption) (*QueryListExtTrackStationsResponse, error) {
	out := new(QueryListExtTrackStationsResponse)
	err := c.cc.Invoke(ctx, Query_ListExtTrackStations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) RetrieveSchemaKey(ctx context.Context, in *QueryRetrieveSchemaKeyRequest, opts ...grpc.CallOption) (*QueryRetrieveSchemaKeyResponse, error) {
	out := new(QueryRetrieveSchemaKeyResponse)
	err := c.cc.Invoke(ctx, Query_RetrieveSchemaKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ListSchemas(ctx context.Context, in *QueryListSchemasRequest, opts ...grpc.CallOption) (*QueryListSchemasResponse, error) {
	out := new(QueryListSchemasResponse)
	err := c.cc.Invoke(ctx, Query_ListSchemas_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetTrackEngagement(ctx context.Context, in *QueryGetTrackEngagementRequest, opts ...grpc.CallOption) (*QueryGetTrackEngagementResponse, error) {
	out := new(QueryGetTrackEngagementResponse)
	err := c.cc.Invoke(ctx, Query_GetTrackEngagement_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of GetExtTrackStation items.
	GetExtTrackStation(context.Context, *QueryGetExtTrackStationRequest) (*QueryGetExtTrackStationResponse, error)
	// Queries a list of ListExtTrackStations items.
	ListExtTrackStations(context.Context, *QueryListExtTrackStationsRequest) (*QueryListExtTrackStationsResponse, error)
	// Queries a list of RetrieveSchemaKey items.
	RetrieveSchemaKey(context.Context, *QueryRetrieveSchemaKeyRequest) (*QueryRetrieveSchemaKeyResponse, error)
	// Queries a list of ListSchemas items.
	ListSchemas(context.Context, *QueryListSchemasRequest) (*QueryListSchemasResponse, error)
	// Queries a list of GetTrackEngagement items.
	GetTrackEngagement(context.Context, *QueryGetTrackEngagementRequest) (*QueryGetTrackEngagementResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) GetExtTrackStation(context.Context, *QueryGetExtTrackStationRequest) (*QueryGetExtTrackStationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExtTrackStation not implemented")
}
func (UnimplementedQueryServer) ListExtTrackStations(context.Context, *QueryListExtTrackStationsRequest) (*QueryListExtTrackStationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListExtTrackStations not implemented")
}
func (UnimplementedQueryServer) RetrieveSchemaKey(context.Context, *QueryRetrieveSchemaKeyRequest) (*QueryRetrieveSchemaKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveSchemaKey not implemented")
}
func (UnimplementedQueryServer) ListSchemas(context.Context, *QueryListSchemasRequest) (*QueryListSchemasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSchemas not implemented")
}
func (UnimplementedQueryServer) GetTrackEngagement(context.Context, *QueryGetTrackEngagementRequest) (*QueryGetTrackEngagementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrackEngagement not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetExtTrackStation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetExtTrackStationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetExtTrackStation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetExtTrackStation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetExtTrackStation(ctx, req.(*QueryGetExtTrackStationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ListExtTrackStations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryListExtTrackStationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ListExtTrackStations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ListExtTrackStations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ListExtTrackStations(ctx, req.(*QueryListExtTrackStationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_RetrieveSchemaKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRetrieveSchemaKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).RetrieveSchemaKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_RetrieveSchemaKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).RetrieveSchemaKey(ctx, req.(*QueryRetrieveSchemaKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ListSchemas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryListSchemasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ListSchemas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ListSchemas_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ListSchemas(ctx, req.(*QueryListSchemasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetTrackEngagement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetTrackEngagementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetTrackEngagement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetTrackEngagement_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetTrackEngagement(ctx, req.(*QueryGetTrackEngagementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "junction.trackgate.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "GetExtTrackStation",
			Handler:    _Query_GetExtTrackStation_Handler,
		},
		{
			MethodName: "ListExtTrackStations",
			Handler:    _Query_ListExtTrackStations_Handler,
		},
		{
			MethodName: "RetrieveSchemaKey",
			Handler:    _Query_RetrieveSchemaKey_Handler,
		},
		{
			MethodName: "ListSchemas",
			Handler:    _Query_ListSchemas_Handler,
		},
		{
			MethodName: "GetTrackEngagement",
			Handler:    _Query_GetTrackEngagement_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "junction/trackgate/query.proto",
}
