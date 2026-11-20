package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "example.com/ordertracking/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type orderServer struct {
	pb.UnimplementedOrderServiceServer
}

func (s *orderServer) GetOrder(ctx context.Context, req *pb.OrderRequest) (*pb.OrderReply, error) {
	if req.GetOrderId() == "" {
		return nil, status.Error(codes.InvalidArgument, "order ID cannot be empty")
	}

	log.Printf("Unary request received for order %s", req.GetOrderId())
	return &pb.OrderReply{
		OrderId: req.GetOrderId(),
		Product: "Laptop",
		Status:  "Order received",
	}, nil
}

func (s *orderServer) TrackOrder(req *pb.OrderRequest, stream grpc.ServerStreamingServer[pb.OrderUpdate]) error {
	if req.GetOrderId() == "" {
		return status.Error(codes.InvalidArgument, "order ID cannot be empty")
	}

	log.Printf("Streaming started for order %s", req.GetOrderId())
	updates := []string{
		"Order confirmed",
		"Preparing the package",
		"Package shipped",
		"Out for delivery",
		"Delivered",
	}

	for _, step := range updates {
		select {
		case <-stream.Context().Done():
			log.Println("Client stopped tracking")
			return stream.Context().Err()
		case <-time.After(2 * time.Second):
		}

		if err := stream.Send(&pb.OrderUpdate{
			OrderId: req.GetOrderId(),
			Status:  step,
		}); err != nil {
			return err
		}
		log.Printf("Sent update: %s", step)
	}

	log.Printf("Streaming finished for order %s", req.GetOrderId())
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, &orderServer{})
	log.Println("Order server is running on port 50051")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
