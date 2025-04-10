package main

import (
	"context"
	"fmt"
	"log"
	"net"

	// "log"
	// "net"

	pb "3.gRPC_coffee_shop/coffee_shop_protos"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedCoffeeShopServiceServer
}

// func (s *Server) GetMenu(MenuRequest *pb.MenuRequest)

func (s *Server) GetMenu(menuRequest *pb.MenuRequest, server grpc.ServerStreamingServer[pb.MenuResponse]) error {
	// return status.Errorf(codes.Unimplemented, "method GetMenu not implemented")

	items := []*pb.Item{
		&pb.Item{
			Id:   "1",
			Name: "Black Coffee",
		},
		&pb.Item{
			Id:   "2",
			Name: "Americano",
		},
		&pb.Item{
			Id:   "3",
			Name: "Latte",
		},
	}

	for i, _ := range items {
		server.Send(&pb.MenuResponse{
			// Items: items[i],
			Items: items[0 : i+1],
		})
	}

	return nil
}
func (s *Server) PlaceOrder(context context.Context, order *pb.Order) (*pb.Receipt, error) {
	// return nil, status.Errorf(codes.Unimplemented, "method PlaceOrder not implemented")

	return &pb.Receipt{
		Id: "ABC123",
	}, nil
}
func (s *Server) GetOrderStatus(context context.Context, receipt *pb.Receipt) (*pb.OrderStatus, error) {
	// return nil, status.Errorf(codes.Unimplemented, "method GetOrderStatus not implemented")

	return &pb.OrderStatus{
		OrderId: receipt.Id,
		Status:  "Your order is being prepared",
	}, nil
}

func main() {
	// setup listener on port 9001
	var port uint16 = 9090
	var address string = fmt.Sprintf(":%d", port)

	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on port %d: %v", port, err)
		panic(err)
	} else {
		log.Printf("Listening on port %d", port)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCoffeeShopServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve gRPC server on port %d: %v", port, err)
		panic(err)
	} else {
		log.Printf("gRPC server running on port %d", port)
	}
}
