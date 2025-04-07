package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	var port uint16 = 9090
	var address string = fmt.Sprintf(":%d", port)
	// address := fmt.Sprintf(":%d", port)

	listen, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("Failed to listen on port %d: %v", port, err)
	}

	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve gRPC server on port %d: %v", port, err)
	} else {
		log.Printf("gRPC server running on port %d", port)
	}

}
