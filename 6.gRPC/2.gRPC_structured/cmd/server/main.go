package main

import (
	"fmt"
	"log"
	"net"

	"go-router/internal/handler"
	pb "go-router/internal/protobuf"
	"go-router/internal/service"

	"google.golang.org/grpc"
)

func main() {
	var port uint16 = 9090
	var address string = fmt.Sprintf(":%d", port)

	listen, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("Failed to listen on port %d: %v", port, err)
	} else {
		log.Printf("Listening on port: %d", port)
	}

	println()

	// Initialize services and handlers
	classificationService := service.NewClassificationService()
	classificationHandler := handler.NewClassificationHandler(classificationService)

	grpcServer := grpc.NewServer()
	pb.RegisterClassificationServiceServer(grpcServer, classificationHandler)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve gRPC server on port %d: %v", port, err)
	}
}
