// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: branches.proto

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

// BranchesServiceClient is the client API for BranchesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BranchesServiceClient interface {
	Create(ctx context.Context, in *CreateBranches, opts ...grpc.CallOption) (*Void, error)
	Update(ctx context.Context, in *UpdateBranches, opts ...grpc.CallOption) (*Void, error)
	Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error)
	GetById(ctx context.Context, in *ById, opts ...grpc.CallOption) (*BranchesRes, error)
	GetList(ctx context.Context, in *GetListBranchesReq, opts ...grpc.CallOption) (*GetListBranchesRes, error)
}

type branchesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBranchesServiceClient(cc grpc.ClientConnInterface) BranchesServiceClient {
	return &branchesServiceClient{cc}
}

func (c *branchesServiceClient) Create(ctx context.Context, in *CreateBranches, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/proto.BranchesService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *branchesServiceClient) Update(ctx context.Context, in *UpdateBranches, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/proto.BranchesService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *branchesServiceClient) Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/proto.BranchesService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *branchesServiceClient) GetById(ctx context.Context, in *ById, opts ...grpc.CallOption) (*BranchesRes, error) {
	out := new(BranchesRes)
	err := c.cc.Invoke(ctx, "/proto.BranchesService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *branchesServiceClient) GetList(ctx context.Context, in *GetListBranchesReq, opts ...grpc.CallOption) (*GetListBranchesRes, error) {
	out := new(GetListBranchesRes)
	err := c.cc.Invoke(ctx, "/proto.BranchesService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BranchesServiceServer is the server API for BranchesService service.
// All implementations must embed UnimplementedBranchesServiceServer
// for forward compatibility
type BranchesServiceServer interface {
	Create(context.Context, *CreateBranches) (*Void, error)
	Update(context.Context, *UpdateBranches) (*Void, error)
	Delete(context.Context, *ById) (*Void, error)
	GetById(context.Context, *ById) (*BranchesRes, error)
	GetList(context.Context, *GetListBranchesReq) (*GetListBranchesRes, error)
	mustEmbedUnimplementedBranchesServiceServer()
}

// UnimplementedBranchesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBranchesServiceServer struct {
}

func (UnimplementedBranchesServiceServer) Create(context.Context, *CreateBranches) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedBranchesServiceServer) Update(context.Context, *UpdateBranches) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedBranchesServiceServer) Delete(context.Context, *ById) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedBranchesServiceServer) GetById(context.Context, *ById) (*BranchesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedBranchesServiceServer) GetList(context.Context, *GetListBranchesReq) (*GetListBranchesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedBranchesServiceServer) mustEmbedUnimplementedBranchesServiceServer() {}

// UnsafeBranchesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BranchesServiceServer will
// result in compilation errors.
type UnsafeBranchesServiceServer interface {
	mustEmbedUnimplementedBranchesServiceServer()
}

func RegisterBranchesServiceServer(s grpc.ServiceRegistrar, srv BranchesServiceServer) {
	s.RegisterService(&BranchesService_ServiceDesc, srv)
}

func _BranchesService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBranches)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchesServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BranchesService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchesServiceServer).Create(ctx, req.(*CreateBranches))
	}
	return interceptor(ctx, in, info, handler)
}

func _BranchesService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBranches)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchesServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BranchesService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchesServiceServer).Update(ctx, req.(*UpdateBranches))
	}
	return interceptor(ctx, in, info, handler)
}

func _BranchesService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchesServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BranchesService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchesServiceServer).Delete(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _BranchesService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchesServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BranchesService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchesServiceServer).GetById(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _BranchesService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListBranchesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchesServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BranchesService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchesServiceServer).GetList(ctx, req.(*GetListBranchesReq))
	}
	return interceptor(ctx, in, info, handler)
}

// BranchesService_ServiceDesc is the grpc.ServiceDesc for BranchesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BranchesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.BranchesService",
	HandlerType: (*BranchesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _BranchesService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _BranchesService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BranchesService_Delete_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _BranchesService_GetById_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _BranchesService_GetList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "branches.proto",
}
