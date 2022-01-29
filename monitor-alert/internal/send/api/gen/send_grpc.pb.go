// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package sendpb

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

// SendServiceClient is the client API for SendService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SendServiceClient interface {
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Response, error)
	Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*Response, error)
}

type sendServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSendServiceClient(cc grpc.ClientConnInterface) SendServiceClient {
	return &sendServiceClient{cc}
}

func (c *sendServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.SendService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sendServiceClient) Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.SendService/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SendServiceServer is the server API for SendService service.
// All implementations must embed UnimplementedSendServiceServer
// for forward compatibility
type SendServiceServer interface {
	Update(context.Context, *UpdateRequest) (*Response, error)
	Init(context.Context, *InitRequest) (*Response, error)
	mustEmbedUnimplementedSendServiceServer()
}

// UnimplementedSendServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSendServiceServer struct {
}

func (UnimplementedSendServiceServer) Update(context.Context, *UpdateRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSendServiceServer) Init(context.Context, *InitRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
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

func _SendService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SendService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SendService_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendServiceServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SendService/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendServiceServer).Init(ctx, req.(*InitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SendService_ServiceDesc is the grpc.ServiceDesc for SendService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SendService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.SendService",
	HandlerType: (*SendServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Update",
			Handler:    _SendService_Update_Handler,
		},
		{
			MethodName: "Init",
			Handler:    _SendService_Init_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "send.proto",
}
