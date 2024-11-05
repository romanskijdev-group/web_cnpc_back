// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: service/NotificationService.proto

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

// NotificationServiceProtoClient is the client API for NotificationServiceProto service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationServiceProtoClient interface {
	// Уведомление пользователя
	NotifyUser(ctx context.Context, in *NotifyParams, opts ...grpc.CallOption) (*Empty, error)
}

type notificationServiceProtoClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationServiceProtoClient(cc grpc.ClientConnInterface) NotificationServiceProtoClient {
	return &notificationServiceProtoClient{cc}
}

func (c *notificationServiceProtoClient) NotifyUser(ctx context.Context, in *NotifyParams, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/msg.NotificationServiceProto/NotifyUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServiceProtoServer is the server API for NotificationServiceProto service.
// All implementations must embed UnimplementedNotificationServiceProtoServer
// for forward compatibility
type NotificationServiceProtoServer interface {
	// Уведомление пользователя
	NotifyUser(context.Context, *NotifyParams) (*Empty, error)
	mustEmbedUnimplementedNotificationServiceProtoServer()
}

// UnimplementedNotificationServiceProtoServer must be embedded to have forward compatible implementations.
type UnimplementedNotificationServiceProtoServer struct {
}

func (UnimplementedNotificationServiceProtoServer) NotifyUser(context.Context, *NotifyParams) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyUser not implemented")
}
func (UnimplementedNotificationServiceProtoServer) mustEmbedUnimplementedNotificationServiceProtoServer() {
}

// UnsafeNotificationServiceProtoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationServiceProtoServer will
// result in compilation errors.
type UnsafeNotificationServiceProtoServer interface {
	mustEmbedUnimplementedNotificationServiceProtoServer()
}

func RegisterNotificationServiceProtoServer(s grpc.ServiceRegistrar, srv NotificationServiceProtoServer) {
	s.RegisterService(&NotificationServiceProto_ServiceDesc, srv)
}

func _NotificationServiceProto_NotifyUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceProtoServer).NotifyUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/msg.NotificationServiceProto/NotifyUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceProtoServer).NotifyUser(ctx, req.(*NotifyParams))
	}
	return interceptor(ctx, in, info, handler)
}

// NotificationServiceProto_ServiceDesc is the grpc.ServiceDesc for NotificationServiceProto service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotificationServiceProto_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "msg.NotificationServiceProto",
	HandlerType: (*NotificationServiceProtoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NotifyUser",
			Handler:    _NotificationServiceProto_NotifyUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/NotificationService.proto",
}
