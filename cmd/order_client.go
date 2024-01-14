package main

import (
	"context"
	pb "github.com/dmitry-shylov/test.grpc.dev/ecommerce"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	orderClientAddress = "localhost:50053"
)

func main() {
	conn, err := grpc.Dial(orderClientAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client := pb.NewOrderManagementClient(conn)
	updateStream, err := client.UpdateOrders(ctx)

	if err != nil {
		log.Fatalf("%v.UpdateOrders(_) = _, %v", client, err)
	}

	updOrder1 := pb.Order{Id: "1", Description: "Помідори"}
	if err := updateStream.Send(&updOrder1); err != nil {
		log.Fatalf("%v.Send(%v) = %v", updateStream, &updOrder1, err)
	}

	updOrder2 := pb.Order{Id: "2", Description: "Борошно"}
	if err := updateStream.Send(&updOrder2); err != nil {
		log.Fatalf("%v.Send(%v) = %v",
			updateStream, &updOrder2, err)
	}

	updOrder3 := pb.Order{Id: "3", Description: "Хліб"}
	if err := updateStream.Send(&updOrder3); err != nil {
		log.Fatalf("%v.Send(%v) = %v",
			updateStream, &updOrder3, err)
	}

	updateRes, err := updateStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v",
			updateStream, err, nil)
	}
	log.Printf("Update Orders Res : %s", updateRes)
}
