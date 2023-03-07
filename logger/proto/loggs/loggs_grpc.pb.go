// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: loggs.proto

package loggs

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

// LoggsProducerServiceClient is the client API for LoggsProducerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoggsProducerServiceClient interface {
	StreamProducer(ctx context.Context, opts ...grpc.CallOption) (LoggsProducerService_StreamProducerClient, error)
}

type loggsProducerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoggsProducerServiceClient(cc grpc.ClientConnInterface) LoggsProducerServiceClient {
	return &loggsProducerServiceClient{cc}
}

func (c *loggsProducerServiceClient) StreamProducer(ctx context.Context, opts ...grpc.CallOption) (LoggsProducerService_StreamProducerClient, error) {
	stream, err := c.cc.NewStream(ctx, &LoggsProducerService_ServiceDesc.Streams[0], "/producer.LoggsProducerService/StreamProducer", opts...)
	if err != nil {
		return nil, err
	}
	x := &loggsProducerServiceStreamProducerClient{stream}
	return x, nil
}

type LoggsProducerService_StreamProducerClient interface {
	Send(*MainLoggerRequest) error
	Recv() (*MainLoggerResponse, error)
	grpc.ClientStream
}

type loggsProducerServiceStreamProducerClient struct {
	grpc.ClientStream
}

func (x *loggsProducerServiceStreamProducerClient) Send(m *MainLoggerRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *loggsProducerServiceStreamProducerClient) Recv() (*MainLoggerResponse, error) {
	m := new(MainLoggerResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LoggsProducerServiceServer is the server API for LoggsProducerService service.
// All implementations must embed UnimplementedLoggsProducerServiceServer
// for forward compatibility
type LoggsProducerServiceServer interface {
	StreamProducer(LoggsProducerService_StreamProducerServer) error
	mustEmbedUnimplementedLoggsProducerServiceServer()
}

// UnimplementedLoggsProducerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLoggsProducerServiceServer struct {
}

func (UnimplementedLoggsProducerServiceServer) StreamProducer(LoggsProducerService_StreamProducerServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamProducer not implemented")
}
func (UnimplementedLoggsProducerServiceServer) mustEmbedUnimplementedLoggsProducerServiceServer() {}

// UnsafeLoggsProducerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoggsProducerServiceServer will
// result in compilation errors.
type UnsafeLoggsProducerServiceServer interface {
	mustEmbedUnimplementedLoggsProducerServiceServer()
}

func RegisterLoggsProducerServiceServer(s grpc.ServiceRegistrar, srv LoggsProducerServiceServer) {
	s.RegisterService(&LoggsProducerService_ServiceDesc, srv)
}

func _LoggsProducerService_StreamProducer_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LoggsProducerServiceServer).StreamProducer(&loggsProducerServiceStreamProducerServer{stream})
}

type LoggsProducerService_StreamProducerServer interface {
	Send(*MainLoggerResponse) error
	Recv() (*MainLoggerRequest, error)
	grpc.ServerStream
}

type loggsProducerServiceStreamProducerServer struct {
	grpc.ServerStream
}

func (x *loggsProducerServiceStreamProducerServer) Send(m *MainLoggerResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *loggsProducerServiceStreamProducerServer) Recv() (*MainLoggerRequest, error) {
	m := new(MainLoggerRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LoggsProducerService_ServiceDesc is the grpc.ServiceDesc for LoggsProducerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LoggsProducerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "producer.LoggsProducerService",
	HandlerType: (*LoggsProducerServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamProducer",
			Handler:       _LoggsProducerService_StreamProducer_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "loggs.proto",
}
