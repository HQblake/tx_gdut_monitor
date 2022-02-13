// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: manage2judgment.receivepb

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

// RuleUpdaterClient is the client API for RuleUpdater service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RuleUpdaterClient interface {
	Update(ctx context.Context, in *AgentRule, opts ...grpc.CallOption) (*Response, error)
}

type ruleUpdaterClient struct {
	cc grpc.ClientConnInterface
}

func NewRuleUpdaterClient(cc grpc.ClientConnInterface) RuleUpdaterClient {
	return &ruleUpdaterClient{cc}
}

func (c *ruleUpdaterClient) Update(ctx context.Context, in *AgentRule, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/receivepb.RuleUpdater/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RuleUpdaterServer is the server API for RuleUpdater service.
// All implementations must embed UnimplementedRuleUpdaterServer
// for forward compatibility
type RuleUpdaterServer interface {
	Update(context.Context, *AgentRule) (*Response, error)
	mustEmbedUnimplementedRuleUpdaterServer()
}

// UnimplementedRuleUpdaterServer must be embedded to have forward compatible implementations.
type UnimplementedRuleUpdaterServer struct {
}

func (UnimplementedRuleUpdaterServer) Update(context.Context, *AgentRule) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRuleUpdaterServer) mustEmbedUnimplementedRuleUpdaterServer() {}

// UnsafeRuleUpdaterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RuleUpdaterServer will
// result in compilation errors.
type UnsafeRuleUpdaterServer interface {
	mustEmbedUnimplementedRuleUpdaterServer()
}

func RegisterRuleUpdaterServer(s grpc.ServiceRegistrar, srv RuleUpdaterServer) {
	s.RegisterService(&RuleUpdater_ServiceDesc, srv)
}

func _RuleUpdater_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentRule)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleUpdaterServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/receivepb.RuleUpdater/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleUpdaterServer).Update(ctx, req.(*AgentRule))
	}
	return interceptor(ctx, in, info, handler)
}

// RuleUpdater_ServiceDesc is the grpc.ServiceDesc for RuleUpdater service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RuleUpdater_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "receivepb.RuleUpdater",
	HandlerType: (*RuleUpdaterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Update",
			Handler:    _RuleUpdater_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "manage2judgment.receivepb",
}
