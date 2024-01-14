package main

import (
	"fmt"
	pb "github.com/dmitry-shylov/test.grpc.dev/ecommerce"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"io"
	"log"
	"net"
	"sync"
)

const (
	orderServerPort = ":50053"
)

type OrderManager struct {
	mu       sync.Mutex
	orderMap map[string]pb.Order
	pb.UnimplementedOrderManagementServer
}

func (o *OrderManager) UpdateOrders(updateOrderServer pb.OrderManagement_UpdateOrdersServer) error {
	o.mu.Lock()
	defer o.mu.Unlock()
	ordersStr := "Updated Order IDs : "
	for {
		order, err := updateOrderServer.Recv()
		if err == io.EOF {
			result := wrapperspb.StringValue{Value: fmt.Sprintf("Orders processed %s", ordersStr)}

			return updateOrderServer.SendAndClose(&result)
		}

		if o.orderMap == nil {
			o.orderMap = make(map[string]pb.Order)
		}
		o.orderMap[order.Id] = *order
		log.Println("Order ID ", order.Id, ": Updated")
		ordersStr += order.Id + ", "
	}
}

func main() {
	lis, err := net.Listen("tcp", orderServerPort)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	productInfoServer := OrderManager{}
	pb.RegisterOrderManagementServer(s, &productInfoServer)
	log.Printf("Starting gRPC listener on port " + orderServerPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
