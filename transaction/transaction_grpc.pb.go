// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: transaction.proto

package transaction

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
	TransactionManager_PerformOperation_FullMethodName = "/TransactionManager/PerformOperation"
)

// TransactionManagerClient is the client API for TransactionManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionManagerClient interface {
	PerformOperation(ctx context.Context, in *OperationRequest, opts ...grpc.CallOption) (*OperationResponse, error)
}

type transactionManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionManagerClient(cc grpc.ClientConnInterface) TransactionManagerClient {
	return &transactionManagerClient{cc}
}

func (c *transactionManagerClient) PerformOperation(ctx context.Context, in *OperationRequest, opts ...grpc.CallOption) (*OperationResponse, error) {
	out := new(OperationResponse)
	err := c.cc.Invoke(ctx, TransactionManager_PerformOperation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionManagerServer is the server API for TransactionManager service.
// All implementations must embed UnimplementedTransactionManagerServer
// for forward compatibility
type TransactionManagerServer interface {
	PerformOperation(context.Context, *OperationRequest) (*OperationResponse, error)
	mustEmbedUnimplementedTransactionManagerServer()
}

// UnimplementedTransactionManagerServer must be embedded to have forward compatible implementations.
type UnimplementedTransactionManagerServer struct {
}

func (UnimplementedTransactionManagerServer) PerformOperation(context.Context, *OperationRequest) (*OperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PerformOperation not implemented")
}
func (UnimplementedTransactionManagerServer) mustEmbedUnimplementedTransactionManagerServer() {}

// UnsafeTransactionManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionManagerServer will
// result in compilation errors.
type UnsafeTransactionManagerServer interface {
	mustEmbedUnimplementedTransactionManagerServer()
}

func RegisterTransactionManagerServer(s grpc.ServiceRegistrar, srv TransactionManagerServer) {
	s.RegisterService(&TransactionManager_ServiceDesc, srv)
}

func _TransactionManager_PerformOperation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionManagerServer).PerformOperation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionManager_PerformOperation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionManagerServer).PerformOperation(ctx, req.(*OperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransactionManager_ServiceDesc is the grpc.ServiceDesc for TransactionManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransactionManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TransactionManager",
	HandlerType: (*TransactionManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PerformOperation",
			Handler:    _TransactionManager_PerformOperation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transaction.proto",
}
