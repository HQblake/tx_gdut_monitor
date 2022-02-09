// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package managepb

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

// HistoryServiceClient is the client API for HistoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HistoryServiceClient interface {
	//
	// 1. 查询所有告警信息，返回一个列表（流式响应）
	// 2. 根据id删除告警信息
	// 3. 根据id获取告警信息
	// 4. 根据agentID、等级、日期、指标等查看告警信息，返回一个列表（流式响应）
	GetAllAlertInfo(ctx context.Context, in *BaseRequest, opts ...grpc.CallOption) (HistoryService_GetAllAlertInfoClient, error)
	// 既可以根据单个agentID或单个ID获取信息，也可以通过其他额外的补充信息获取某个区间的告警信息
	GetAlertInfo(ctx context.Context, in *AlertRequest, opts ...grpc.CallOption) (HistoryService_GetAlertInfoClient, error)
	DelAlterInfo(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*BaseResponse, error)
}

type historyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHistoryServiceClient(cc grpc.ClientConnInterface) HistoryServiceClient {
	return &historyServiceClient{cc}
}

func (c *historyServiceClient) GetAllAlertInfo(ctx context.Context, in *BaseRequest, opts ...grpc.CallOption) (HistoryService_GetAllAlertInfoClient, error) {
	stream, err := c.cc.NewStream(ctx, &HistoryService_ServiceDesc.Streams[0], "/judgment2store.HistoryService/GetAllAlertInfo", opts...)
	if err != nil {
		return nil, err
	}
	x := &historyServiceGetAllAlertInfoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HistoryService_GetAllAlertInfoClient interface {
	Recv() (*AlertResponse, error)
	grpc.ClientStream
}

type historyServiceGetAllAlertInfoClient struct {
	grpc.ClientStream
}

func (x *historyServiceGetAllAlertInfoClient) Recv() (*AlertResponse, error) {
	m := new(AlertResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *historyServiceClient) GetAlertInfo(ctx context.Context, in *AlertRequest, opts ...grpc.CallOption) (HistoryService_GetAlertInfoClient, error) {
	stream, err := c.cc.NewStream(ctx, &HistoryService_ServiceDesc.Streams[1], "/judgment2store.HistoryService/GetAlertInfo", opts...)
	if err != nil {
		return nil, err
	}
	x := &historyServiceGetAlertInfoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HistoryService_GetAlertInfoClient interface {
	Recv() (*AlertResponse, error)
	grpc.ClientStream
}

type historyServiceGetAlertInfoClient struct {
	grpc.ClientStream
}

func (x *historyServiceGetAlertInfoClient) Recv() (*AlertResponse, error) {
	m := new(AlertResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *historyServiceClient) DelAlterInfo(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/judgment2store.HistoryService/DelAlterInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HistoryServiceServer is the server API for HistoryService service.
// All implementations must embed UnimplementedHistoryServiceServer
// for forward compatibility
type HistoryServiceServer interface {
	//
	// 1. 查询所有告警信息，返回一个列表（流式响应）
	// 2. 根据id删除告警信息
	// 3. 根据id获取告警信息
	// 4. 根据agentID、等级、日期、指标等查看告警信息，返回一个列表（流式响应）
	GetAllAlertInfo(*BaseRequest, HistoryService_GetAllAlertInfoServer) error
	// 既可以根据单个agentID或单个ID获取信息，也可以通过其他额外的补充信息获取某个区间的告警信息
	GetAlertInfo(*AlertRequest, HistoryService_GetAlertInfoServer) error
	DelAlterInfo(context.Context, *IDRequest) (*BaseResponse, error)
	mustEmbedUnimplementedHistoryServiceServer()
}

// UnimplementedHistoryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHistoryServiceServer struct {
}

func (UnimplementedHistoryServiceServer) GetAllAlertInfo(*BaseRequest, HistoryService_GetAllAlertInfoServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllAlertInfo not implemented")
}
func (UnimplementedHistoryServiceServer) GetAlertInfo(*AlertRequest, HistoryService_GetAlertInfoServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAlertInfo not implemented")
}
func (UnimplementedHistoryServiceServer) DelAlterInfo(context.Context, *IDRequest) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelAlterInfo not implemented")
}
func (UnimplementedHistoryServiceServer) mustEmbedUnimplementedHistoryServiceServer() {}

// UnsafeHistoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HistoryServiceServer will
// result in compilation errors.
type UnsafeHistoryServiceServer interface {
	mustEmbedUnimplementedHistoryServiceServer()
}

func RegisterHistoryServiceServer(s grpc.ServiceRegistrar, srv HistoryServiceServer) {
	s.RegisterService(&HistoryService_ServiceDesc, srv)
}

func _HistoryService_GetAllAlertInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BaseRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HistoryServiceServer).GetAllAlertInfo(m, &historyServiceGetAllAlertInfoServer{stream})
}

type HistoryService_GetAllAlertInfoServer interface {
	Send(*AlertResponse) error
	grpc.ServerStream
}

type historyServiceGetAllAlertInfoServer struct {
	grpc.ServerStream
}

func (x *historyServiceGetAllAlertInfoServer) Send(m *AlertResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _HistoryService_GetAlertInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AlertRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HistoryServiceServer).GetAlertInfo(m, &historyServiceGetAlertInfoServer{stream})
}

type HistoryService_GetAlertInfoServer interface {
	Send(*AlertResponse) error
	grpc.ServerStream
}

type historyServiceGetAlertInfoServer struct {
	grpc.ServerStream
}

func (x *historyServiceGetAlertInfoServer) Send(m *AlertResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _HistoryService_DelAlterInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HistoryServiceServer).DelAlterInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/judgment2store.HistoryService/DelAlterInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HistoryServiceServer).DelAlterInfo(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HistoryService_ServiceDesc is the grpc.ServiceDesc for HistoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HistoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "judgment2store.HistoryService",
	HandlerType: (*HistoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DelAlterInfo",
			Handler:    _HistoryService_DelAlterInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllAlertInfo",
			Handler:       _HistoryService_GetAllAlertInfo_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetAlertInfo",
			Handler:       _HistoryService_GetAlertInfo_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "manage2store.proto",
}

// MetricServiceClient is the client API for MetricService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetricServiceClient interface {
	//
	// 1. 根据AgentID、metric、开始时间戳、结束时间戳、聚合时间获取上报的指标数据列表（流式响应）
	GetMetricData(ctx context.Context, in *MetricRequest, opts ...grpc.CallOption) (MetricService_GetMetricDataClient, error)
}

type metricServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricServiceClient(cc grpc.ClientConnInterface) MetricServiceClient {
	return &metricServiceClient{cc}
}

func (c *metricServiceClient) GetMetricData(ctx context.Context, in *MetricRequest, opts ...grpc.CallOption) (MetricService_GetMetricDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &MetricService_ServiceDesc.Streams[0], "/judgment2store.MetricService/GetMetricData", opts...)
	if err != nil {
		return nil, err
	}
	x := &metricServiceGetMetricDataClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MetricService_GetMetricDataClient interface {
	Recv() (*MetricResponse, error)
	grpc.ClientStream
}

type metricServiceGetMetricDataClient struct {
	grpc.ClientStream
}

func (x *metricServiceGetMetricDataClient) Recv() (*MetricResponse, error) {
	m := new(MetricResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MetricServiceServer is the server API for MetricService service.
// All implementations must embed UnimplementedMetricServiceServer
// for forward compatibility
type MetricServiceServer interface {
	//
	// 1. 根据AgentID、metric、开始时间戳、结束时间戳、聚合时间获取上报的指标数据列表（流式响应）
	GetMetricData(*MetricRequest, MetricService_GetMetricDataServer) error
	mustEmbedUnimplementedMetricServiceServer()
}

// UnimplementedMetricServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMetricServiceServer struct {
}

func (UnimplementedMetricServiceServer) GetMetricData(*MetricRequest, MetricService_GetMetricDataServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMetricData not implemented")
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

func _MetricService_GetMetricData_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MetricRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MetricServiceServer).GetMetricData(m, &metricServiceGetMetricDataServer{stream})
}

type MetricService_GetMetricDataServer interface {
	Send(*MetricResponse) error
	grpc.ServerStream
}

type metricServiceGetMetricDataServer struct {
	grpc.ServerStream
}

func (x *metricServiceGetMetricDataServer) Send(m *MetricResponse) error {
	return x.ServerStream.SendMsg(m)
}

// MetricService_ServiceDesc is the grpc.ServiceDesc for MetricService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetricService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "judgment2store.MetricService",
	HandlerType: (*MetricServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetMetricData",
			Handler:       _MetricService_GetMetricData_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "manage2store.proto",
}

// AgentServiceClient is the client API for AgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentServiceClient interface {
	// 查询Agent信息
	// 查询所有存活的Agent信息，返回结果：ip、port、local、isLive、metrics列表，流形式返回
	GetAllAgentInfo(ctx context.Context, in *BaseRequest, opts ...grpc.CallOption) (AgentService_GetAllAgentInfoClient, error)
	// 查询指定Agent(ip和local)的信息，返回结果：ip、port、local、isLive、metrics列表
	GetAgentInfoByAgentID(ctx context.Context, in *AgentRequest, opts ...grpc.CallOption) (*AgentResponse, error)
}

type agentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentServiceClient(cc grpc.ClientConnInterface) AgentServiceClient {
	return &agentServiceClient{cc}
}

func (c *agentServiceClient) GetAllAgentInfo(ctx context.Context, in *BaseRequest, opts ...grpc.CallOption) (AgentService_GetAllAgentInfoClient, error) {
	stream, err := c.cc.NewStream(ctx, &AgentService_ServiceDesc.Streams[0], "/judgment2store.AgentService/GetAllAgentInfo", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentServiceGetAllAgentInfoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AgentService_GetAllAgentInfoClient interface {
	Recv() (*AgentResponse, error)
	grpc.ClientStream
}

type agentServiceGetAllAgentInfoClient struct {
	grpc.ClientStream
}

func (x *agentServiceGetAllAgentInfoClient) Recv() (*AgentResponse, error) {
	m := new(AgentResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *agentServiceClient) GetAgentInfoByAgentID(ctx context.Context, in *AgentRequest, opts ...grpc.CallOption) (*AgentResponse, error) {
	out := new(AgentResponse)
	err := c.cc.Invoke(ctx, "/judgment2store.AgentService/GetAgentInfoByAgentID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServiceServer is the server API for AgentService service.
// All implementations must embed UnimplementedAgentServiceServer
// for forward compatibility
type AgentServiceServer interface {
	// 查询Agent信息
	// 查询所有存活的Agent信息，返回结果：ip、port、local、isLive、metrics列表，流形式返回
	GetAllAgentInfo(*BaseRequest, AgentService_GetAllAgentInfoServer) error
	// 查询指定Agent(ip和local)的信息，返回结果：ip、port、local、isLive、metrics列表
	GetAgentInfoByAgentID(context.Context, *AgentRequest) (*AgentResponse, error)
	mustEmbedUnimplementedAgentServiceServer()
}

// UnimplementedAgentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAgentServiceServer struct {
}

func (UnimplementedAgentServiceServer) GetAllAgentInfo(*BaseRequest, AgentService_GetAllAgentInfoServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllAgentInfo not implemented")
}
func (UnimplementedAgentServiceServer) GetAgentInfoByAgentID(context.Context, *AgentRequest) (*AgentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAgentInfoByAgentID not implemented")
}
func (UnimplementedAgentServiceServer) mustEmbedUnimplementedAgentServiceServer() {}

// UnsafeAgentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentServiceServer will
// result in compilation errors.
type UnsafeAgentServiceServer interface {
	mustEmbedUnimplementedAgentServiceServer()
}

func RegisterAgentServiceServer(s grpc.ServiceRegistrar, srv AgentServiceServer) {
	s.RegisterService(&AgentService_ServiceDesc, srv)
}

func _AgentService_GetAllAgentInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BaseRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AgentServiceServer).GetAllAgentInfo(m, &agentServiceGetAllAgentInfoServer{stream})
}

type AgentService_GetAllAgentInfoServer interface {
	Send(*AgentResponse) error
	grpc.ServerStream
}

type agentServiceGetAllAgentInfoServer struct {
	grpc.ServerStream
}

func (x *agentServiceGetAllAgentInfoServer) Send(m *AgentResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _AgentService_GetAgentInfoByAgentID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).GetAgentInfoByAgentID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/judgment2store.AgentService/GetAgentInfoByAgentID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).GetAgentInfoByAgentID(ctx, req.(*AgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AgentService_ServiceDesc is the grpc.ServiceDesc for AgentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "judgment2store.AgentService",
	HandlerType: (*AgentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAgentInfoByAgentID",
			Handler:    _AgentService_GetAgentInfoByAgentID_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllAgentInfo",
			Handler:       _AgentService_GetAllAgentInfo_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "manage2store.proto",
}

// JudgmentServiceClient is the client API for JudgmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JudgmentServiceClient interface {
	//
	// 1. 根据agentID（ip和local）查看配置，返回一个列表（流式响应）
	// 2. 根据id更新配置
	// 3. 根据id删除配置
	GetConfigsByAgent(ctx context.Context, in *AgentRequest, opts ...grpc.CallOption) (JudgmentService_GetConfigsByAgentClient, error)
	UpdateConfig(ctx context.Context, in *JudgmentEntry, opts ...grpc.CallOption) (*BaseResponse, error)
	DeleteConfig(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*BaseResponse, error)
}

type judgmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJudgmentServiceClient(cc grpc.ClientConnInterface) JudgmentServiceClient {
	return &judgmentServiceClient{cc}
}

func (c *judgmentServiceClient) GetConfigsByAgent(ctx context.Context, in *AgentRequest, opts ...grpc.CallOption) (JudgmentService_GetConfigsByAgentClient, error) {
	stream, err := c.cc.NewStream(ctx, &JudgmentService_ServiceDesc.Streams[0], "/judgment2store.JudgmentService/GetConfigsByAgent", opts...)
	if err != nil {
		return nil, err
	}
	x := &judgmentServiceGetConfigsByAgentClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type JudgmentService_GetConfigsByAgentClient interface {
	Recv() (*JudgmentConfigResponse, error)
	grpc.ClientStream
}

type judgmentServiceGetConfigsByAgentClient struct {
	grpc.ClientStream
}

func (x *judgmentServiceGetConfigsByAgentClient) Recv() (*JudgmentConfigResponse, error) {
	m := new(JudgmentConfigResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *judgmentServiceClient) UpdateConfig(ctx context.Context, in *JudgmentEntry, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/judgment2store.JudgmentService/UpdateConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *judgmentServiceClient) DeleteConfig(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/judgment2store.JudgmentService/DeleteConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JudgmentServiceServer is the server API for JudgmentService service.
// All implementations must embed UnimplementedJudgmentServiceServer
// for forward compatibility
type JudgmentServiceServer interface {
	//
	// 1. 根据agentID（ip和local）查看配置，返回一个列表（流式响应）
	// 2. 根据id更新配置
	// 3. 根据id删除配置
	GetConfigsByAgent(*AgentRequest, JudgmentService_GetConfigsByAgentServer) error
	UpdateConfig(context.Context, *JudgmentEntry) (*BaseResponse, error)
	DeleteConfig(context.Context, *IDRequest) (*BaseResponse, error)
	mustEmbedUnimplementedJudgmentServiceServer()
}

// UnimplementedJudgmentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedJudgmentServiceServer struct {
}

func (UnimplementedJudgmentServiceServer) GetConfigsByAgent(*AgentRequest, JudgmentService_GetConfigsByAgentServer) error {
	return status.Errorf(codes.Unimplemented, "method GetConfigsByAgent not implemented")
}
func (UnimplementedJudgmentServiceServer) UpdateConfig(context.Context, *JudgmentEntry) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConfig not implemented")
}
func (UnimplementedJudgmentServiceServer) DeleteConfig(context.Context, *IDRequest) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConfig not implemented")
}
func (UnimplementedJudgmentServiceServer) mustEmbedUnimplementedJudgmentServiceServer() {}

// UnsafeJudgmentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JudgmentServiceServer will
// result in compilation errors.
type UnsafeJudgmentServiceServer interface {
	mustEmbedUnimplementedJudgmentServiceServer()
}

func RegisterJudgmentServiceServer(s grpc.ServiceRegistrar, srv JudgmentServiceServer) {
	s.RegisterService(&JudgmentService_ServiceDesc, srv)
}

func _JudgmentService_GetConfigsByAgent_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AgentRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(JudgmentServiceServer).GetConfigsByAgent(m, &judgmentServiceGetConfigsByAgentServer{stream})
}

type JudgmentService_GetConfigsByAgentServer interface {
	Send(*JudgmentConfigResponse) error
	grpc.ServerStream
}

type judgmentServiceGetConfigsByAgentServer struct {
	grpc.ServerStream
}

func (x *judgmentServiceGetConfigsByAgentServer) Send(m *JudgmentConfigResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _JudgmentService_UpdateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JudgmentEntry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudgmentServiceServer).UpdateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/judgment2store.JudgmentService/UpdateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudgmentServiceServer).UpdateConfig(ctx, req.(*JudgmentEntry))
	}
	return interceptor(ctx, in, info, handler)
}

func _JudgmentService_DeleteConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudgmentServiceServer).DeleteConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/judgment2store.JudgmentService/DeleteConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudgmentServiceServer).DeleteConfig(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// JudgmentService_ServiceDesc is the grpc.ServiceDesc for JudgmentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JudgmentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "judgment2store.JudgmentService",
	HandlerType: (*JudgmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateConfig",
			Handler:    _JudgmentService_UpdateConfig_Handler,
		},
		{
			MethodName: "DeleteConfig",
			Handler:    _JudgmentService_DeleteConfig_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetConfigsByAgent",
			Handler:       _JudgmentService_GetConfigsByAgent_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "manage2store.proto",
}

// SendServiceClient is the client API for SendService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SendServiceClient interface {
	//
	// 1. 新增配置
	// 2. 根据id更新配置
	// 3. 根据id删除配置
	// 4. 根据id获取配置详情
	// 5. 根据agentID（或ip和local）查看配置，返回一个列表（流式响应）
	AddConfig(ctx context.Context, in *AddSendRequest, opts ...grpc.CallOption) (*BaseResponse, error)
	UpdateConfig(ctx context.Context, in *SendEntry, opts ...grpc.CallOption) (*BaseResponse, error)
	DeleteConfig(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*BaseResponse, error)
	GetConfigByID(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*SendConfigResponse, error)
	GetConfigsByAgent(ctx context.Context, in *AgentRequest, opts ...grpc.CallOption) (SendService_GetConfigsByAgentClient, error)
}

type sendServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSendServiceClient(cc grpc.ClientConnInterface) SendServiceClient {
	return &sendServiceClient{cc}
}

func (c *sendServiceClient) AddConfig(ctx context.Context, in *AddSendRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/judgment2store.SendService/AddConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sendServiceClient) UpdateConfig(ctx context.Context, in *SendEntry, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/judgment2store.SendService/UpdateConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sendServiceClient) DeleteConfig(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/judgment2store.SendService/DeleteConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sendServiceClient) GetConfigByID(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*SendConfigResponse, error) {
	out := new(SendConfigResponse)
	err := c.cc.Invoke(ctx, "/judgment2store.SendService/GetConfigByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sendServiceClient) GetConfigsByAgent(ctx context.Context, in *AgentRequest, opts ...grpc.CallOption) (SendService_GetConfigsByAgentClient, error) {
	stream, err := c.cc.NewStream(ctx, &SendService_ServiceDesc.Streams[0], "/judgment2store.SendService/GetConfigsByAgent", opts...)
	if err != nil {
		return nil, err
	}
	x := &sendServiceGetConfigsByAgentClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SendService_GetConfigsByAgentClient interface {
	Recv() (*SendConfigResponse, error)
	grpc.ClientStream
}

type sendServiceGetConfigsByAgentClient struct {
	grpc.ClientStream
}

func (x *sendServiceGetConfigsByAgentClient) Recv() (*SendConfigResponse, error) {
	m := new(SendConfigResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SendServiceServer is the server API for SendService service.
// All implementations must embed UnimplementedSendServiceServer
// for forward compatibility
type SendServiceServer interface {
	//
	// 1. 新增配置
	// 2. 根据id更新配置
	// 3. 根据id删除配置
	// 4. 根据id获取配置详情
	// 5. 根据agentID（或ip和local）查看配置，返回一个列表（流式响应）
	AddConfig(context.Context, *AddSendRequest) (*BaseResponse, error)
	UpdateConfig(context.Context, *SendEntry) (*BaseResponse, error)
	DeleteConfig(context.Context, *IDRequest) (*BaseResponse, error)
	GetConfigByID(context.Context, *IDRequest) (*SendConfigResponse, error)
	GetConfigsByAgent(*AgentRequest, SendService_GetConfigsByAgentServer) error
	mustEmbedUnimplementedSendServiceServer()
}

// UnimplementedSendServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSendServiceServer struct {
}

func (UnimplementedSendServiceServer) AddConfig(context.Context, *AddSendRequest) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddConfig not implemented")
}
func (UnimplementedSendServiceServer) UpdateConfig(context.Context, *SendEntry) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConfig not implemented")
}
func (UnimplementedSendServiceServer) DeleteConfig(context.Context, *IDRequest) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConfig not implemented")
}
func (UnimplementedSendServiceServer) GetConfigByID(context.Context, *IDRequest) (*SendConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfigByID not implemented")
}
func (UnimplementedSendServiceServer) GetConfigsByAgent(*AgentRequest, SendService_GetConfigsByAgentServer) error {
	return status.Errorf(codes.Unimplemented, "method GetConfigsByAgent not implemented")
}
func (UnimplementedSendServiceServer) mustEmbedUnimplementedSendServiceServer() {}

// UnsafeSendServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SendServiceServer will
// result in compilation errors.
type UnsafeSendServiceServer interface {
	mustEmbedUnimplementedSendServiceServer()
}

func RegisterSendServiceServer(s grpc.ServiceRegistrar, srv SendServiceServer) {
	s.RegisterService(&SendService_ServiceDesc, srv)
}

func _SendService_AddConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendServiceServer).AddConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/judgment2store.SendService/AddConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendServiceServer).AddConfig(ctx, req.(*AddSendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SendService_UpdateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEntry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendServiceServer).UpdateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/judgment2store.SendService/UpdateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendServiceServer).UpdateConfig(ctx, req.(*SendEntry))
	}
	return interceptor(ctx, in, info, handler)
}

func _SendService_DeleteConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendServiceServer).DeleteConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/judgment2store.SendService/DeleteConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendServiceServer).DeleteConfig(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SendService_GetConfigByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendServiceServer).GetConfigByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/judgment2store.SendService/GetConfigByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendServiceServer).GetConfigByID(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SendService_GetConfigsByAgent_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AgentRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SendServiceServer).GetConfigsByAgent(m, &sendServiceGetConfigsByAgentServer{stream})
}

type SendService_GetConfigsByAgentServer interface {
	Send(*SendConfigResponse) error
	grpc.ServerStream
}

type sendServiceGetConfigsByAgentServer struct {
	grpc.ServerStream
}

func (x *sendServiceGetConfigsByAgentServer) Send(m *SendConfigResponse) error {
	return x.ServerStream.SendMsg(m)
}

// SendService_ServiceDesc is the grpc.ServiceDesc for SendService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SendService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "judgment2store.SendService",
	HandlerType: (*SendServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddConfig",
			Handler:    _SendService_AddConfig_Handler,
		},
		{
			MethodName: "UpdateConfig",
			Handler:    _SendService_UpdateConfig_Handler,
		},
		{
			MethodName: "DeleteConfig",
			Handler:    _SendService_DeleteConfig_Handler,
		},
		{
			MethodName: "GetConfigByID",
			Handler:    _SendService_GetConfigByID_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetConfigsByAgent",
			Handler:       _SendService_GetConfigsByAgent_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "manage2store.proto",
}
