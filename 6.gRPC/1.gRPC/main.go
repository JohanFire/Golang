package main

import (
	"fmt"
	"log"
	"net"

	pb "1.gRPC/classification"
	"google.golang.org/grpc"
)

func main() {
	var port uint16 = 9090
	var address string = fmt.Sprintf(":%d", port)

	listen, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("Failed to listen on port %d: %v", port, err)
	} else {
		log.Printf("Listening on port %d", port)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterClassificationServiceServer(grpcServer, &pb.Server{})

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve gRPC server on port %d: %v", port, err)
	} else {
		log.Printf("gRPC server running on port %d", port)
	}

}
