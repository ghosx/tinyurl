// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/external/external.proto

package external

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

// ExternalClient is the client API for External service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExternalClient interface {
	CreateUrl(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	GetUrl(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type externalClient struct {
	cc grpc.ClientConnInterface
}

func NewExternalClient(cc grpc.ClientConnInterface) ExternalClient {
	return &externalClient{cc}
}

func (c *externalClient) CreateUrl(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/external.External/CreateUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalClient) GetUrl(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/external.External/GetUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExternalServer is the server API for External service.
// All implementations must embed UnimplementedExternalServer
// for forward compatibility
type ExternalServer interface {
	CreateUrl(context.Context, *CreateRequest) (*CreateResponse, error)
	GetUrl(context.Context, *GetRequest) (*GetResponse, error)
	mustEmbedUnimplementedExternalServer()
}

// UnimplementedExternalServer must be embedded to have forward compatible implementations.
type UnimplementedExternalServer struct {
}

func (UnimplementedExternalServer) CreateUrl(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUrl not implemented")
}
func (UnimplementedExternalServer) GetUrl(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUrl not implemented")
}
func (UnimplementedExternalServer) mustEmbedUnimplementedExternalServer() {}

// UnsafeExternalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExternalServer will
// result in compilation errors.
type UnsafeExternalServer interface {
	mustEmbedUnimplementedExternalServer()
}

func RegisterExternalServer(s grpc.ServiceRegistrar, srv ExternalServer) {
	s.RegisterService(&External_ServiceDesc, srv)
}

func _External_CreateUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalServer).CreateUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/external.External/CreateUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalServer).CreateUrl(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _External_GetUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalServer).GetUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/external.External/GetUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalServer).GetUrl(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// External_ServiceDesc is the grpc.ServiceDesc for External service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var External_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "external.External",
	HandlerType: (*ExternalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUrl",
			Handler:    _External_CreateUrl_Handler,
		},
		{
			MethodName: "GetUrl",
			Handler:    _External_GetUrl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/external/external.proto",
}
