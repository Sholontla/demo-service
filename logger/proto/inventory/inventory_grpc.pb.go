// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: inventory.proto

package inventory

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

// InventoryServiceClient is the client API for InventoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InventoryServiceClient interface {
	Inventory(ctx context.Context, in *InvenoryRequest, opts ...grpc.CallOption) (*InventoryResponse, error)
	StreamInventory(ctx context.Context, opts ...grpc.CallOption) (InventoryService_StreamInventoryClient, error)
}

type inventoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInventoryServiceClient(cc grpc.ClientConnInterface) InventoryServiceClient {
	return &inventoryServiceClient{cc}
}

func (c *inventoryServiceClient) Inventory(ctx context.Context, in *InvenoryRequest, opts ...grpc.CallOption) (*InventoryResponse, error) {
	out := new(InventoryResponse)
	err := c.cc.Invoke(ctx, "/inventory.InventoryService/Inventory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) StreamInventory(ctx context.Context, opts ...grpc.CallOption) (InventoryService_StreamInventoryClient, error) {
	stream, err := c.cc.NewStream(ctx, &InventoryService_ServiceDesc.Streams[0], "/inventory.InventoryService/StreamInventory", opts...)
	if err != nil {
		return nil, err
	}
	x := &inventoryServiceStreamInventoryClient{stream}
	return x, nil
}

type InventoryService_StreamInventoryClient interface {
	Send(*InvenoryRequest) error
	Recv() (*InventoryResponse, error)
	grpc.ClientStream
}

type inventoryServiceStreamInventoryClient struct {
	grpc.ClientStream
}

func (x *inventoryServiceStreamInventoryClient) Send(m *InvenoryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *inventoryServiceStreamInventoryClient) Recv() (*InventoryResponse, error) {
	m := new(InventoryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// InventoryServiceServer is the server API for InventoryService service.
// All implementations must embed UnimplementedInventoryServiceServer
// for forward compatibility
type InventoryServiceServer interface {
	Inventory(context.Context, *InvenoryRequest) (*InventoryResponse, error)
	StreamInventory(InventoryService_StreamInventoryServer) error
	mustEmbedUnimplementedInventoryServiceServer()
}

// UnimplementedInventoryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInventoryServiceServer struct {
}

func (UnimplementedInventoryServiceServer) Inventory(context.Context, *InvenoryRequest) (*InventoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Inventory not implemented")
}
func (UnimplementedInventoryServiceServer) StreamInventory(InventoryService_StreamInventoryServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamInventory not implemented")
}
func (UnimplementedInventoryServiceServer) mustEmbedUnimplementedInventoryServiceServer() {}

// UnsafeInventoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InventoryServiceServer will
// result in compilation errors.
type UnsafeInventoryServiceServer interface {
	mustEmbedUnimplementedInventoryServiceServer()
}

func RegisterInventoryServiceServer(s grpc.ServiceRegistrar, srv InventoryServiceServer) {
	s.RegisterService(&InventoryService_ServiceDesc, srv)
}

func _InventoryService_Inventory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvenoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).Inventory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.InventoryService/Inventory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).Inventory(ctx, req.(*InvenoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_StreamInventory_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(InventoryServiceServer).StreamInventory(&inventoryServiceStreamInventoryServer{stream})
}

type InventoryService_StreamInventoryServer interface {
	Send(*InventoryResponse) error
	Recv() (*InvenoryRequest, error)
	grpc.ServerStream
}

type inventoryServiceStreamInventoryServer struct {
	grpc.ServerStream
}

func (x *inventoryServiceStreamInventoryServer) Send(m *InventoryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *inventoryServiceStreamInventoryServer) Recv() (*InvenoryRequest, error) {
	m := new(InvenoryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// InventoryService_ServiceDesc is the grpc.ServiceDesc for InventoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InventoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inventory.InventoryService",
	HandlerType: (*InventoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Inventory",
			Handler:    _InventoryService_Inventory_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamInventory",
			Handler:       _InventoryService_StreamInventory_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "inventory.proto",
}
