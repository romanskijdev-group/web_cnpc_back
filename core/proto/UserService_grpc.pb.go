// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: service/UserService.proto

package protoobj

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
	// вход пользователя в аккаунт
	UserLoginAccount(ctx context.Context, in *UserAuthReqAccountReq, opts ...grpc.CallOption) (*LogInInfoRes, error)
	// получение профиля пользователя
	GetUserProfile(ctx context.Context, in *UsersMsg, opts ...grpc.CallOption) (*UsersMsg, error)
	// получение профилей пользователей
	GetUsersInfoList(ctx context.Context, in *UsersMsg, opts ...grpc.CallOption) (*UsersMsgList, error)
	// обновление информаций о пользователе
	UpdateUserProfile(ctx context.Context, in *UsersMsg, opts ...grpc.CallOption) (*UsersMsg, error)
	// создание пользователя
	CreateNewUser(ctx context.Context, in *UsersMsg, opts ...grpc.CallOption) (*UsersMsg, error)
	// удаление пользователя
	DeleteUser(ctx context.Context, in *UsersMsg, opts ...grpc.CallOption) (*Empty, error)
	// проверка временного кода входа телеграм бота пользователя (если упех то очистка кода из памяти)
	CheckCodeBotAuthUserValid(ctx context.Context, in *UserAuthReqAccountReq, opts ...grpc.CallOption) (*LogInInfoRes, error)
	// обновление аватара пользователя
	UpdateUserAvatarURL(ctx context.Context, in *UpdateUserAvatarURLReq, opts ...grpc.CallOption) (*Empty, error)
}

func (u UserServiceClient) UserLoginAccount(ctx context.Context, req *UserAuthReqAccountReq) (*LogInInfoRes, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceClient) GetUserProfile(ctx context.Context, msg *UsersMsg) (*UsersMsg, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceClient) GetUsersInfoList(ctx context.Context, msg *UsersMsg) (*UsersMsgList, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceClient) UpdateUserProfile(ctx context.Context, msg *UsersMsg) (*UsersMsg, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceClient) CreateNewUser(ctx context.Context, msg *UsersMsg) (*UsersMsg, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceClient) DeleteUser(ctx context.Context, msg *UsersMsg) (*Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceClient) CheckCodeBotAuthUserValid(ctx context.Context, req *UserAuthReqAccountReq) (*LogInInfoRes, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceClient) UpdateUserAvatarURL(ctx context.Context, req *UpdateUserAvatarURLReq) (*Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceClient) mustEmbedUnimplementedUserServiceServer() {
	//TODO implement me
	panic("implement me")
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) UserLoginAccount(ctx context.Context, in *UserAuthReqAccountReq, opts ...grpc.CallOption) (*LogInInfoRes, error) {
	out := new(LogInInfoRes)
	err := c.cc.Invoke(ctx, "/msg.UserService/UserLoginAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserProfile(ctx context.Context, in *UsersMsg, opts ...grpc.CallOption) (*UsersMsg, error) {
	out := new(UsersMsg)
	err := c.cc.Invoke(ctx, "/msg.UserService/GetUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUsersInfoList(ctx context.Context, in *UsersMsg, opts ...grpc.CallOption) (*UsersMsgList, error) {
	out := new(UsersMsgList)
	err := c.cc.Invoke(ctx, "/msg.UserService/GetUsersInfoList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserProfile(ctx context.Context, in *UsersMsg, opts ...grpc.CallOption) (*UsersMsg, error) {
	out := new(UsersMsg)
	err := c.cc.Invoke(ctx, "/msg.UserService/UpdateUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateNewUser(ctx context.Context, in *UsersMsg, opts ...grpc.CallOption) (*UsersMsg, error) {
	out := new(UsersMsg)
	err := c.cc.Invoke(ctx, "/msg.UserService/CreateNewUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *UsersMsg, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/msg.UserService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CheckCodeBotAuthUserValid(ctx context.Context, in *UserAuthReqAccountReq, opts ...grpc.CallOption) (*LogInInfoRes, error) {
	out := new(LogInInfoRes)
	err := c.cc.Invoke(ctx, "/msg.UserService/CheckCodeBotAuthUserValid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserAvatarURL(ctx context.Context, in *UpdateUserAvatarURLReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/msg.UserService/UpdateUserAvatarURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	// вход пользователя в аккаунт
	UserLoginAccount(context.Context, *UserAuthReqAccountReq) (*LogInInfoRes, error)
	// получение профиля пользователя
	GetUserProfile(context.Context, *UsersMsg) (*UsersMsg, error)
	// получение профилей пользователей
	GetUsersInfoList(context.Context, *UsersMsg) (*UsersMsgList, error)
	// обновление информаций о пользователе
	UpdateUserProfile(context.Context, *UsersMsg) (*UsersMsg, error)
	// создание пользователя
	CreateNewUser(context.Context, *UsersMsg) (*UsersMsg, error)
	// удаление пользователя
	DeleteUser(context.Context, *UsersMsg) (*Empty, error)
	// проверка временного кода входа телеграм бота пользователя (если упех то очистка кода из памяти)
	CheckCodeBotAuthUserValid(context.Context, *UserAuthReqAccountReq) (*LogInInfoRes, error)
	// обновление аватара пользователя
	UpdateUserAvatarURL(context.Context, *UpdateUserAvatarURLReq) (*Empty, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) UserLoginAccount(context.Context, *UserAuthReqAccountReq) (*LogInInfoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLoginAccount not implemented")
}
func (UnimplementedUserServiceServer) GetUserProfile(context.Context, *UsersMsg) (*UsersMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfile not implemented")
}
func (UnimplementedUserServiceServer) GetUsersInfoList(context.Context, *UsersMsg) (*UsersMsgList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersInfoList not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserProfile(context.Context, *UsersMsg) (*UsersMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserProfile not implemented")
}
func (UnimplementedUserServiceServer) CreateNewUser(context.Context, *UsersMsg) (*UsersMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewUser not implemented")
}
func (UnimplementedUserServiceServer) DeleteUser(context.Context, *UsersMsg) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServiceServer) CheckCodeBotAuthUserValid(context.Context, *UserAuthReqAccountReq) (*LogInInfoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckCodeBotAuthUserValid not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserAvatarURL(context.Context, *UpdateUserAvatarURLReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserAvatarURL not implemented")
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

func _UserService_UserLoginAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAuthReqAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserLoginAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.UserService/UserLoginAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserLoginAccount(ctx, req.(*UserAuthReqAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.UserService/GetUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserProfile(ctx, req.(*UsersMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUsersInfoList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUsersInfoList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.UserService/GetUsersInfoList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUsersInfoList(ctx, req.(*UsersMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.UserService/UpdateUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserProfile(ctx, req.(*UsersMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateNewUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateNewUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.UserService/CreateNewUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateNewUser(ctx, req.(*UsersMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.UserService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*UsersMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CheckCodeBotAuthUserValid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAuthReqAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CheckCodeBotAuthUserValid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.UserService/CheckCodeBotAuthUserValid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CheckCodeBotAuthUserValid(ctx, req.(*UserAuthReqAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserAvatarURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserAvatarURLReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserAvatarURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.UserService/UpdateUserAvatarURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserAvatarURL(ctx, req.(*UpdateUserAvatarURLReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "msg.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserLoginAccount",
			Handler:    _UserService_UserLoginAccount_Handler,
		},
		{
			MethodName: "GetUserProfile",
			Handler:    _UserService_GetUserProfile_Handler,
		},
		{
			MethodName: "GetUsersInfoList",
			Handler:    _UserService_GetUsersInfoList_Handler,
		},
		{
			MethodName: "UpdateUserProfile",
			Handler:    _UserService_UpdateUserProfile_Handler,
		},
		{
			MethodName: "CreateNewUser",
			Handler:    _UserService_CreateNewUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
		{
			MethodName: "CheckCodeBotAuthUserValid",
			Handler:    _UserService_CheckCodeBotAuthUserValid_Handler,
		},
		{
			MethodName: "UpdateUserAvatarURL",
			Handler:    _UserService_UpdateUserAvatarURL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/UserService.proto",
}
