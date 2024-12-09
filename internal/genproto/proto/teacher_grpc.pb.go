// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: teacher.proto

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

// TeacherServiceClient is the client API for TeacherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TeacherServiceClient interface {
	Create(ctx context.Context, in *CreateTeacher, opts ...grpc.CallOption) (*Void, error)
	Update(ctx context.Context, in *UpdateTeacher, opts ...grpc.CallOption) (*Void, error)
	Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error)
	GetById(ctx context.Context, in *ById, opts ...grpc.CallOption) (*TeacherRes, error)
	GetList(ctx context.Context, in *GetListTeacherReq, opts ...grpc.CallOption) (*GetListTeacherRes, error)
}

type teacherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTeacherServiceClient(cc grpc.ClientConnInterface) TeacherServiceClient {
	return &teacherServiceClient{cc}
}

func (c *teacherServiceClient) Create(ctx context.Context, in *CreateTeacher, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/proto.TeacherService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teacherServiceClient) Update(ctx context.Context, in *UpdateTeacher, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/proto.TeacherService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teacherServiceClient) Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/proto.TeacherService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teacherServiceClient) GetById(ctx context.Context, in *ById, opts ...grpc.CallOption) (*TeacherRes, error) {
	out := new(TeacherRes)
	err := c.cc.Invoke(ctx, "/proto.TeacherService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teacherServiceClient) GetList(ctx context.Context, in *GetListTeacherReq, opts ...grpc.CallOption) (*GetListTeacherRes, error) {
	out := new(GetListTeacherRes)
	err := c.cc.Invoke(ctx, "/proto.TeacherService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeacherServiceServer is the server API for TeacherService service.
// All implementations must embed UnimplementedTeacherServiceServer
// for forward compatibility
type TeacherServiceServer interface {
	Create(context.Context, *CreateTeacher) (*Void, error)
	Update(context.Context, *UpdateTeacher) (*Void, error)
	Delete(context.Context, *ById) (*Void, error)
	GetById(context.Context, *ById) (*TeacherRes, error)
	GetList(context.Context, *GetListTeacherReq) (*GetListTeacherRes, error)
	mustEmbedUnimplementedTeacherServiceServer()
}

// UnimplementedTeacherServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTeacherServiceServer struct {
}

func (UnimplementedTeacherServiceServer) Create(context.Context, *CreateTeacher) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTeacherServiceServer) Update(context.Context, *UpdateTeacher) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTeacherServiceServer) Delete(context.Context, *ById) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTeacherServiceServer) GetById(context.Context, *ById) (*TeacherRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedTeacherServiceServer) GetList(context.Context, *GetListTeacherReq) (*GetListTeacherRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedTeacherServiceServer) mustEmbedUnimplementedTeacherServiceServer() {}

// UnsafeTeacherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TeacherServiceServer will
// result in compilation errors.
type UnsafeTeacherServiceServer interface {
	mustEmbedUnimplementedTeacherServiceServer()
}

func RegisterTeacherServiceServer(s grpc.ServiceRegistrar, srv TeacherServiceServer) {
	s.RegisterService(&TeacherService_ServiceDesc, srv)
}

func _TeacherService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTeacher)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeacherServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TeacherService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeacherServiceServer).Create(ctx, req.(*CreateTeacher))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeacherService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTeacher)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeacherServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TeacherService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeacherServiceServer).Update(ctx, req.(*UpdateTeacher))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeacherService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeacherServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TeacherService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeacherServiceServer).Delete(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeacherService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeacherServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TeacherService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeacherServiceServer).GetById(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeacherService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListTeacherReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeacherServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TeacherService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeacherServiceServer).GetList(ctx, req.(*GetListTeacherReq))
	}
	return interceptor(ctx, in, info, handler)
}

// TeacherService_ServiceDesc is the grpc.ServiceDesc for TeacherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TeacherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TeacherService",
	HandlerType: (*TeacherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TeacherService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TeacherService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TeacherService_Delete_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _TeacherService_GetById_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _TeacherService_GetList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teacher.proto",
}
