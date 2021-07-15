// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v2

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

// ListerClient is the client API for Lister service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ListerClient interface {
	GetAllLists(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Lists, error)
	GetList(ctx context.Context, in *List, opts ...grpc.CallOption) (*List, error)
	AddList(ctx context.Context, in *List, opts ...grpc.CallOption) (*Empty, error)
	ChangeList(ctx context.Context, in *List, opts ...grpc.CallOption) (*Empty, error)
	DeleteList(ctx context.Context, in *List, opts ...grpc.CallOption) (*Empty, error)
}

type listerClient struct {
	cc grpc.ClientConnInterface
}

func NewListerClient(cc grpc.ClientConnInterface) ListerClient {
	return &listerClient{cc}
}

func (c *listerClient) GetAllLists(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Lists, error) {
	out := new(Lists)
	err := c.cc.Invoke(ctx, "/lister.v2.Lister/GetAllLists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listerClient) GetList(ctx context.Context, in *List, opts ...grpc.CallOption) (*List, error) {
	out := new(List)
	err := c.cc.Invoke(ctx, "/lister.v2.Lister/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listerClient) AddList(ctx context.Context, in *List, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/lister.v2.Lister/AddList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listerClient) ChangeList(ctx context.Context, in *List, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/lister.v2.Lister/ChangeList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listerClient) DeleteList(ctx context.Context, in *List, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/lister.v2.Lister/DeleteList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ListerServer is the server API for Lister service.
// All implementations must embed UnimplementedListerServer
// for forward compatibility
type ListerServer interface {
	GetAllLists(context.Context, *Empty) (*Lists, error)
	GetList(context.Context, *List) (*List, error)
	AddList(context.Context, *List) (*Empty, error)
	ChangeList(context.Context, *List) (*Empty, error)
	DeleteList(context.Context, *List) (*Empty, error)
	mustEmbedUnimplementedListerServer()
}

// UnimplementedListerServer must be embedded to have forward compatible implementations.
type UnimplementedListerServer struct {
}

func (UnimplementedListerServer) GetAllLists(context.Context, *Empty) (*Lists, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllLists not implemented")
}
func (UnimplementedListerServer) GetList(context.Context, *List) (*List, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedListerServer) AddList(context.Context, *List) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddList not implemented")
}
func (UnimplementedListerServer) ChangeList(context.Context, *List) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeList not implemented")
}
func (UnimplementedListerServer) DeleteList(context.Context, *List) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteList not implemented")
}
func (UnimplementedListerServer) mustEmbedUnimplementedListerServer() {}

// UnsafeListerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ListerServer will
// result in compilation errors.
type UnsafeListerServer interface {
	mustEmbedUnimplementedListerServer()
}

func RegisterListerServer(s grpc.ServiceRegistrar, srv ListerServer) {
	s.RegisterService(&Lister_ServiceDesc, srv)
}

func _Lister_GetAllLists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListerServer).GetAllLists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lister.v2.Lister/GetAllLists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListerServer).GetAllLists(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lister_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(List)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListerServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lister.v2.Lister/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListerServer).GetList(ctx, req.(*List))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lister_AddList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(List)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListerServer).AddList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lister.v2.Lister/AddList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListerServer).AddList(ctx, req.(*List))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lister_ChangeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(List)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListerServer).ChangeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lister.v2.Lister/ChangeList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListerServer).ChangeList(ctx, req.(*List))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lister_DeleteList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(List)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListerServer).DeleteList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lister.v2.Lister/DeleteList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListerServer).DeleteList(ctx, req.(*List))
	}
	return interceptor(ctx, in, info, handler)
}

// Lister_ServiceDesc is the grpc.ServiceDesc for Lister service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Lister_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lister.v2.Lister",
	HandlerType: (*ListerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllLists",
			Handler:    _Lister_GetAllLists_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _Lister_GetList_Handler,
		},
		{
			MethodName: "AddList",
			Handler:    _Lister_AddList_Handler,
		},
		{
			MethodName: "ChangeList",
			Handler:    _Lister_ChangeList_Handler,
		},
		{
			MethodName: "DeleteList",
			Handler:    _Lister_DeleteList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v2/lister.proto",
}
