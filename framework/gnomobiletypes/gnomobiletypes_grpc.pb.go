// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: gnomobiletypes.proto

package gnomobiletypes

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
	GnomobileService_CreateReply_FullMethodName = "/gomobile.v1.GnomobileService/CreateReply"
	GnomobileService_Hello_FullMethodName       = "/gomobile.v1.GnomobileService/Hello"
)

// GnomobileServiceClient is the client API for GnomobileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GnomobileServiceClient interface {
	// CreateReply creates a reply to the Gnoboard realms
	CreateReply(ctx context.Context, in *CreateReply_Request, opts ...grpc.CallOption) (*CreateReply_Reply, error)
	// Hello greets someone
	Hello(ctx context.Context, in *Hello_Request, opts ...grpc.CallOption) (*Hello_Reply, error)
}

type gnomobileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGnomobileServiceClient(cc grpc.ClientConnInterface) GnomobileServiceClient {
	return &gnomobileServiceClient{cc}
}

func (c *gnomobileServiceClient) CreateReply(ctx context.Context, in *CreateReply_Request, opts ...grpc.CallOption) (*CreateReply_Reply, error) {
	out := new(CreateReply_Reply)
	err := c.cc.Invoke(ctx, GnomobileService_CreateReply_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gnomobileServiceClient) Hello(ctx context.Context, in *Hello_Request, opts ...grpc.CallOption) (*Hello_Reply, error) {
	out := new(Hello_Reply)
	err := c.cc.Invoke(ctx, GnomobileService_Hello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GnomobileServiceServer is the server API for GnomobileService service.
// All implementations must embed UnimplementedGnomobileServiceServer
// for forward compatibility
type GnomobileServiceServer interface {
	// CreateReply creates a reply to the Gnoboard realms
	CreateReply(context.Context, *CreateReply_Request) (*CreateReply_Reply, error)
	// Hello greets someone
	Hello(context.Context, *Hello_Request) (*Hello_Reply, error)
	mustEmbedUnimplementedGnomobileServiceServer()
}

// UnimplementedGnomobileServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGnomobileServiceServer struct {
}

func (UnimplementedGnomobileServiceServer) CreateReply(context.Context, *CreateReply_Request) (*CreateReply_Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReply not implemented")
}
func (UnimplementedGnomobileServiceServer) Hello(context.Context, *Hello_Request) (*Hello_Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedGnomobileServiceServer) mustEmbedUnimplementedGnomobileServiceServer() {}

// UnsafeGnomobileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GnomobileServiceServer will
// result in compilation errors.
type UnsafeGnomobileServiceServer interface {
	mustEmbedUnimplementedGnomobileServiceServer()
}

func RegisterGnomobileServiceServer(s grpc.ServiceRegistrar, srv GnomobileServiceServer) {
	s.RegisterService(&GnomobileService_ServiceDesc, srv)
}

func _GnomobileService_CreateReply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReply_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GnomobileServiceServer).CreateReply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GnomobileService_CreateReply_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GnomobileServiceServer).CreateReply(ctx, req.(*CreateReply_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _GnomobileService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hello_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GnomobileServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GnomobileService_Hello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GnomobileServiceServer).Hello(ctx, req.(*Hello_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// GnomobileService_ServiceDesc is the grpc.ServiceDesc for GnomobileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GnomobileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gomobile.v1.GnomobileService",
	HandlerType: (*GnomobileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateReply",
			Handler:    _GnomobileService_CreateReply_Handler,
		},
		{
			MethodName: "Hello",
			Handler:    _GnomobileService_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gnomobiletypes.proto",
}
