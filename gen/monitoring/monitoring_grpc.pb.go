// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: monitoring/monitoring.proto

package monitoring

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

// MonitoringClient is the client API for Monitoring service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MonitoringClient interface {
	InsertNode(ctx context.Context, in *InsertNodeReq, opts ...grpc.CallOption) (*NodeRecord, error)
	GetNode(ctx context.Context, in *GetNodeReq, opts ...grpc.CallOption) (*NodeRecord, error)
	ListRuns(ctx context.Context, in *ListRunsReq, opts ...grpc.CallOption) (*ListRunsRes, error)
	GetRun(ctx context.Context, in *GetRunReq, opts ...grpc.CallOption) (*GetRunRes, error)
	GetChart(ctx context.Context, in *ListRunsReq, opts ...grpc.CallOption) (*GetChartRes, error)
}

type monitoringClient struct {
	cc grpc.ClientConnInterface
}

func NewMonitoringClient(cc grpc.ClientConnInterface) MonitoringClient {
	return &monitoringClient{cc}
}

func (c *monitoringClient) InsertNode(ctx context.Context, in *InsertNodeReq, opts ...grpc.CallOption) (*NodeRecord, error) {
	out := new(NodeRecord)
	err := c.cc.Invoke(ctx, "/monitoring.Monitoring/InsertNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitoringClient) GetNode(ctx context.Context, in *GetNodeReq, opts ...grpc.CallOption) (*NodeRecord, error) {
	out := new(NodeRecord)
	err := c.cc.Invoke(ctx, "/monitoring.Monitoring/GetNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitoringClient) ListRuns(ctx context.Context, in *ListRunsReq, opts ...grpc.CallOption) (*ListRunsRes, error) {
	out := new(ListRunsRes)
	err := c.cc.Invoke(ctx, "/monitoring.Monitoring/ListRuns", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitoringClient) GetRun(ctx context.Context, in *GetRunReq, opts ...grpc.CallOption) (*GetRunRes, error) {
	out := new(GetRunRes)
	err := c.cc.Invoke(ctx, "/monitoring.Monitoring/GetRun", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitoringClient) GetChart(ctx context.Context, in *ListRunsReq, opts ...grpc.CallOption) (*GetChartRes, error) {
	out := new(GetChartRes)
	err := c.cc.Invoke(ctx, "/monitoring.Monitoring/GetChart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MonitoringServer is the server API for Monitoring service.
// All implementations must embed UnimplementedMonitoringServer
// for forward compatibility
type MonitoringServer interface {
	InsertNode(context.Context, *InsertNodeReq) (*NodeRecord, error)
	GetNode(context.Context, *GetNodeReq) (*NodeRecord, error)
	ListRuns(context.Context, *ListRunsReq) (*ListRunsRes, error)
	GetRun(context.Context, *GetRunReq) (*GetRunRes, error)
	GetChart(context.Context, *ListRunsReq) (*GetChartRes, error)
	mustEmbedUnimplementedMonitoringServer()
}

// UnimplementedMonitoringServer must be embedded to have forward compatible implementations.
type UnimplementedMonitoringServer struct {
}

func (UnimplementedMonitoringServer) InsertNode(context.Context, *InsertNodeReq) (*NodeRecord, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertNode not implemented")
}
func (UnimplementedMonitoringServer) GetNode(context.Context, *GetNodeReq) (*NodeRecord, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNode not implemented")
}
func (UnimplementedMonitoringServer) ListRuns(context.Context, *ListRunsReq) (*ListRunsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRuns not implemented")
}
func (UnimplementedMonitoringServer) GetRun(context.Context, *GetRunReq) (*GetRunRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRun not implemented")
}
func (UnimplementedMonitoringServer) GetChart(context.Context, *ListRunsReq) (*GetChartRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChart not implemented")
}
func (UnimplementedMonitoringServer) mustEmbedUnimplementedMonitoringServer() {}

// UnsafeMonitoringServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MonitoringServer will
// result in compilation errors.
type UnsafeMonitoringServer interface {
	mustEmbedUnimplementedMonitoringServer()
}

func RegisterMonitoringServer(s grpc.ServiceRegistrar, srv MonitoringServer) {
	s.RegisterService(&Monitoring_ServiceDesc, srv)
}

func _Monitoring_InsertNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertNodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringServer).InsertNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitoring.Monitoring/InsertNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringServer).InsertNode(ctx, req.(*InsertNodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitoring_GetNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringServer).GetNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitoring.Monitoring/GetNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringServer).GetNode(ctx, req.(*GetNodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitoring_ListRuns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRunsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringServer).ListRuns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitoring.Monitoring/ListRuns",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringServer).ListRuns(ctx, req.(*ListRunsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitoring_GetRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRunReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringServer).GetRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitoring.Monitoring/GetRun",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringServer).GetRun(ctx, req.(*GetRunReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Monitoring_GetChart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRunsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitoringServer).GetChart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitoring.Monitoring/GetChart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitoringServer).GetChart(ctx, req.(*ListRunsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Monitoring_ServiceDesc is the grpc.ServiceDesc for Monitoring service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Monitoring_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "monitoring.Monitoring",
	HandlerType: (*MonitoringServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertNode",
			Handler:    _Monitoring_InsertNode_Handler,
		},
		{
			MethodName: "GetNode",
			Handler:    _Monitoring_GetNode_Handler,
		},
		{
			MethodName: "ListRuns",
			Handler:    _Monitoring_ListRuns_Handler,
		},
		{
			MethodName: "GetRun",
			Handler:    _Monitoring_GetRun_Handler,
		},
		{
			MethodName: "GetChart",
			Handler:    _Monitoring_GetChart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "monitoring/monitoring.proto",
}

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	QueryPublishedWorkflows(ctx context.Context, in *QueryReq, opts ...grpc.CallOption) (*QueryRes, error)
	CreateQuery(ctx context.Context, in *SavedQueryReq, opts ...grpc.CallOption) (*SavedQueryRecord, error)
	UpdateQuery(ctx context.Context, in *SavedQueryReq, opts ...grpc.CallOption) (*SavedQueryRecord, error)
	GetQueries(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SavedQueryList, error)
	DeleteQuery(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*IdRequest, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) QueryPublishedWorkflows(ctx context.Context, in *QueryReq, opts ...grpc.CallOption) (*QueryRes, error) {
	out := new(QueryRes)
	err := c.cc.Invoke(ctx, "/monitoring.Query/QueryPublishedWorkflows", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CreateQuery(ctx context.Context, in *SavedQueryReq, opts ...grpc.CallOption) (*SavedQueryRecord, error) {
	out := new(SavedQueryRecord)
	err := c.cc.Invoke(ctx, "/monitoring.Query/CreateQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) UpdateQuery(ctx context.Context, in *SavedQueryReq, opts ...grpc.CallOption) (*SavedQueryRecord, error) {
	out := new(SavedQueryRecord)
	err := c.cc.Invoke(ctx, "/monitoring.Query/UpdateQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetQueries(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SavedQueryList, error) {
	out := new(SavedQueryList)
	err := c.cc.Invoke(ctx, "/monitoring.Query/GetQueries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DeleteQuery(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*IdRequest, error) {
	out := new(IdRequest)
	err := c.cc.Invoke(ctx, "/monitoring.Query/DeleteQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	QueryPublishedWorkflows(context.Context, *QueryReq) (*QueryRes, error)
	CreateQuery(context.Context, *SavedQueryReq) (*SavedQueryRecord, error)
	UpdateQuery(context.Context, *SavedQueryReq) (*SavedQueryRecord, error)
	GetQueries(context.Context, *Empty) (*SavedQueryList, error)
	DeleteQuery(context.Context, *IdRequest) (*IdRequest, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) QueryPublishedWorkflows(context.Context, *QueryReq) (*QueryRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPublishedWorkflows not implemented")
}
func (UnimplementedQueryServer) CreateQuery(context.Context, *SavedQueryReq) (*SavedQueryRecord, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQuery not implemented")
}
func (UnimplementedQueryServer) UpdateQuery(context.Context, *SavedQueryReq) (*SavedQueryRecord, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQuery not implemented")
}
func (UnimplementedQueryServer) GetQueries(context.Context, *Empty) (*SavedQueryList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQueries not implemented")
}
func (UnimplementedQueryServer) DeleteQuery(context.Context, *IdRequest) (*IdRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQuery not implemented")
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

func _Query_QueryPublishedWorkflows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryPublishedWorkflows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitoring.Query/QueryPublishedWorkflows",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryPublishedWorkflows(ctx, req.(*QueryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CreateQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SavedQueryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CreateQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitoring.Query/CreateQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CreateQuery(ctx, req.(*SavedQueryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_UpdateQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SavedQueryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).UpdateQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitoring.Query/UpdateQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).UpdateQuery(ctx, req.(*SavedQueryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetQueries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetQueries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitoring.Query/GetQueries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetQueries(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DeleteQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DeleteQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitoring.Query/DeleteQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DeleteQuery(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "monitoring.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryPublishedWorkflows",
			Handler:    _Query_QueryPublishedWorkflows_Handler,
		},
		{
			MethodName: "CreateQuery",
			Handler:    _Query_CreateQuery_Handler,
		},
		{
			MethodName: "UpdateQuery",
			Handler:    _Query_UpdateQuery_Handler,
		},
		{
			MethodName: "GetQueries",
			Handler:    _Query_GetQueries_Handler,
		},
		{
			MethodName: "DeleteQuery",
			Handler:    _Query_DeleteQuery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "monitoring/monitoring.proto",
}
