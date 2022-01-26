// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package greetpb

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

// GreetingServiceClient is the client API for GreetingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetingServiceClient interface {
	Greeting(ctx context.Context, in *GreetingRequest, opts ...grpc.CallOption) (*GreetingResponse, error)
}

type greetingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetingServiceClient(cc grpc.ClientConnInterface) GreetingServiceClient {
	return &greetingServiceClient{cc}
}

func (c *greetingServiceClient) Greeting(ctx context.Context, in *GreetingRequest, opts ...grpc.CallOption) (*GreetingResponse, error) {
	out := new(GreetingResponse)
	err := c.cc.Invoke(ctx, "/greet.GreetingService/Greeting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreetingServiceServer is the server API for GreetingService service.
// All implementations must embed UnimplementedGreetingServiceServer
// for forward compatibility
type GreetingServiceServer interface {
	Greeting(context.Context, *GreetingRequest) (*GreetingResponse, error)
	mustEmbedUnimplementedGreetingServiceServer()
}

// UnimplementedGreetingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGreetingServiceServer struct {
}

func (UnimplementedGreetingServiceServer) Greeting(context.Context, *GreetingRequest) (*GreetingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greeting not implemented")
}
func (UnimplementedGreetingServiceServer) mustEmbedUnimplementedGreetingServiceServer() {}

// UnsafeGreetingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetingServiceServer will
// result in compilation errors.
type UnsafeGreetingServiceServer interface {
	mustEmbedUnimplementedGreetingServiceServer()
}

func RegisterGreetingServiceServer(s grpc.ServiceRegistrar, srv GreetingServiceServer) {
	s.RegisterService(&GreetingService_ServiceDesc, srv)
}

func _GreetingService_Greeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetingServiceServer).Greeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.GreetingService/Greeting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetingServiceServer).Greeting(ctx, req.(*GreetingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GreetingService_ServiceDesc is the grpc.ServiceDesc for GreetingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreetingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreetingService",
	HandlerType: (*GreetingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greeting",
			Handler:    _GreetingService_Greeting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/greet.proto",
}
