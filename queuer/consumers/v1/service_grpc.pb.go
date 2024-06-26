// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: queuer/consumers/v1/service.proto

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
	PresenceService_IAmAlive_FullMethodName = "/queuer.consumers.v1.PresenceService/IAmAlive"
)

// PresenceServiceClient is the client API for PresenceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PresenceServiceClient interface {
	IAmAlive(ctx context.Context, in *IAmAliveRequest, opts ...grpc.CallOption) (*IAmAliveResponse, error)
}

type presenceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPresenceServiceClient(cc grpc.ClientConnInterface) PresenceServiceClient {
	return &presenceServiceClient{cc}
}

func (c *presenceServiceClient) IAmAlive(ctx context.Context, in *IAmAliveRequest, opts ...grpc.CallOption) (*IAmAliveResponse, error) {
	out := new(IAmAliveResponse)
	err := c.cc.Invoke(ctx, PresenceService_IAmAlive_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PresenceServiceServer is the server API for PresenceService service.
// All implementations must embed UnimplementedPresenceServiceServer
// for forward compatibility
type PresenceServiceServer interface {
	IAmAlive(context.Context, *IAmAliveRequest) (*IAmAliveResponse, error)
	mustEmbedUnimplementedPresenceServiceServer()
}

// UnimplementedPresenceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPresenceServiceServer struct {
}

func (UnimplementedPresenceServiceServer) IAmAlive(context.Context, *IAmAliveRequest) (*IAmAliveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IAmAlive not implemented")
}
func (UnimplementedPresenceServiceServer) mustEmbedUnimplementedPresenceServiceServer() {}

// UnsafePresenceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PresenceServiceServer will
// result in compilation errors.
type UnsafePresenceServiceServer interface {
	mustEmbedUnimplementedPresenceServiceServer()
}

func RegisterPresenceServiceServer(s grpc.ServiceRegistrar, srv PresenceServiceServer) {
	s.RegisterService(&PresenceService_ServiceDesc, srv)
}

func _PresenceService_IAmAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IAmAliveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresenceServiceServer).IAmAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PresenceService_IAmAlive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresenceServiceServer).IAmAlive(ctx, req.(*IAmAliveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PresenceService_ServiceDesc is the grpc.ServiceDesc for PresenceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PresenceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "queuer.consumers.v1.PresenceService",
	HandlerType: (*PresenceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IAmAlive",
			Handler:    _PresenceService_IAmAlive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "queuer/consumers/v1/service.proto",
}

const (
	ConsumerService_RegisterConsumer_FullMethodName            = "/queuer.consumers.v1.ConsumerService/RegisterConsumer"
	ConsumerService_GetConsumer_FullMethodName                 = "/queuer.consumers.v1.ConsumerService/GetConsumer"
	ConsumerService_GetAuthorizedConsumers_FullMethodName      = "/queuer.consumers.v1.ConsumerService/GetAuthorizedConsumers"
	ConsumerService_GetUnauthorizedConsumers_FullMethodName    = "/queuer.consumers.v1.ConsumerService/GetUnauthorizedConsumers"
	ConsumerService_AuthorizeConsumer_FullMethodName           = "/queuer.consumers.v1.ConsumerService/AuthorizeConsumer"
	ConsumerService_DeauthorizeConsumer_FullMethodName         = "/queuer.consumers.v1.ConsumerService/DeauthorizeConsumer"
	ConsumerService_RetrieveConsumerCredentials_FullMethodName = "/queuer.consumers.v1.ConsumerService/RetrieveConsumerCredentials"
	ConsumerService_PublishConsumerStats_FullMethodName        = "/queuer.consumers.v1.ConsumerService/PublishConsumerStats"
	ConsumerService_GetConsumerStreams_FullMethodName          = "/queuer.consumers.v1.ConsumerService/GetConsumerStreams"
	ConsumerService_GetSubscribedStreams_FullMethodName        = "/queuer.consumers.v1.ConsumerService/GetSubscribedStreams"
)

// ConsumerServiceClient is the client API for ConsumerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConsumerServiceClient interface {
	RegisterConsumer(ctx context.Context, in *RegisterConsumerRequest, opts ...grpc.CallOption) (*RegisterConsumerResponse, error)
	GetConsumer(ctx context.Context, in *GetConsumerRequest, opts ...grpc.CallOption) (*GetConsumerResponse, error)
	GetAuthorizedConsumers(ctx context.Context, in *GetAuthorizedConsumersRequest, opts ...grpc.CallOption) (*GetAuthorizedConsumersResponse, error)
	GetUnauthorizedConsumers(ctx context.Context, in *GetUnauthorizedConsumersRequest, opts ...grpc.CallOption) (*GetUnauthorizedConsumersResponse, error)
	AuthorizeConsumer(ctx context.Context, in *AuthorizeConsumerRequest, opts ...grpc.CallOption) (*AuthorizeConsumerResponse, error)
	DeauthorizeConsumer(ctx context.Context, in *DeauthorizeConsumerRequest, opts ...grpc.CallOption) (*DeauthorizeConsumerResponse, error)
	RetrieveConsumerCredentials(ctx context.Context, in *RetrieveConsumerCredentialsRequest, opts ...grpc.CallOption) (*RetrieveConsumerCredentialsResponse, error)
	PublishConsumerStats(ctx context.Context, opts ...grpc.CallOption) (ConsumerService_PublishConsumerStatsClient, error)
	GetConsumerStreams(ctx context.Context, in *GetConsumerStreamsRequest, opts ...grpc.CallOption) (*GetConsumerStreamsResponse, error)
	GetSubscribedStreams(ctx context.Context, in *GetSubscribedStreamsRequest, opts ...grpc.CallOption) (*GetSubscribedStreamsResponse, error)
}

type consumerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConsumerServiceClient(cc grpc.ClientConnInterface) ConsumerServiceClient {
	return &consumerServiceClient{cc}
}

func (c *consumerServiceClient) RegisterConsumer(ctx context.Context, in *RegisterConsumerRequest, opts ...grpc.CallOption) (*RegisterConsumerResponse, error) {
	out := new(RegisterConsumerResponse)
	err := c.cc.Invoke(ctx, ConsumerService_RegisterConsumer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerServiceClient) GetConsumer(ctx context.Context, in *GetConsumerRequest, opts ...grpc.CallOption) (*GetConsumerResponse, error) {
	out := new(GetConsumerResponse)
	err := c.cc.Invoke(ctx, ConsumerService_GetConsumer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerServiceClient) GetAuthorizedConsumers(ctx context.Context, in *GetAuthorizedConsumersRequest, opts ...grpc.CallOption) (*GetAuthorizedConsumersResponse, error) {
	out := new(GetAuthorizedConsumersResponse)
	err := c.cc.Invoke(ctx, ConsumerService_GetAuthorizedConsumers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerServiceClient) GetUnauthorizedConsumers(ctx context.Context, in *GetUnauthorizedConsumersRequest, opts ...grpc.CallOption) (*GetUnauthorizedConsumersResponse, error) {
	out := new(GetUnauthorizedConsumersResponse)
	err := c.cc.Invoke(ctx, ConsumerService_GetUnauthorizedConsumers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerServiceClient) AuthorizeConsumer(ctx context.Context, in *AuthorizeConsumerRequest, opts ...grpc.CallOption) (*AuthorizeConsumerResponse, error) {
	out := new(AuthorizeConsumerResponse)
	err := c.cc.Invoke(ctx, ConsumerService_AuthorizeConsumer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerServiceClient) DeauthorizeConsumer(ctx context.Context, in *DeauthorizeConsumerRequest, opts ...grpc.CallOption) (*DeauthorizeConsumerResponse, error) {
	out := new(DeauthorizeConsumerResponse)
	err := c.cc.Invoke(ctx, ConsumerService_DeauthorizeConsumer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerServiceClient) RetrieveConsumerCredentials(ctx context.Context, in *RetrieveConsumerCredentialsRequest, opts ...grpc.CallOption) (*RetrieveConsumerCredentialsResponse, error) {
	out := new(RetrieveConsumerCredentialsResponse)
	err := c.cc.Invoke(ctx, ConsumerService_RetrieveConsumerCredentials_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerServiceClient) PublishConsumerStats(ctx context.Context, opts ...grpc.CallOption) (ConsumerService_PublishConsumerStatsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ConsumerService_ServiceDesc.Streams[0], ConsumerService_PublishConsumerStats_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &consumerServicePublishConsumerStatsClient{stream}
	return x, nil
}

type ConsumerService_PublishConsumerStatsClient interface {
	Send(*PublishConsumerStatsRequest) error
	CloseAndRecv() (*PublishConsumerStatsResponse, error)
	grpc.ClientStream
}

type consumerServicePublishConsumerStatsClient struct {
	grpc.ClientStream
}

func (x *consumerServicePublishConsumerStatsClient) Send(m *PublishConsumerStatsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *consumerServicePublishConsumerStatsClient) CloseAndRecv() (*PublishConsumerStatsResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PublishConsumerStatsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *consumerServiceClient) GetConsumerStreams(ctx context.Context, in *GetConsumerStreamsRequest, opts ...grpc.CallOption) (*GetConsumerStreamsResponse, error) {
	out := new(GetConsumerStreamsResponse)
	err := c.cc.Invoke(ctx, ConsumerService_GetConsumerStreams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerServiceClient) GetSubscribedStreams(ctx context.Context, in *GetSubscribedStreamsRequest, opts ...grpc.CallOption) (*GetSubscribedStreamsResponse, error) {
	out := new(GetSubscribedStreamsResponse)
	err := c.cc.Invoke(ctx, ConsumerService_GetSubscribedStreams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConsumerServiceServer is the server API for ConsumerService service.
// All implementations must embed UnimplementedConsumerServiceServer
// for forward compatibility
type ConsumerServiceServer interface {
	RegisterConsumer(context.Context, *RegisterConsumerRequest) (*RegisterConsumerResponse, error)
	GetConsumer(context.Context, *GetConsumerRequest) (*GetConsumerResponse, error)
	GetAuthorizedConsumers(context.Context, *GetAuthorizedConsumersRequest) (*GetAuthorizedConsumersResponse, error)
	GetUnauthorizedConsumers(context.Context, *GetUnauthorizedConsumersRequest) (*GetUnauthorizedConsumersResponse, error)
	AuthorizeConsumer(context.Context, *AuthorizeConsumerRequest) (*AuthorizeConsumerResponse, error)
	DeauthorizeConsumer(context.Context, *DeauthorizeConsumerRequest) (*DeauthorizeConsumerResponse, error)
	RetrieveConsumerCredentials(context.Context, *RetrieveConsumerCredentialsRequest) (*RetrieveConsumerCredentialsResponse, error)
	PublishConsumerStats(ConsumerService_PublishConsumerStatsServer) error
	GetConsumerStreams(context.Context, *GetConsumerStreamsRequest) (*GetConsumerStreamsResponse, error)
	GetSubscribedStreams(context.Context, *GetSubscribedStreamsRequest) (*GetSubscribedStreamsResponse, error)
	mustEmbedUnimplementedConsumerServiceServer()
}

// UnimplementedConsumerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConsumerServiceServer struct {
}

func (UnimplementedConsumerServiceServer) RegisterConsumer(context.Context, *RegisterConsumerRequest) (*RegisterConsumerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterConsumer not implemented")
}
func (UnimplementedConsumerServiceServer) GetConsumer(context.Context, *GetConsumerRequest) (*GetConsumerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConsumer not implemented")
}
func (UnimplementedConsumerServiceServer) GetAuthorizedConsumers(context.Context, *GetAuthorizedConsumersRequest) (*GetAuthorizedConsumersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthorizedConsumers not implemented")
}
func (UnimplementedConsumerServiceServer) GetUnauthorizedConsumers(context.Context, *GetUnauthorizedConsumersRequest) (*GetUnauthorizedConsumersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUnauthorizedConsumers not implemented")
}
func (UnimplementedConsumerServiceServer) AuthorizeConsumer(context.Context, *AuthorizeConsumerRequest) (*AuthorizeConsumerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorizeConsumer not implemented")
}
func (UnimplementedConsumerServiceServer) DeauthorizeConsumer(context.Context, *DeauthorizeConsumerRequest) (*DeauthorizeConsumerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeauthorizeConsumer not implemented")
}
func (UnimplementedConsumerServiceServer) RetrieveConsumerCredentials(context.Context, *RetrieveConsumerCredentialsRequest) (*RetrieveConsumerCredentialsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveConsumerCredentials not implemented")
}
func (UnimplementedConsumerServiceServer) PublishConsumerStats(ConsumerService_PublishConsumerStatsServer) error {
	return status.Errorf(codes.Unimplemented, "method PublishConsumerStats not implemented")
}
func (UnimplementedConsumerServiceServer) GetConsumerStreams(context.Context, *GetConsumerStreamsRequest) (*GetConsumerStreamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConsumerStreams not implemented")
}
func (UnimplementedConsumerServiceServer) GetSubscribedStreams(context.Context, *GetSubscribedStreamsRequest) (*GetSubscribedStreamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscribedStreams not implemented")
}
func (UnimplementedConsumerServiceServer) mustEmbedUnimplementedConsumerServiceServer() {}

// UnsafeConsumerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConsumerServiceServer will
// result in compilation errors.
type UnsafeConsumerServiceServer interface {
	mustEmbedUnimplementedConsumerServiceServer()
}

func RegisterConsumerServiceServer(s grpc.ServiceRegistrar, srv ConsumerServiceServer) {
	s.RegisterService(&ConsumerService_ServiceDesc, srv)
}

func _ConsumerService_RegisterConsumer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterConsumerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).RegisterConsumer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConsumerService_RegisterConsumer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).RegisterConsumer(ctx, req.(*RegisterConsumerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsumerService_GetConsumer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConsumerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).GetConsumer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConsumerService_GetConsumer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).GetConsumer(ctx, req.(*GetConsumerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsumerService_GetAuthorizedConsumers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthorizedConsumersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).GetAuthorizedConsumers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConsumerService_GetAuthorizedConsumers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).GetAuthorizedConsumers(ctx, req.(*GetAuthorizedConsumersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsumerService_GetUnauthorizedConsumers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUnauthorizedConsumersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).GetUnauthorizedConsumers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConsumerService_GetUnauthorizedConsumers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).GetUnauthorizedConsumers(ctx, req.(*GetUnauthorizedConsumersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsumerService_AuthorizeConsumer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeConsumerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).AuthorizeConsumer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConsumerService_AuthorizeConsumer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).AuthorizeConsumer(ctx, req.(*AuthorizeConsumerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsumerService_DeauthorizeConsumer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeauthorizeConsumerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).DeauthorizeConsumer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConsumerService_DeauthorizeConsumer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).DeauthorizeConsumer(ctx, req.(*DeauthorizeConsumerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsumerService_RetrieveConsumerCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveConsumerCredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).RetrieveConsumerCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConsumerService_RetrieveConsumerCredentials_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).RetrieveConsumerCredentials(ctx, req.(*RetrieveConsumerCredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsumerService_PublishConsumerStats_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ConsumerServiceServer).PublishConsumerStats(&consumerServicePublishConsumerStatsServer{stream})
}

type ConsumerService_PublishConsumerStatsServer interface {
	SendAndClose(*PublishConsumerStatsResponse) error
	Recv() (*PublishConsumerStatsRequest, error)
	grpc.ServerStream
}

type consumerServicePublishConsumerStatsServer struct {
	grpc.ServerStream
}

func (x *consumerServicePublishConsumerStatsServer) SendAndClose(m *PublishConsumerStatsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *consumerServicePublishConsumerStatsServer) Recv() (*PublishConsumerStatsRequest, error) {
	m := new(PublishConsumerStatsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ConsumerService_GetConsumerStreams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConsumerStreamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).GetConsumerStreams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConsumerService_GetConsumerStreams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).GetConsumerStreams(ctx, req.(*GetConsumerStreamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsumerService_GetSubscribedStreams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubscribedStreamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).GetSubscribedStreams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConsumerService_GetSubscribedStreams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).GetSubscribedStreams(ctx, req.(*GetSubscribedStreamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConsumerService_ServiceDesc is the grpc.ServiceDesc for ConsumerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConsumerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "queuer.consumers.v1.ConsumerService",
	HandlerType: (*ConsumerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterConsumer",
			Handler:    _ConsumerService_RegisterConsumer_Handler,
		},
		{
			MethodName: "GetConsumer",
			Handler:    _ConsumerService_GetConsumer_Handler,
		},
		{
			MethodName: "GetAuthorizedConsumers",
			Handler:    _ConsumerService_GetAuthorizedConsumers_Handler,
		},
		{
			MethodName: "GetUnauthorizedConsumers",
			Handler:    _ConsumerService_GetUnauthorizedConsumers_Handler,
		},
		{
			MethodName: "AuthorizeConsumer",
			Handler:    _ConsumerService_AuthorizeConsumer_Handler,
		},
		{
			MethodName: "DeauthorizeConsumer",
			Handler:    _ConsumerService_DeauthorizeConsumer_Handler,
		},
		{
			MethodName: "RetrieveConsumerCredentials",
			Handler:    _ConsumerService_RetrieveConsumerCredentials_Handler,
		},
		{
			MethodName: "GetConsumerStreams",
			Handler:    _ConsumerService_GetConsumerStreams_Handler,
		},
		{
			MethodName: "GetSubscribedStreams",
			Handler:    _ConsumerService_GetSubscribedStreams_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PublishConsumerStats",
			Handler:       _ConsumerService_PublishConsumerStats_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "queuer/consumers/v1/service.proto",
}
