// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: service/DialogService.proto

package protoobj

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DialogServiceClient is the client API for DialogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DialogServiceClient interface {
}

type dialogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDialogServiceClient(cc grpc.ClientConnInterface) DialogServiceClient {
	return &dialogServiceClient{cc}
}

// DialogServiceServer is the server API for DialogService service.
// All implementations must embed UnimplementedDialogServiceServer
// for forward compatibility
type DialogServiceServer interface {
	mustEmbedUnimplementedDialogServiceServer()
}

// UnimplementedDialogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDialogServiceServer struct {
}

func (UnimplementedDialogServiceServer) mustEmbedUnimplementedDialogServiceServer() {}

// UnsafeDialogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DialogServiceServer will
// result in compilation errors.
type UnsafeDialogServiceServer interface {
	mustEmbedUnimplementedDialogServiceServer()
}

func RegisterDialogServiceServer(s grpc.ServiceRegistrar, srv DialogServiceServer) {
	s.RegisterService(&DialogService_ServiceDesc, srv)
}

// DialogService_ServiceDesc is the grpc.ServiceDesc for DialogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DialogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "msg.DialogService",
	HandlerType: (*DialogServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "service/DialogService.proto",
}
