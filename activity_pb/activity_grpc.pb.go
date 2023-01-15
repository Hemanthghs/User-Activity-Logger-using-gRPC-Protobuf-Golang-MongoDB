// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: activity_pb/activity.proto

package activity_pb

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	UserAdd(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	ActivityAdd(ctx context.Context, in *ActivityRequest, opts ...grpc.CallOption) (*ActivityResponse, error)
	ActivityIsValid(ctx context.Context, in *ActivityIsValidRequest, opts ...grpc.CallOption) (*ActivityIsValidResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	ActivityIsDone(ctx context.Context, in *ActivityIsDoneRequest, opts ...grpc.CallOption) (*ActivityIsDoneResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	GetActivity(ctx context.Context, in *GetActivityRequest, opts ...grpc.CallOption) (*GetActivityResponse, error)
	RemoveUser(ctx context.Context, in *RemoveUserRequest, opts ...grpc.CallOption) (*RemoveUserResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) UserAdd(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/activity_pb.UserService/UserAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ActivityAdd(ctx context.Context, in *ActivityRequest, opts ...grpc.CallOption) (*ActivityResponse, error) {
	out := new(ActivityResponse)
	err := c.cc.Invoke(ctx, "/activity_pb.UserService/ActivityAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ActivityIsValid(ctx context.Context, in *ActivityIsValidRequest, opts ...grpc.CallOption) (*ActivityIsValidResponse, error) {
	out := new(ActivityIsValidResponse)
	err := c.cc.Invoke(ctx, "/activity_pb.UserService/ActivityIsValid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, "/activity_pb.UserService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ActivityIsDone(ctx context.Context, in *ActivityIsDoneRequest, opts ...grpc.CallOption) (*ActivityIsDoneResponse, error) {
	out := new(ActivityIsDoneResponse)
	err := c.cc.Invoke(ctx, "/activity_pb.UserService/ActivityIsDone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/activity_pb.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetActivity(ctx context.Context, in *GetActivityRequest, opts ...grpc.CallOption) (*GetActivityResponse, error) {
	out := new(GetActivityResponse)
	err := c.cc.Invoke(ctx, "/activity_pb.UserService/GetActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RemoveUser(ctx context.Context, in *RemoveUserRequest, opts ...grpc.CallOption) (*RemoveUserResponse, error) {
	out := new(RemoveUserResponse)
	err := c.cc.Invoke(ctx, "/activity_pb.UserService/RemoveUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	UserAdd(context.Context, *UserRequest) (*UserResponse, error)
	ActivityAdd(context.Context, *ActivityRequest) (*ActivityResponse, error)
	ActivityIsValid(context.Context, *ActivityIsValidRequest) (*ActivityIsValidResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	ActivityIsDone(context.Context, *ActivityIsDoneRequest) (*ActivityIsDoneResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	GetActivity(context.Context, *GetActivityRequest) (*GetActivityResponse, error)
	RemoveUser(context.Context, *RemoveUserRequest) (*RemoveUserResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) UserAdd(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserAdd not implemented")
}
func (UnimplementedUserServiceServer) ActivityAdd(context.Context, *ActivityRequest) (*ActivityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivityAdd not implemented")
}
func (UnimplementedUserServiceServer) ActivityIsValid(context.Context, *ActivityIsValidRequest) (*ActivityIsValidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivityIsValid not implemented")
}
func (UnimplementedUserServiceServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserServiceServer) ActivityIsDone(context.Context, *ActivityIsDoneRequest) (*ActivityIsDoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivityIsDone not implemented")
}
func (UnimplementedUserServiceServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServiceServer) GetActivity(context.Context, *GetActivityRequest) (*GetActivityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActivity not implemented")
}
func (UnimplementedUserServiceServer) RemoveUser(context.Context, *RemoveUserRequest) (*RemoveUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUser not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_UserAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activity_pb.UserService/UserAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserAdd(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ActivityAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ActivityAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activity_pb.UserService/ActivityAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ActivityAdd(ctx, req.(*ActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ActivityIsValid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityIsValidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ActivityIsValid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activity_pb.UserService/ActivityIsValid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ActivityIsValid(ctx, req.(*ActivityIsValidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activity_pb.UserService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ActivityIsDone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityIsDoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ActivityIsDone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activity_pb.UserService/ActivityIsDone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ActivityIsDone(ctx, req.(*ActivityIsDoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activity_pb.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activity_pb.UserService/GetActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetActivity(ctx, req.(*GetActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RemoveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RemoveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activity_pb.UserService/RemoveUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RemoveUser(ctx, req.(*RemoveUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "activity_pb.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserAdd",
			Handler:    _UserService_UserAdd_Handler,
		},
		{
			MethodName: "ActivityAdd",
			Handler:    _UserService_ActivityAdd_Handler,
		},
		{
			MethodName: "ActivityIsValid",
			Handler:    _UserService_ActivityIsValid_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
		{
			MethodName: "ActivityIsDone",
			Handler:    _UserService_ActivityIsDone_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "GetActivity",
			Handler:    _UserService_GetActivity_Handler,
		},
		{
			MethodName: "RemoveUser",
			Handler:    _UserService_RemoveUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "activity_pb/activity.proto",
}
