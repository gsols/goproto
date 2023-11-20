// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: queuer/queues/v1/queue.proto

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
	QueueService_CreateQueue_FullMethodName    = "/queuer.queues.v1.QueueService/CreateQueue"
	QueueService_GetQueue_FullMethodName       = "/queuer.queues.v1.QueueService/GetQueue"
	QueueService_DeleteQueue_FullMethodName    = "/queuer.queues.v1.QueueService/DeleteQueue"
	QueueService_FlushQueue_FullMethodName     = "/queuer.queues.v1.QueueService/FlushQueue"
	QueueService_PublishMessage_FullMethodName = "/queuer.queues.v1.QueueService/PublishMessage"
	QueueService_ClientStats_FullMethodName    = "/queuer.queues.v1.QueueService/ClientStats"
)

// QueueServiceClient is the client API for QueueService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueueServiceClient interface {
	CreateQueue(ctx context.Context, in *CreateQueueRequest, opts ...grpc.CallOption) (*CreateQueueResponse, error)
	GetQueue(ctx context.Context, in *GetQueueRequest, opts ...grpc.CallOption) (*GetQueueResponse, error)
	DeleteQueue(ctx context.Context, in *DeleteQueueRequest, opts ...grpc.CallOption) (*DeleteQueueResponse, error)
	FlushQueue(ctx context.Context, in *FlushQueueRequest, opts ...grpc.CallOption) (*FlushQueueResponse, error)
	PublishMessage(ctx context.Context, in *PublishMessageRequest, opts ...grpc.CallOption) (*PublishMessageResponse, error)
	ClientStats(ctx context.Context, opts ...grpc.CallOption) (QueueService_ClientStatsClient, error)
}

type queueServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQueueServiceClient(cc grpc.ClientConnInterface) QueueServiceClient {
	return &queueServiceClient{cc}
}

func (c *queueServiceClient) CreateQueue(ctx context.Context, in *CreateQueueRequest, opts ...grpc.CallOption) (*CreateQueueResponse, error) {
	out := new(CreateQueueResponse)
	err := c.cc.Invoke(ctx, QueueService_CreateQueue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueServiceClient) GetQueue(ctx context.Context, in *GetQueueRequest, opts ...grpc.CallOption) (*GetQueueResponse, error) {
	out := new(GetQueueResponse)
	err := c.cc.Invoke(ctx, QueueService_GetQueue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueServiceClient) DeleteQueue(ctx context.Context, in *DeleteQueueRequest, opts ...grpc.CallOption) (*DeleteQueueResponse, error) {
	out := new(DeleteQueueResponse)
	err := c.cc.Invoke(ctx, QueueService_DeleteQueue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueServiceClient) FlushQueue(ctx context.Context, in *FlushQueueRequest, opts ...grpc.CallOption) (*FlushQueueResponse, error) {
	out := new(FlushQueueResponse)
	err := c.cc.Invoke(ctx, QueueService_FlushQueue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueServiceClient) PublishMessage(ctx context.Context, in *PublishMessageRequest, opts ...grpc.CallOption) (*PublishMessageResponse, error) {
	out := new(PublishMessageResponse)
	err := c.cc.Invoke(ctx, QueueService_PublishMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueServiceClient) ClientStats(ctx context.Context, opts ...grpc.CallOption) (QueueService_ClientStatsClient, error) {
	stream, err := c.cc.NewStream(ctx, &QueueService_ServiceDesc.Streams[0], QueueService_ClientStats_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &queueServiceClientStatsClient{stream}
	return x, nil
}

type QueueService_ClientStatsClient interface {
	Send(*ClientStatsRequest) error
	CloseAndRecv() (*ClientStatsResponse, error)
	grpc.ClientStream
}

type queueServiceClientStatsClient struct {
	grpc.ClientStream
}

func (x *queueServiceClientStatsClient) Send(m *ClientStatsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *queueServiceClientStatsClient) CloseAndRecv() (*ClientStatsResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ClientStatsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// QueueServiceServer is the server API for QueueService service.
// All implementations must embed UnimplementedQueueServiceServer
// for forward compatibility
type QueueServiceServer interface {
	CreateQueue(context.Context, *CreateQueueRequest) (*CreateQueueResponse, error)
	GetQueue(context.Context, *GetQueueRequest) (*GetQueueResponse, error)
	DeleteQueue(context.Context, *DeleteQueueRequest) (*DeleteQueueResponse, error)
	FlushQueue(context.Context, *FlushQueueRequest) (*FlushQueueResponse, error)
	PublishMessage(context.Context, *PublishMessageRequest) (*PublishMessageResponse, error)
	ClientStats(QueueService_ClientStatsServer) error
	mustEmbedUnimplementedQueueServiceServer()
}

// UnimplementedQueueServiceServer must be embedded to have forward compatible implementations.
type UnimplementedQueueServiceServer struct {
}

func (UnimplementedQueueServiceServer) CreateQueue(context.Context, *CreateQueueRequest) (*CreateQueueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQueue not implemented")
}
func (UnimplementedQueueServiceServer) GetQueue(context.Context, *GetQueueRequest) (*GetQueueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQueue not implemented")
}
func (UnimplementedQueueServiceServer) DeleteQueue(context.Context, *DeleteQueueRequest) (*DeleteQueueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQueue not implemented")
}
func (UnimplementedQueueServiceServer) FlushQueue(context.Context, *FlushQueueRequest) (*FlushQueueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FlushQueue not implemented")
}
func (UnimplementedQueueServiceServer) PublishMessage(context.Context, *PublishMessageRequest) (*PublishMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishMessage not implemented")
}
func (UnimplementedQueueServiceServer) ClientStats(QueueService_ClientStatsServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientStats not implemented")
}
func (UnimplementedQueueServiceServer) mustEmbedUnimplementedQueueServiceServer() {}

// UnsafeQueueServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueueServiceServer will
// result in compilation errors.
type UnsafeQueueServiceServer interface {
	mustEmbedUnimplementedQueueServiceServer()
}

func RegisterQueueServiceServer(s grpc.ServiceRegistrar, srv QueueServiceServer) {
	s.RegisterService(&QueueService_ServiceDesc, srv)
}

func _QueueService_CreateQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServiceServer).CreateQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueueService_CreateQueue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServiceServer).CreateQueue(ctx, req.(*CreateQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueService_GetQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServiceServer).GetQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueueService_GetQueue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServiceServer).GetQueue(ctx, req.(*GetQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueService_DeleteQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServiceServer).DeleteQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueueService_DeleteQueue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServiceServer).DeleteQueue(ctx, req.(*DeleteQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueService_FlushQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlushQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServiceServer).FlushQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueueService_FlushQueue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServiceServer).FlushQueue(ctx, req.(*FlushQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueService_PublishMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServiceServer).PublishMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueueService_PublishMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServiceServer).PublishMessage(ctx, req.(*PublishMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueService_ClientStats_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(QueueServiceServer).ClientStats(&queueServiceClientStatsServer{stream})
}

type QueueService_ClientStatsServer interface {
	SendAndClose(*ClientStatsResponse) error
	Recv() (*ClientStatsRequest, error)
	grpc.ServerStream
}

type queueServiceClientStatsServer struct {
	grpc.ServerStream
}

func (x *queueServiceClientStatsServer) SendAndClose(m *ClientStatsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *queueServiceClientStatsServer) Recv() (*ClientStatsRequest, error) {
	m := new(ClientStatsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// QueueService_ServiceDesc is the grpc.ServiceDesc for QueueService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QueueService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "queuer.queues.v1.QueueService",
	HandlerType: (*QueueServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateQueue",
			Handler:    _QueueService_CreateQueue_Handler,
		},
		{
			MethodName: "GetQueue",
			Handler:    _QueueService_GetQueue_Handler,
		},
		{
			MethodName: "DeleteQueue",
			Handler:    _QueueService_DeleteQueue_Handler,
		},
		{
			MethodName: "FlushQueue",
			Handler:    _QueueService_FlushQueue_Handler,
		},
		{
			MethodName: "PublishMessage",
			Handler:    _QueueService_PublishMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ClientStats",
			Handler:       _QueueService_ClientStats_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "queuer/queues/v1/queue.proto",
}
