package main

import (
	"fmt"
	"log"
	"net"

	pb "1.gRPC/classification"
	"google.golang.org/grpc"
	// Import the generated classification package
	// "1.gRPC/classification"
)

func main() {
	var port uint16 = 9090
	var address string = fmt.Sprintf(":%d", port)
	// address := fmt.Sprintf(":%d", port)

	listen, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("Failed to listen on port %d: %v", port, err)
	} else {
		log.Printf("Listening on port %d", port)
	}

	// classificationServer := grpc.NewServer()
	// classificationService := &classification.ClassificationService{}

	grpcServer := grpc.NewServer()

	pb.RegisterClassificationServiceServer(grpcServer, &pb.Server{})

	// Register the classification server with the gRPC server
	// classification.RegisterClassificationServiceServer(grpcServer, &classificationServer)
	// classification.RegisterClassificationServiceServer(grpcServer, &classificationServer{})
	// classification.RegisterClassificationServiceServer(grpcServer, classificationService)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve gRPC server on port %d: %v", port, err)
	} else {
		log.Printf("gRPC server running on port %d", port)
	}

}
