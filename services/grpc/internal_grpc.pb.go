// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: internal/services/grpc/internal.proto

package grpc

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

// EsServiceClient is the client API for EsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EsServiceClient interface {
	TestUnary(ctx context.Context, in *SendUserSegmentRequest, opts ...grpc.CallOption) (*SendUserSegmentResponse, error)
}

type esServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEsServiceClient(cc grpc.ClientConnInterface) EsServiceClient {
	return &esServiceClient{cc}
}

func (c *esServiceClient) TestUnary(ctx context.Context, in *SendUserSegmentRequest, opts ...grpc.CallOption) (*SendUserSegmentResponse, error) {
	out := new(SendUserSegmentResponse)
	err := c.cc.Invoke(ctx, "/arman_task.EsService/TestUnary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EsServiceServer is the server API for EsService service.
// All implementations must embed UnimplementedEsServiceServer
// for forward compatibility
type EsServiceServer interface {
	TestUnary(context.Context, *SendUserSegmentRequest) (*SendUserSegmentResponse, error)
	mustEmbedUnimplementedEsServiceServer()
}

// UnimplementedEsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEsServiceServer struct {
}

func (UnimplementedEsServiceServer) TestUnary(context.Context, *SendUserSegmentRequest) (*SendUserSegmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestUnary not implemented")
}
func (UnimplementedEsServiceServer) mustEmbedUnimplementedEsServiceServer() {}

// UnsafeEsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EsServiceServer will
// result in compilation errors.
type UnsafeEsServiceServer interface {
	mustEmbedUnimplementedEsServiceServer()
}

func RegisterEsServiceServer(s grpc.ServiceRegistrar, srv EsServiceServer) {
	s.RegisterService(&EsService_ServiceDesc, srv)
}

func _EsService_TestUnary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendUserSegmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EsServiceServer).TestUnary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/arman_task.EsService/TestUnary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EsServiceServer).TestUnary(ctx, req.(*SendUserSegmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EsService_ServiceDesc is the grpc.ServiceDesc for EsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "arman_task.EsService",
	HandlerType: (*EsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestUnary",
			Handler:    _EsService_TestUnary_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/services/grpc/internal.proto",
}
