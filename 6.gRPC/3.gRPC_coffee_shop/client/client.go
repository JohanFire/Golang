package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "3.gRPC_coffee_shop/coffee_shop_protos"
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

	client := pb.NewCoffeeShopServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	menuStream, err := client.GetMenu(ctx, &pb.MenuRequest{})
	if err != nil {
		log.Fatalf("Failed to call getMenu. \n %v", err)
	}

	done := make(chan bool)

	var items []*pb.Item
	go func() {
		for {
			response, err := menuStream.Recv()
			if err == io.EOF {
				done <- true
				return
			}

			if err != nil {
				log.Fatalf("Error while receiving stream. \n %v", err)
			}

			items = response.Items
			log.Printf("Received item: %v", items)
		}
	}()

	<-done // wait for the stream to finish

	fmt.Println()

	receipt, err := client.PlaceOrder(ctx, &pb.Order{
		Items: items,
	})
	if err != nil {
		log.Fatalf("Failed to call PlaceOrder. \n %v", err)
	}

	log.Printf("Receipt: %v", receipt)

	fmt.Println()

	status, err := client.GetOrderStatus(ctx, receipt)
	if err != nil {
		log.Fatalf("Failed to call GetOrderStatus. \n %v", err)
	}

	log.Printf("Order status: %v", status)
}
