package main

import (
	"context"
	"log"
	"time"

	pb "go-router/internal/protobuf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:9090", grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}

	defer conn.Close()

	client := pb.NewClassificationServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	request := &pb.ClassificationRequest{
		UserName:      "Johan Tristán",
		ApplianceType: "Washing Machine",
		ApplianceVIB:  "DishWasher123456",
	}

	println()

	response, err := client.Classify(ctx, request)
	if err != nil {
		log.Fatalf("Error while calling Classify: %v", err)
	}

	log.Printf("Response from Classify: \n %s", response.Response)
	println()
}
