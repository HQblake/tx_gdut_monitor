// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// ManageConfigClient is the client API for ManageConfig service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManageConfigClient interface {
	//
	// 1. 查询所有存活的Agent信息，返回结果：ip、port、local、isLive
	// 2. 根据AgentID（或IP和Local）查询Agent的所有metric
	GetAllAgentInfo(ctx context.Context, in *AgentRequest, opts ...grpc.CallOption) (ManageConfig_GetAllAgentInfoClient, error)
	GetMetricsByAgentID(ctx context.Context, in *AgentRequest, opts ...grpc.CallOption) (ManageConfig_GetMetricsByAgentIDClient, error)
	//
	// 1. 根据AgentID、metric、开始时间戳、结束时间戳、聚合时间获取上报的指标数据列表（流式响应）
	GetMetricData(ctx context.Context, in *MetricRequest, opts ...grpc.CallOption) (ManageConfig_GetMetricDataClient, error)
	//
	// 1. 查询所有告警信息，返回一个列表（流式响应）
	// 2. 根据id删除告警信息
	// 3. 根据id获取告警信息
	// 4. 根据agentID、等级、日期、指标等查看告警信息，返回一个列表（流式响应）
	GetAllAlertInfo(ctx context.Context, in *AlertRequest, opts ...grpc.CallOption) (ManageConfig_GetAllAlertInfoClient, error)
	GetAlertInfo(ctx context.Context, in *AlertRequest, opts ...grpc.CallOption) (ManageConfig_GetAlertInfoClient, error)
	DelAlterInfo(ctx context.Context, in *AlertRequest, opts ...grpc.CallOption) (*AlertResponse, error)
	//
	// 1. 新增配置
	// 2. 根据id更新配置
	// 3. 根据id删除配置
	// 4. 根据id获取配置详情
	// 5. 获取所有的配置列表，返回一个列表（流式响应）
	// 6. 根据agentID（或ip和local）查看配置，返回一个列表（流式响应）
	AddConfig(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error)
	UpdateConfigById(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error)
	DeleteConfigById(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error)
	GetConfig(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (ManageConfig_GetConfigClient, error)
	GetAllConfigs(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (ManageConfig_GetAllConfigsClient, error)
}

type manageConfigClient struct {
	cc grpc.ClientConnInterface
}

func NewManageConfigClient(cc grpc.ClientConnInterface) ManageConfigClient {
	return &manageConfigClient{cc}
}

func (c *manageConfigClient) GetAllAgentInfo(ctx context.Context, in *AgentRequest, opts ...grpc.CallOption) (ManageConfig_GetAllAgentInfoClient, error) {
	stream, err := c.cc.NewStream(ctx, &ManageConfig_ServiceDesc.Streams[0], "/proto.ManageConfig/GetAllAgentInfo", opts...)
	if err != nil {
		return nil, err
	}
	x := &manageConfigGetAllAgentInfoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ManageConfig_GetAllAgentInfoClient interface {
	Recv() (*AgentResponse, error)
	grpc.ClientStream
}

type manageConfigGetAllAgentInfoClient struct {
	grpc.ClientStream
}

func (x *manageConfigGetAllAgentInfoClient) Recv() (*AgentResponse, error) {
	m := new(AgentResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *manageConfigClient) GetMetricsByAgentID(ctx context.Context, in *AgentRequest, opts ...grpc.CallOption) (ManageConfig_GetMetricsByAgentIDClient, error) {
	stream, err := c.cc.NewStream(ctx, &ManageConfig_ServiceDesc.Streams[1], "/proto.ManageConfig/GetMetricsByAgentID", opts...)
	if err != nil {
		return nil, err
	}
	x := &manageConfigGetMetricsByAgentIDClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ManageConfig_GetMetricsByAgentIDClient interface {
	Recv() (*AgentResponse, error)
	grpc.ClientStream
}

type manageConfigGetMetricsByAgentIDClient struct {
	grpc.ClientStream
}

func (x *manageConfigGetMetricsByAgentIDClient) Recv() (*AgentResponse, error) {
	m := new(AgentResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *manageConfigClient) GetMetricData(ctx context.Context, in *MetricRequest, opts ...grpc.CallOption) (ManageConfig_GetMetricDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &ManageConfig_ServiceDesc.Streams[2], "/proto.ManageConfig/GetMetricData", opts...)
	if err != nil {
		return nil, err
	}
	x := &manageConfigGetMetricDataClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ManageConfig_GetMetricDataClient interface {
	Recv() (*MetricResponse, error)
	grpc.ClientStream
}

type manageConfigGetMetricDataClient struct {
	grpc.ClientStream
}

func (x *manageConfigGetMetricDataClient) Recv() (*MetricResponse, error) {
	m := new(MetricResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *manageConfigClient) GetAllAlertInfo(ctx context.Context, in *AlertRequest, opts ...grpc.CallOption) (ManageConfig_GetAllAlertInfoClient, error) {
	stream, err := c.cc.NewStream(ctx, &ManageConfig_ServiceDesc.Streams[3], "/proto.ManageConfig/GetAllAlertInfo", opts...)
	if err != nil {
		return nil, err
	}
	x := &manageConfigGetAllAlertInfoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ManageConfig_GetAllAlertInfoClient interface {
	Recv() (*AlertResponse, error)
	grpc.ClientStream
}

type manageConfigGetAllAlertInfoClient struct {
	grpc.ClientStream
}

func (x *manageConfigGetAllAlertInfoClient) Recv() (*AlertResponse, error) {
	m := new(AlertResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *manageConfigClient) GetAlertInfo(ctx context.Context, in *AlertRequest, opts ...grpc.CallOption) (ManageConfig_GetAlertInfoClient, error) {
	stream, err := c.cc.NewStream(ctx, &ManageConfig_ServiceDesc.Streams[4], "/proto.ManageConfig/GetAlertInfo", opts...)
	if err != nil {
		return nil, err
	}
	x := &manageConfigGetAlertInfoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ManageConfig_GetAlertInfoClient interface {
	Recv() (*AlertResponse, error)
	grpc.ClientStream
}

type manageConfigGetAlertInfoClient struct {
	grpc.ClientStream
}

func (x *manageConfigGetAlertInfoClient) Recv() (*AlertResponse, error) {
	m := new(AlertResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *manageConfigClient) DelAlterInfo(ctx context.Context, in *AlertRequest, opts ...grpc.CallOption) (*AlertResponse, error) {
	out := new(AlertResponse)
	err := c.cc.Invoke(ctx, "/proto.ManageConfig/DelAlterInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageConfigClient) AddConfig(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error) {
	out := new(ConfigResponse)
	err := c.cc.Invoke(ctx, "/proto.ManageConfig/AddConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageConfigClient) UpdateConfigById(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error) {
	out := new(ConfigResponse)
	err := c.cc.Invoke(ctx, "/proto.ManageConfig/UpdateConfigById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageConfigClient) DeleteConfigById(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error) {
	out := new(ConfigResponse)
	err := c.cc.Invoke(ctx, "/proto.ManageConfig/DeleteConfigById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageConfigClient) GetConfig(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (ManageConfig_GetConfigClient, error) {
	stream, err := c.cc.NewStream(ctx, &ManageConfig_ServiceDesc.Streams[5], "/proto.ManageConfig/GetConfig", opts...)
	if err != nil {
		return nil, err
	}
	x := &manageConfigGetConfigClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ManageConfig_GetConfigClient interface {
	Recv() (*ConfigResponse, error)
	grpc.ClientStream
}

type manageConfigGetConfigClient struct {
	grpc.ClientStream
}

func (x *manageConfigGetConfigClient) Recv() (*ConfigResponse, error) {
	m := new(ConfigResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *manageConfigClient) GetAllConfigs(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (ManageConfig_GetAllConfigsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ManageConfig_ServiceDesc.Streams[6], "/proto.ManageConfig/GetAllConfigs", opts...)
	if err != nil {
		return nil, err
	}
	x := &manageConfigGetAllConfigsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ManageConfig_GetAllConfigsClient interface {
	Recv() (*ConfigResponse, error)
	grpc.ClientStream
}

type manageConfigGetAllConfigsClient struct {
	grpc.ClientStream
}

func (x *manageConfigGetAllConfigsClient) Recv() (*ConfigResponse, error) {
	m := new(ConfigResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ManageConfigServer is the server API for ManageConfig service.
// All implementations must embed UnimplementedManageConfigServer
// for forward compatibility
type ManageConfigServer interface {
	//
	// 1. 查询所有存活的Agent信息，返回结果：ip、port、local、isLive
	// 2. 根据AgentID（或IP和Local）查询Agent的所有metric
	GetAllAgentInfo(*AgentRequest, ManageConfig_GetAllAgentInfoServer) error
	GetMetricsByAgentID(*AgentRequest, ManageConfig_GetMetricsByAgentIDServer) error
	//
	// 1. 根据AgentID、metric、开始时间戳、结束时间戳、聚合时间获取上报的指标数据列表（流式响应）
	GetMetricData(*MetricRequest, ManageConfig_GetMetricDataServer) error
	//
	// 1. 查询所有告警信息，返回一个列表（流式响应）
	// 2. 根据id删除告警信息
	// 3. 根据id获取告警信息
	// 4. 根据agentID、等级、日期、指标等查看告警信息，返回一个列表（流式响应）
	GetAllAlertInfo(*AlertRequest, ManageConfig_GetAllAlertInfoServer) error
	GetAlertInfo(*AlertRequest, ManageConfig_GetAlertInfoServer) error
	DelAlterInfo(context.Context, *AlertRequest) (*AlertResponse, error)
	//
	// 1. 新增配置
	// 2. 根据id更新配置
	// 3. 根据id删除配置
	// 4. 根据id获取配置详情
	// 5. 获取所有的配置列表，返回一个列表（流式响应）
	// 6. 根据agentID（或ip和local）查看配置，返回一个列表（流式响应）
	AddConfig(context.Context, *ConfigRequest) (*ConfigResponse, error)
	UpdateConfigById(context.Context, *ConfigRequest) (*ConfigResponse, error)
	DeleteConfigById(context.Context, *ConfigRequest) (*ConfigResponse, error)
	GetConfig(*ConfigRequest, ManageConfig_GetConfigServer) error
	GetAllConfigs(*ConfigRequest, ManageConfig_GetAllConfigsServer) error
	mustEmbedUnimplementedManageConfigServer()
}

// UnimplementedManageConfigServer must be embedded to have forward compatible implementations.
type UnimplementedManageConfigServer struct {
}

func (UnimplementedManageConfigServer) GetAllAgentInfo(*AgentRequest, ManageConfig_GetAllAgentInfoServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllAgentInfo not implemented")
}
func (UnimplementedManageConfigServer) GetMetricsByAgentID(*AgentRequest, ManageConfig_GetMetricsByAgentIDServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMetricsByAgentID not implemented")
}
func (UnimplementedManageConfigServer) GetMetricData(*MetricRequest, ManageConfig_GetMetricDataServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMetricData not implemented")
}
func (UnimplementedManageConfigServer) GetAllAlertInfo(*AlertRequest, ManageConfig_GetAllAlertInfoServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllAlertInfo not implemented")
}
func (UnimplementedManageConfigServer) GetAlertInfo(*AlertRequest, ManageConfig_GetAlertInfoServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAlertInfo not implemented")
}
func (UnimplementedManageConfigServer) DelAlterInfo(context.Context, *AlertRequest) (*AlertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelAlterInfo not implemented")
}
func (UnimplementedManageConfigServer) AddConfig(context.Context, *ConfigRequest) (*ConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddConfig not implemented")
}
func (UnimplementedManageConfigServer) UpdateConfigById(context.Context, *ConfigRequest) (*ConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConfigById not implemented")
}
func (UnimplementedManageConfigServer) DeleteConfigById(context.Context, *ConfigRequest) (*ConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConfigById not implemented")
}
func (UnimplementedManageConfigServer) GetConfig(*ConfigRequest, ManageConfig_GetConfigServer) error {
	return status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (UnimplementedManageConfigServer) GetAllConfigs(*ConfigRequest, ManageConfig_GetAllConfigsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllConfigs not implemented")
}
func (UnimplementedManageConfigServer) mustEmbedUnimplementedManageConfigServer() {}

// UnsafeManageConfigServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManageConfigServer will
// result in compilation errors.
type UnsafeManageConfigServer interface {
	mustEmbedUnimplementedManageConfigServer()
}

func RegisterManageConfigServer(s grpc.ServiceRegistrar, srv ManageConfigServer) {
	s.RegisterService(&ManageConfig_ServiceDesc, srv)
}

func _ManageConfig_GetAllAgentInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AgentRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ManageConfigServer).GetAllAgentInfo(m, &manageConfigGetAllAgentInfoServer{stream})
}

type ManageConfig_GetAllAgentInfoServer interface {
	Send(*AgentResponse) error
	grpc.ServerStream
}

type manageConfigGetAllAgentInfoServer struct {
	grpc.ServerStream
}

func (x *manageConfigGetAllAgentInfoServer) Send(m *AgentResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ManageConfig_GetMetricsByAgentID_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AgentRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ManageConfigServer).GetMetricsByAgentID(m, &manageConfigGetMetricsByAgentIDServer{stream})
}

type ManageConfig_GetMetricsByAgentIDServer interface {
	Send(*AgentResponse) error
	grpc.ServerStream
}

type manageConfigGetMetricsByAgentIDServer struct {
	grpc.ServerStream
}

func (x *manageConfigGetMetricsByAgentIDServer) Send(m *AgentResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ManageConfig_GetMetricData_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MetricRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ManageConfigServer).GetMetricData(m, &manageConfigGetMetricDataServer{stream})
}

type ManageConfig_GetMetricDataServer interface {
	Send(*MetricResponse) error
	grpc.ServerStream
}

type manageConfigGetMetricDataServer struct {
	grpc.ServerStream
}

func (x *manageConfigGetMetricDataServer) Send(m *MetricResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ManageConfig_GetAllAlertInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AlertRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ManageConfigServer).GetAllAlertInfo(m, &manageConfigGetAllAlertInfoServer{stream})
}

type ManageConfig_GetAllAlertInfoServer interface {
	Send(*AlertResponse) error
	grpc.ServerStream
}

type manageConfigGetAllAlertInfoServer struct {
	grpc.ServerStream
}

func (x *manageConfigGetAllAlertInfoServer) Send(m *AlertResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ManageConfig_GetAlertInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AlertRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ManageConfigServer).GetAlertInfo(m, &manageConfigGetAlertInfoServer{stream})
}

type ManageConfig_GetAlertInfoServer interface {
	Send(*AlertResponse) error
	grpc.ServerStream
}

type manageConfigGetAlertInfoServer struct {
	grpc.ServerStream
}

func (x *manageConfigGetAlertInfoServer) Send(m *AlertResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ManageConfig_DelAlterInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageConfigServer).DelAlterInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ManageConfig/DelAlterInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageConfigServer).DelAlterInfo(ctx, req.(*AlertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManageConfig_AddConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageConfigServer).AddConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ManageConfig/AddConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageConfigServer).AddConfig(ctx, req.(*ConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManageConfig_UpdateConfigById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageConfigServer).UpdateConfigById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ManageConfig/UpdateConfigById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageConfigServer).UpdateConfigById(ctx, req.(*ConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManageConfig_DeleteConfigById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageConfigServer).DeleteConfigById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ManageConfig/DeleteConfigById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageConfigServer).DeleteConfigById(ctx, req.(*ConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManageConfig_GetConfig_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConfigRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ManageConfigServer).GetConfig(m, &manageConfigGetConfigServer{stream})
}

type ManageConfig_GetConfigServer interface {
	Send(*ConfigResponse) error
	grpc.ServerStream
}

type manageConfigGetConfigServer struct {
	grpc.ServerStream
}

func (x *manageConfigGetConfigServer) Send(m *ConfigResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ManageConfig_GetAllConfigs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConfigRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ManageConfigServer).GetAllConfigs(m, &manageConfigGetAllConfigsServer{stream})
}

type ManageConfig_GetAllConfigsServer interface {
	Send(*ConfigResponse) error
	grpc.ServerStream
}

type manageConfigGetAllConfigsServer struct {
	grpc.ServerStream
}

func (x *manageConfigGetAllConfigsServer) Send(m *ConfigResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ManageConfig_ServiceDesc is the grpc.ServiceDesc for ManageConfig service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ManageConfig_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ManageConfig",
	HandlerType: (*ManageConfigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DelAlterInfo",
			Handler:    _ManageConfig_DelAlterInfo_Handler,
		},
		{
			MethodName: "AddConfig",
			Handler:    _ManageConfig_AddConfig_Handler,
		},
		{
			MethodName: "UpdateConfigById",
			Handler:    _ManageConfig_UpdateConfigById_Handler,
		},
		{
			MethodName: "DeleteConfigById",
			Handler:    _ManageConfig_DeleteConfigById_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllAgentInfo",
			Handler:       _ManageConfig_GetAllAgentInfo_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetMetricsByAgentID",
			Handler:       _ManageConfig_GetMetricsByAgentID_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetMetricData",
			Handler:       _ManageConfig_GetMetricData_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetAllAlertInfo",
			Handler:       _ManageConfig_GetAllAlertInfo_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetAlertInfo",
			Handler:       _ManageConfig_GetAlertInfo_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetConfig",
			Handler:       _ManageConfig_GetConfig_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetAllConfigs",
			Handler:       _ManageConfig_GetAllConfigs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "manage2store.proto",
}
