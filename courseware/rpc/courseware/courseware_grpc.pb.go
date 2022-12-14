// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package courseware

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

// CoursewareClient is the client API for Courseware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoursewareClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type coursewareClient struct {
	cc grpc.ClientConnInterface
}

func NewCoursewareClient(cc grpc.ClientConnInterface) CoursewareClient {
	return &coursewareClient{cc}
}

func (c *coursewareClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/user.Courseware/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coursewareClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/user.Courseware/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coursewareClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/user.Courseware/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coursewareClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/user.Courseware/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoursewareServer is the server API for Courseware service.
// All implementations must embed UnimplementedCoursewareServer
// for forward compatibility
type CoursewareServer interface {
	Add(context.Context, *AddRequest) (*AddResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	mustEmbedUnimplementedCoursewareServer()
}

// UnimplementedCoursewareServer must be embedded to have forward compatible implementations.
type UnimplementedCoursewareServer struct {
}

func (UnimplementedCoursewareServer) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedCoursewareServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCoursewareServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCoursewareServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCoursewareServer) mustEmbedUnimplementedCoursewareServer() {}

// UnsafeCoursewareServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoursewareServer will
// result in compilation errors.
type UnsafeCoursewareServer interface {
	mustEmbedUnimplementedCoursewareServer()
}

func RegisterCoursewareServer(s grpc.ServiceRegistrar, srv CoursewareServer) {
	s.RegisterService(&Courseware_ServiceDesc, srv)
}

func _Courseware_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoursewareServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Courseware/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoursewareServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Courseware_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoursewareServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Courseware/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoursewareServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Courseware_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoursewareServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Courseware/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoursewareServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Courseware_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoursewareServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Courseware/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoursewareServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Courseware_ServiceDesc is the grpc.ServiceDesc for Courseware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Courseware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.Courseware",
	HandlerType: (*CoursewareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Courseware_Add_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Courseware_Update_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Courseware_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Courseware_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "courseware.proto",
}
