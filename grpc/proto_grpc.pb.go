// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: grpc/proto.proto

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

const (
	Myservice_MethodName_FullMethodName = "/handin4.myservice/MethodName"
)

// MyserviceClient is the client API for Myservice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MyserviceClient interface {
	MethodName(ctx context.Context, in *EnterRequest, opts ...grpc.CallOption) (*Exit, error)
}

type myserviceClient struct {
	cc grpc.ClientConnInterface
}

func NewMyserviceClient(cc grpc.ClientConnInterface) MyserviceClient {
	return &myserviceClient{cc}
}

func (c *myserviceClient) MethodName(ctx context.Context, in *EnterRequest, opts ...grpc.CallOption) (*Exit, error) {
	out := new(Exit)
	err := c.cc.Invoke(ctx, Myservice_MethodName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MyserviceServer is the server API for Myservice service.
// All implementations must embed UnimplementedMyserviceServer
// for forward compatibility
type MyserviceServer interface {
	MethodName(context.Context, *EnterRequest) (*Exit, error)
	mustEmbedUnimplementedMyserviceServer()
}

// UnimplementedMyserviceServer must be embedded to have forward compatible implementations.
type UnimplementedMyserviceServer struct {
}

func (UnimplementedMyserviceServer) MethodName(context.Context, *EnterRequest) (*Exit, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MethodName not implemented")
}
func (UnimplementedMyserviceServer) mustEmbedUnimplementedMyserviceServer() {}

// UnsafeMyserviceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MyserviceServer will
// result in compilation errors.
type UnsafeMyserviceServer interface {
	mustEmbedUnimplementedMyserviceServer()
}

func RegisterMyserviceServer(s grpc.ServiceRegistrar, srv MyserviceServer) {
	s.RegisterService(&Myservice_ServiceDesc, srv)
}

func _Myservice_MethodName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyserviceServer).MethodName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Myservice_MethodName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyserviceServer).MethodName(ctx, req.(*EnterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Myservice_ServiceDesc is the grpc.ServiceDesc for Myservice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Myservice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "handin4.myservice",
	HandlerType: (*MyserviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MethodName",
			Handler:    _Myservice_MethodName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/proto.proto",
}
