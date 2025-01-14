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

// UserAccountServiceProtoClient is the client API for UserAccountServiceProto service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserAccountServiceProtoClient interface {
	// вход пользователя в аккаунт
	UserLoginAccount(ctx context.Context, in *UserAuthReqAccountReq, opts ...grpc.CallOption) (*LogInInfoRes, error)
	// получение профиля пользователя
	GetUserProfile(ctx context.Context, in *UsersMsg, opts ...grpc.CallOption) (*UsersMsg, error)
	// получение профилей пользователей
	GetUsersInfoList(ctx context.Context, in *UsersMsgReq, opts ...grpc.CallOption) (*UsersMsgList, error)
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
	// получение уведомлений пользователя
	GetUserAlerts(ctx context.Context, in *UserAlertMsg, opts ...grpc.CallOption) (*UserAlertMsgList, error)
	// изменение уведомлений
	UpdateUserAlerts(ctx context.Context, in *UserAlertMsg, opts ...grpc.CallOption) (*UserAlertMsgList, error)
}

type userAccountServiceProtoClient struct {
	cc grpc.ClientConnInterface
}

func NewUserAccountServiceProtoClient(cc grpc.ClientConnInterface) UserAccountServiceProtoClient {
	return &userAccountServiceProtoClient{cc}
}

func (c *userAccountServiceProtoClient) UserLoginAccount(ctx context.Context, in *UserAuthReqAccountReq, opts ...grpc.CallOption) (*LogInInfoRes, error) {
	out := new(LogInInfoRes)
	err := c.cc.Invoke(ctx, "/msg.UserAccountServiceProto/UserLoginAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccountServiceProtoClient) GetUserProfile(ctx context.Context, in *UsersMsg, opts ...grpc.CallOption) (*UsersMsg, error) {
	out := new(UsersMsg)
	err := c.cc.Invoke(ctx, "/msg.UserAccountServiceProto/GetUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccountServiceProtoClient) GetUsersInfoList(ctx context.Context, in *UsersMsgReq, opts ...grpc.CallOption) (*UsersMsgList, error) {
	out := new(UsersMsgList)
	err := c.cc.Invoke(ctx, "/msg.UserAccountServiceProto/GetUsersInfoList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccountServiceProtoClient) UpdateUserProfile(ctx context.Context, in *UsersMsg, opts ...grpc.CallOption) (*UsersMsg, error) {
	out := new(UsersMsg)
	err := c.cc.Invoke(ctx, "/msg.UserAccountServiceProto/UpdateUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccountServiceProtoClient) CreateNewUser(ctx context.Context, in *UsersMsg, opts ...grpc.CallOption) (*UsersMsg, error) {
	out := new(UsersMsg)
	err := c.cc.Invoke(ctx, "/msg.UserAccountServiceProto/CreateNewUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccountServiceProtoClient) DeleteUser(ctx context.Context, in *UsersMsg, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/msg.UserAccountServiceProto/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccountServiceProtoClient) CheckCodeBotAuthUserValid(ctx context.Context, in *UserAuthReqAccountReq, opts ...grpc.CallOption) (*LogInInfoRes, error) {
	out := new(LogInInfoRes)
	err := c.cc.Invoke(ctx, "/msg.UserAccountServiceProto/CheckCodeBotAuthUserValid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccountServiceProtoClient) UpdateUserAvatarURL(ctx context.Context, in *UpdateUserAvatarURLReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/msg.UserAccountServiceProto/UpdateUserAvatarURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccountServiceProtoClient) GetUserAlerts(ctx context.Context, in *UserAlertMsg, opts ...grpc.CallOption) (*UserAlertMsgList, error) {
	out := new(UserAlertMsgList)
	err := c.cc.Invoke(ctx, "/msg.UserAccountServiceProto/GetUserAlerts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAccountServiceProtoClient) UpdateUserAlerts(ctx context.Context, in *UserAlertMsg, opts ...grpc.CallOption) (*UserAlertMsgList, error) {
	out := new(UserAlertMsgList)
	err := c.cc.Invoke(ctx, "/msg.UserAccountServiceProto/UpdateUserAlerts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserAccountServiceProtoServer is the server API for UserAccountServiceProto service.
// All implementations must embed UnimplementedUserAccountServiceProtoServer
// for forward compatibility
type UserAccountServiceProtoServer interface {
	// вход пользователя в аккаунт
	UserLoginAccount(context.Context, *UserAuthReqAccountReq) (*LogInInfoRes, error)
	// получение профиля пользователя
	GetUserProfile(context.Context, *UsersMsg) (*UsersMsg, error)
	// получение профилей пользователей
	GetUsersInfoList(context.Context, *UsersMsgReq) (*UsersMsgList, error)
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
	// получение уведомлений пользователя
	GetUserAlerts(context.Context, *UserAlertMsg) (*UserAlertMsgList, error)
	// изменение уведомлений
	UpdateUserAlerts(context.Context, *UserAlertMsg) (*UserAlertMsgList, error)
	mustEmbedUnimplementedUserAccountServiceProtoServer()
}

// UnimplementedUserAccountServiceProtoServer must be embedded to have forward compatible implementations.
type UnimplementedUserAccountServiceProtoServer struct {
}

func (UnimplementedUserAccountServiceProtoServer) UserLoginAccount(context.Context, *UserAuthReqAccountReq) (*LogInInfoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLoginAccount not implemented")
}
func (UnimplementedUserAccountServiceProtoServer) GetUserProfile(context.Context, *UsersMsg) (*UsersMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfile not implemented")
}
func (UnimplementedUserAccountServiceProtoServer) GetUsersInfoList(context.Context, *UsersMsgReq) (*UsersMsgList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersInfoList not implemented")
}
func (UnimplementedUserAccountServiceProtoServer) UpdateUserProfile(context.Context, *UsersMsg) (*UsersMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserProfile not implemented")
}
func (UnimplementedUserAccountServiceProtoServer) CreateNewUser(context.Context, *UsersMsg) (*UsersMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewUser not implemented")
}
func (UnimplementedUserAccountServiceProtoServer) DeleteUser(context.Context, *UsersMsg) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserAccountServiceProtoServer) CheckCodeBotAuthUserValid(context.Context, *UserAuthReqAccountReq) (*LogInInfoRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckCodeBotAuthUserValid not implemented")
}
func (UnimplementedUserAccountServiceProtoServer) UpdateUserAvatarURL(context.Context, *UpdateUserAvatarURLReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserAvatarURL not implemented")
}
func (UnimplementedUserAccountServiceProtoServer) GetUserAlerts(context.Context, *UserAlertMsg) (*UserAlertMsgList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAlerts not implemented")
}
func (UnimplementedUserAccountServiceProtoServer) UpdateUserAlerts(context.Context, *UserAlertMsg) (*UserAlertMsgList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserAlerts not implemented")
}
func (UnimplementedUserAccountServiceProtoServer) mustEmbedUnimplementedUserAccountServiceProtoServer() {
}

// UnsafeUserAccountServiceProtoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserAccountServiceProtoServer will
// result in compilation errors.
type UnsafeUserAccountServiceProtoServer interface {
	mustEmbedUnimplementedUserAccountServiceProtoServer()
}

func RegisterUserAccountServiceProtoServer(s grpc.ServiceRegistrar, srv UserAccountServiceProtoServer) {
	s.RegisterService(&UserAccountServiceProto_ServiceDesc, srv)
}

func _UserAccountServiceProto_UserLoginAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAuthReqAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccountServiceProtoServer).UserLoginAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.UserAccountServiceProto/UserLoginAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccountServiceProtoServer).UserLoginAccount(ctx, req.(*UserAuthReqAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccountServiceProto_GetUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccountServiceProtoServer).GetUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.UserAccountServiceProto/GetUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccountServiceProtoServer).GetUserProfile(ctx, req.(*UsersMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccountServiceProto_GetUsersInfoList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccountServiceProtoServer).GetUsersInfoList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.UserAccountServiceProto/GetUsersInfoList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccountServiceProtoServer).GetUsersInfoList(ctx, req.(*UsersMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccountServiceProto_UpdateUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccountServiceProtoServer).UpdateUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.UserAccountServiceProto/UpdateUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccountServiceProtoServer).UpdateUserProfile(ctx, req.(*UsersMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccountServiceProto_CreateNewUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccountServiceProtoServer).CreateNewUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.UserAccountServiceProto/CreateNewUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccountServiceProtoServer).CreateNewUser(ctx, req.(*UsersMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccountServiceProto_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccountServiceProtoServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.UserAccountServiceProto/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccountServiceProtoServer).DeleteUser(ctx, req.(*UsersMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccountServiceProto_CheckCodeBotAuthUserValid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAuthReqAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccountServiceProtoServer).CheckCodeBotAuthUserValid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.UserAccountServiceProto/CheckCodeBotAuthUserValid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccountServiceProtoServer).CheckCodeBotAuthUserValid(ctx, req.(*UserAuthReqAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccountServiceProto_UpdateUserAvatarURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserAvatarURLReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccountServiceProtoServer).UpdateUserAvatarURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.UserAccountServiceProto/UpdateUserAvatarURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccountServiceProtoServer).UpdateUserAvatarURL(ctx, req.(*UpdateUserAvatarURLReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccountServiceProto_GetUserAlerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAlertMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccountServiceProtoServer).GetUserAlerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.UserAccountServiceProto/GetUserAlerts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccountServiceProtoServer).GetUserAlerts(ctx, req.(*UserAlertMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAccountServiceProto_UpdateUserAlerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAlertMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccountServiceProtoServer).UpdateUserAlerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.UserAccountServiceProto/UpdateUserAlerts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccountServiceProtoServer).UpdateUserAlerts(ctx, req.(*UserAlertMsg))
	}
	return interceptor(ctx, in, info, handler)
}

// UserAccountServiceProto_ServiceDesc is the grpc.ServiceDesc for UserAccountServiceProto service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserAccountServiceProto_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "msg.UserAccountServiceProto",
	HandlerType: (*UserAccountServiceProtoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserLoginAccount",
			Handler:    _UserAccountServiceProto_UserLoginAccount_Handler,
		},
		{
			MethodName: "GetUserProfile",
			Handler:    _UserAccountServiceProto_GetUserProfile_Handler,
		},
		{
			MethodName: "GetUsersInfoList",
			Handler:    _UserAccountServiceProto_GetUsersInfoList_Handler,
		},
		{
			MethodName: "UpdateUserProfile",
			Handler:    _UserAccountServiceProto_UpdateUserProfile_Handler,
		},
		{
			MethodName: "CreateNewUser",
			Handler:    _UserAccountServiceProto_CreateNewUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserAccountServiceProto_DeleteUser_Handler,
		},
		{
			MethodName: "CheckCodeBotAuthUserValid",
			Handler:    _UserAccountServiceProto_CheckCodeBotAuthUserValid_Handler,
		},
		{
			MethodName: "UpdateUserAvatarURL",
			Handler:    _UserAccountServiceProto_UpdateUserAvatarURL_Handler,
		},
		{
			MethodName: "GetUserAlerts",
			Handler:    _UserAccountServiceProto_GetUserAlerts_Handler,
		},
		{
			MethodName: "UpdateUserAlerts",
			Handler:    _UserAccountServiceProto_UpdateUserAlerts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/UserService.proto",
}
