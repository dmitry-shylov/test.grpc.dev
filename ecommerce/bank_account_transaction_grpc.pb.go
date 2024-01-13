// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: bank_account_transaction.proto

package __

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
	BankAccountTransactionInfo_GetBankAccountTransactions_FullMethodName = "/ecommerce.BankAccountTransactionInfo/getBankAccountTransactions"
)

// BankAccountTransactionInfoClient is the client API for BankAccountTransactionInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BankAccountTransactionInfoClient interface {
	GetBankAccountTransactions(ctx context.Context, in *BankAccountId, opts ...grpc.CallOption) (BankAccountTransactionInfo_GetBankAccountTransactionsClient, error)
}

type bankAccountTransactionInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewBankAccountTransactionInfoClient(cc grpc.ClientConnInterface) BankAccountTransactionInfoClient {
	return &bankAccountTransactionInfoClient{cc}
}

func (c *bankAccountTransactionInfoClient) GetBankAccountTransactions(ctx context.Context, in *BankAccountId, opts ...grpc.CallOption) (BankAccountTransactionInfo_GetBankAccountTransactionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &BankAccountTransactionInfo_ServiceDesc.Streams[0], BankAccountTransactionInfo_GetBankAccountTransactions_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &bankAccountTransactionInfoGetBankAccountTransactionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BankAccountTransactionInfo_GetBankAccountTransactionsClient interface {
	Recv() (*BankAccountTransaction, error)
	grpc.ClientStream
}

type bankAccountTransactionInfoGetBankAccountTransactionsClient struct {
	grpc.ClientStream
}

func (x *bankAccountTransactionInfoGetBankAccountTransactionsClient) Recv() (*BankAccountTransaction, error) {
	m := new(BankAccountTransaction)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BankAccountTransactionInfoServer is the server API for BankAccountTransactionInfo service.
// All implementations must embed UnimplementedBankAccountTransactionInfoServer
// for forward compatibility
type BankAccountTransactionInfoServer interface {
	GetBankAccountTransactions(*BankAccountId, BankAccountTransactionInfo_GetBankAccountTransactionsServer) error
	mustEmbedUnimplementedBankAccountTransactionInfoServer()
}

// UnimplementedBankAccountTransactionInfoServer must be embedded to have forward compatible implementations.
type UnimplementedBankAccountTransactionInfoServer struct {
}

func (UnimplementedBankAccountTransactionInfoServer) GetBankAccountTransactions(*BankAccountId, BankAccountTransactionInfo_GetBankAccountTransactionsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetBankAccountTransactions not implemented")
}
func (UnimplementedBankAccountTransactionInfoServer) mustEmbedUnimplementedBankAccountTransactionInfoServer() {
}

// UnsafeBankAccountTransactionInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BankAccountTransactionInfoServer will
// result in compilation errors.
type UnsafeBankAccountTransactionInfoServer interface {
	mustEmbedUnimplementedBankAccountTransactionInfoServer()
}

func RegisterBankAccountTransactionInfoServer(s grpc.ServiceRegistrar, srv BankAccountTransactionInfoServer) {
	s.RegisterService(&BankAccountTransactionInfo_ServiceDesc, srv)
}

func _BankAccountTransactionInfo_GetBankAccountTransactions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BankAccountId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BankAccountTransactionInfoServer).GetBankAccountTransactions(m, &bankAccountTransactionInfoGetBankAccountTransactionsServer{stream})
}

type BankAccountTransactionInfo_GetBankAccountTransactionsServer interface {
	Send(*BankAccountTransaction) error
	grpc.ServerStream
}

type bankAccountTransactionInfoGetBankAccountTransactionsServer struct {
	grpc.ServerStream
}

func (x *bankAccountTransactionInfoGetBankAccountTransactionsServer) Send(m *BankAccountTransaction) error {
	return x.ServerStream.SendMsg(m)
}

// BankAccountTransactionInfo_ServiceDesc is the grpc.ServiceDesc for BankAccountTransactionInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BankAccountTransactionInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ecommerce.BankAccountTransactionInfo",
	HandlerType: (*BankAccountTransactionInfoServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "getBankAccountTransactions",
			Handler:       _BankAccountTransactionInfo_GetBankAccountTransactions_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "bank_account_transaction.proto",
}