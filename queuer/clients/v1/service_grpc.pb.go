// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: queuer/clients/v1/service.proto

package v1

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

const (
	ClientService_RegisterClient_FullMethodName      = "/queuer.clients.v1.ClientService/RegisterClient"
	ClientService_PublishClientStats_FullMethodName  = "/queuer.clients.v1.ClientService/PublishClientStats"
	ClientService_SubscribeToCommands_FullMethodName = "/queuer.clients.v1.ClientService/SubscribeToCommands"
	ClientService_AckCommand_FullMethodName          = "/queuer.clients.v1.ClientService/AckCommand"
)

// ClientServiceClient is the client API for ClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientServiceClient interface {
	RegisterClient(ctx context.Context, in *RegisterClientRequest, opts ...grpc.CallOption) (*RegisterClientResponse, error)
	PublishClientStats(ctx context.Context, opts ...grpc.CallOption) (ClientService_PublishClientStatsClient, error)
	SubscribeToCommands(ctx context.Context, in *SubscribeToCommandsRequest, opts ...grpc.CallOption) (ClientService_SubscribeToCommandsClient, error)
	AckCommand(ctx context.Context, in *AckCommandRequest, opts ...grpc.CallOption) (*AckCommandResponse, error)
}

type clientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientServiceClient(cc grpc.ClientConnInterface) ClientServiceClient {
	return &clientServiceClient{cc}
}

func (c *clientServiceClient) RegisterClient(ctx context.Context, in *RegisterClientRequest, opts ...grpc.CallOption) (*RegisterClientResponse, error) {
	out := new(RegisterClientResponse)
	err := c.cc.Invoke(ctx, ClientService_RegisterClient_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) PublishClientStats(ctx context.Context, opts ...grpc.CallOption) (ClientService_PublishClientStatsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClientService_ServiceDesc.Streams[0], ClientService_PublishClientStats_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &clientServicePublishClientStatsClient{stream}
	return x, nil
}

type ClientService_PublishClientStatsClient interface {
	Send(*PublishClientStatsRequest) error
	CloseAndRecv() (*PublishClientStatsResponse, error)
	grpc.ClientStream
}

type clientServicePublishClientStatsClient struct {
	grpc.ClientStream
}

func (x *clientServicePublishClientStatsClient) Send(m *PublishClientStatsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clientServicePublishClientStatsClient) CloseAndRecv() (*PublishClientStatsResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PublishClientStatsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clientServiceClient) SubscribeToCommands(ctx context.Context, in *SubscribeToCommandsRequest, opts ...grpc.CallOption) (ClientService_SubscribeToCommandsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClientService_ServiceDesc.Streams[1], ClientService_SubscribeToCommands_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &clientServiceSubscribeToCommandsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ClientService_SubscribeToCommandsClient interface {
	Recv() (*SubscribeToCommandsResponse, error)
	grpc.ClientStream
}

type clientServiceSubscribeToCommandsClient struct {
	grpc.ClientStream
}

func (x *clientServiceSubscribeToCommandsClient) Recv() (*SubscribeToCommandsResponse, error) {
	m := new(SubscribeToCommandsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clientServiceClient) AckCommand(ctx context.Context, in *AckCommandRequest, opts ...grpc.CallOption) (*AckCommandResponse, error) {
	out := new(AckCommandResponse)
	err := c.cc.Invoke(ctx, ClientService_AckCommand_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientServiceServer is the server API for ClientService service.
// All implementations must embed UnimplementedClientServiceServer
// for forward compatibility
type ClientServiceServer interface {
	RegisterClient(context.Context, *RegisterClientRequest) (*RegisterClientResponse, error)
	PublishClientStats(ClientService_PublishClientStatsServer) error
	SubscribeToCommands(*SubscribeToCommandsRequest, ClientService_SubscribeToCommandsServer) error
	AckCommand(context.Context, *AckCommandRequest) (*AckCommandResponse, error)
	mustEmbedUnimplementedClientServiceServer()
}

// UnimplementedClientServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClientServiceServer struct {
}

func (UnimplementedClientServiceServer) RegisterClient(context.Context, *RegisterClientRequest) (*RegisterClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterClient not implemented")
}
func (UnimplementedClientServiceServer) PublishClientStats(ClientService_PublishClientStatsServer) error {
	return status.Errorf(codes.Unimplemented, "method PublishClientStats not implemented")
}
func (UnimplementedClientServiceServer) SubscribeToCommands(*SubscribeToCommandsRequest, ClientService_SubscribeToCommandsServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToCommands not implemented")
}
func (UnimplementedClientServiceServer) AckCommand(context.Context, *AckCommandRequest) (*AckCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AckCommand not implemented")
}
func (UnimplementedClientServiceServer) mustEmbedUnimplementedClientServiceServer() {}

// UnsafeClientServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientServiceServer will
// result in compilation errors.
type UnsafeClientServiceServer interface {
	mustEmbedUnimplementedClientServiceServer()
}

func RegisterClientServiceServer(s grpc.ServiceRegistrar, srv ClientServiceServer) {
	s.RegisterService(&ClientService_ServiceDesc, srv)
}

func _ClientService_RegisterClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).RegisterClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientService_RegisterClient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).RegisterClient(ctx, req.(*RegisterClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_PublishClientStats_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClientServiceServer).PublishClientStats(&clientServicePublishClientStatsServer{stream})
}

type ClientService_PublishClientStatsServer interface {
	SendAndClose(*PublishClientStatsResponse) error
	Recv() (*PublishClientStatsRequest, error)
	grpc.ServerStream
}

type clientServicePublishClientStatsServer struct {
	grpc.ServerStream
}

func (x *clientServicePublishClientStatsServer) SendAndClose(m *PublishClientStatsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clientServicePublishClientStatsServer) Recv() (*PublishClientStatsRequest, error) {
	m := new(PublishClientStatsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ClientService_SubscribeToCommands_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeToCommandsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ClientServiceServer).SubscribeToCommands(m, &clientServiceSubscribeToCommandsServer{stream})
}

type ClientService_SubscribeToCommandsServer interface {
	Send(*SubscribeToCommandsResponse) error
	grpc.ServerStream
}

type clientServiceSubscribeToCommandsServer struct {
	grpc.ServerStream
}

func (x *clientServiceSubscribeToCommandsServer) Send(m *SubscribeToCommandsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ClientService_AckCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AckCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).AckCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientService_AckCommand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).AckCommand(ctx, req.(*AckCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientService_ServiceDesc is the grpc.ServiceDesc for ClientService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "queuer.clients.v1.ClientService",
	HandlerType: (*ClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterClient",
			Handler:    _ClientService_RegisterClient_Handler,
		},
		{
			MethodName: "AckCommand",
			Handler:    _ClientService_AckCommand_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PublishClientStats",
			Handler:       _ClientService_PublishClientStats_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SubscribeToCommands",
			Handler:       _ClientService_SubscribeToCommands_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "queuer/clients/v1/service.proto",
}