package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "1.gRPC/classification"
	// pb "classification"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:9090", grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server. \n %v", err)
	}

	defer conn.Close()

	client := pb.NewClassificationServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	fmt.Println()

	request, err := client.Classify(ctx, &pb.ClassificationRequest{
		Name:         "Johan Tristán",
		Appliance:    "Washing Machine",
		ApplianceVIB: "DishWasher123456",
	})
	if err != nil {
		log.Fatal("Error while calling Classify: ", err)
	}

	log.Printf("Response from Classify: %s", request.Response)

}
