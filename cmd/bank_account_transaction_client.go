package main

import (
	"context"
	pb "github.com/dmitry-shylov/test.grpc.dev/ecommerce"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

const (
	bankAccountTransactionAddress = "localhost:50052"
)

func main() {
	conn, err := grpc.Dial(bankAccountTransactionAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewBankAccountTransactionInfoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetBankAccountTransactions(ctx,
		&pb.BankAccountId{Value: "11111"})
	if err != nil {
		log.Fatalf("Could not add product: %v", err)
	}

	for {
		searchOrder, err := r.Recv()
		if err == io.EOF {
			break
		}

		log.Print("Search Result : ", searchOrder)
	}
}
