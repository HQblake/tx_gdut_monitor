// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: judgment2store.proto

package judgpb

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

// MetricServiceClient is the client API for MetricService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetricServiceClient interface {
	// 获取聚合数据
	GetAggregatedData(ctx context.Context, in *AggregatedRequest, opts ...grpc.CallOption) (*AggregatedResponse, error)
	// 插入告警信息
	InsertAlertInfo(ctx context.Context, in *HistoryInfoRequest, opts ...grpc.CallOption) (*HistoryInfoResponse, error)
}

type metricServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricServiceClient(cc grpc.ClientConnInterface) MetricServiceClient {
	return &metricServiceClient{cc}
}

func (c *metricServiceClient) GetAggregatedData(ctx context.Context, in *AggregatedRequest, opts ...grpc.CallOption) (*AggregatedResponse, error) {
	out := new(AggregatedResponse)
	err := c.cc.Invoke(ctx, "/proto.MetricService/GetAggregatedData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricServiceClient) InsertAlertInfo(ctx context.Context, in *HistoryInfoRequest, opts ...grpc.CallOption) (*HistoryInfoResponse, error) {
	out := new(HistoryInfoResponse)
	err := c.cc.Invoke(ctx, "/proto.MetricService/InsertAlertInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricServiceServer is the server API for MetricService service.
// All implementations must embed UnimplementedMetricServiceServer
// for forward compatibility
type MetricServiceServer interface {
	// 获取聚合数据
	GetAggregatedData(context.Context, *AggregatedRequest) (*AggregatedResponse, error)
	// 插入告警信息
	InsertAlertInfo(context.Context, *HistoryInfoRequest) (*HistoryInfoResponse, error)
	mustEmbedUnimplementedMetricServiceServer()
}

// UnimplementedMetricServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMetricServiceServer struct {
}

func (UnimplementedMetricServiceServer) GetAggregatedData(context.Context, *AggregatedRequest) (*AggregatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAggregatedData not implemented")
}
func (UnimplementedMetricServiceServer) InsertAlertInfo(context.Context, *HistoryInfoRequest) (*HistoryInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertAlertInfo not implemented")
}
func (UnimplementedMetricServiceServer) mustEmbedUnimplementedMetricServiceServer() {}

// UnsafeMetricServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetricServiceServer will
// result in compilation errors.
type UnsafeMetricServiceServer interface {
	mustEmbedUnimplementedMetricServiceServer()
}

func RegisterMetricServiceServer(s grpc.ServiceRegistrar, srv MetricServiceServer) {
	s.RegisterService(&MetricService_ServiceDesc, srv)
}

func _MetricService_GetAggregatedData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AggregatedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricServiceServer).GetAggregatedData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MetricService/GetAggregatedData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricServiceServer).GetAggregatedData(ctx, req.(*AggregatedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricService_InsertAlertInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HistoryInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricServiceServer).InsertAlertInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MetricService/InsertAlertInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricServiceServer).InsertAlertInfo(ctx, req.(*HistoryInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MetricService_ServiceDesc is the grpc.ServiceDesc for MetricService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetricService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MetricService",
	HandlerType: (*MetricServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAggregatedData",
			Handler:    _MetricService_GetAggregatedData_Handler,
		},
		{
			MethodName: "InsertAlertInfo",
			Handler:    _MetricService_InsertAlertInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "judgment2store.proto",
}
