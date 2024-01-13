package main

import (
	"fmt"
	pb "github.com/dmitry-shylov/test.grpc.dev/ecommerce"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	bankAccountTransactionPort = ":50052"
)

type BankAccountServer struct {
	pb.UnimplementedBankAccountTransactionInfoServer
}

func (b *BankAccountServer) GetBankAccountTransactions(bankAccountID *pb.BankAccountId, server pb.BankAccountTransactionInfo_GetBankAccountTransactionsServer) error {
	bankTransaction := pb.BankAccountTransaction{Name: "PrivatBank", Id: bankAccountID.Value, Amount: 3000}
	bankTransaction2 := pb.BankAccountTransaction{Name: "OschadBank", Id: bankAccountID.Value, Amount: 2000}

	err := server.Send(&bankTransaction)
	if err != nil {
		return fmt.Errorf(
			"error sending message to stream : %v", err)
	}

	log.Print("Matching Bank Account Found : " + bankAccountID.Value)

	err = server.Send(&bankTransaction2)
	if err != nil {
		return fmt.Errorf(
			"error sending message to stream : %v", err)
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", bankAccountTransactionPort)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	bankAccountServer := BankAccountServer{}
	pb.RegisterBankAccountTransactionInfoServer(s, &bankAccountServer)
	log.Printf("Starting gRPC listener on port " + bankAccountTransactionPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
