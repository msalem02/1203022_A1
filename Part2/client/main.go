
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "example.com/ordertracking/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	connection, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	client := pb.NewOrderServiceClient(connection)
	request := &pb.OrderRequest{OrderId: "1001"}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	info, err := client.GetOrder(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("UNARY RPC RESULT")
	fmt.Println("Order ID:", info.GetOrderId())
	fmt.Println("Product:", info.GetProduct())
	fmt.Println("Status:", info.GetStatus())

	fmt.Println("\nSTREAMING RPC UPDATES")
	stream, err := client.TrackOrder(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}

	for {
		update, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("Stream finished")
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Received:", update.GetStatus())
	}
}
