package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	pb "example.com/distributedpow/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type workerServer struct {
	pb.UnimplementedWorkerServiceServer
	id string
}

func (w *workerServer) SearchRange(req *pb.SearchRequest, stream grpc.ServerStreamingServer[pb.SearchUpdate]) error {
	log.Printf("Worker %s searching job %s range [%d, %d)", w.id, req.JobId, req.Start, req.End)
	prefix := strings.Repeat("0", int(req.Difficulty))
	var batch uint64
	for nonce := req.Start; nonce < req.End; nonce++ {
		if nonce%1024 == 0 {
			if err := stream.Context().Err(); err != nil {
				log.Printf("Worker %s STOPPED job %s due to cancellation", w.id, req.JobId)
				return err
			}
		}
		h := sha256.Sum256([]byte(req.Data + strconv.FormatUint(nonce, 10)))
		hash := hex.EncodeToString(h[:])
		batch++
		if strings.HasPrefix(hash, prefix) {
			log.Printf("Worker %s FOUND nonce %d for job %s", w.id, nonce, req.JobId)
			return stream.Send(&pb.SearchUpdate{JobId: req.JobId, Checked: batch, Finished: true, Found: true, Nonce: nonce, Hash: hash})
		}
		if batch >= 10000 {
			if err := stream.Send(&pb.SearchUpdate{JobId: req.JobId, Checked: batch}); err != nil {
				log.Printf("Worker %s STOPPED job %s: %v", w.id, req.JobId, err)
				return err
			}
			batch = 0
		}
	}
	log.Printf("Worker %s finished range [%d, %d) without a result", w.id, req.Start, req.End)
	return stream.Send(&pb.SearchUpdate{JobId: req.JobId, Checked: batch, Finished: true})
}

func main() {
	id := flag.String("id", "worker1", "unique worker ID")
	port := flag.Int("port", 6001, "worker gRPC port")
	host := flag.String("host", "127.0.0.1", "address the coordinator can reach")
	coordinator := flag.String("coordinator", "127.0.0.1:50051", "coordinator gRPC address")
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	pb.RegisterWorkerServiceServer(server, &workerServer{id: *id})
	go func() {
		log.Printf("Worker %s listening on port %d", *id, *port)
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	for {
		conn, err := grpc.NewClient(*coordinator, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Printf("Dial problem: %v", err)
			time.Sleep(2 * time.Second)
			continue
		}
		client := pb.NewCoordinatorServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		reply, err := client.RegisterWorker(ctx, &pb.RegisterRequest{WorkerId: *id, Address: fmt.Sprintf("%s:%d", *host, *port)})
		cancel()
		if err != nil || !reply.Accepted {
			log.Printf("Register failed: %v", err)
			conn.Close()
			time.Sleep(2 * time.Second)
			continue
		}
		log.Printf("Worker %s registered with coordinator", *id)
		ticker := time.NewTicker(2 * time.Second)
		connected := true
		for connected {
			<-ticker.C
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			_, err := client.Heartbeat(ctx, &pb.HeartbeatRequest{WorkerId: *id})
			cancel()
			if err != nil {
				log.Printf("Heartbeat failed. Reconnecting: %v", err)
				connected = false
			}
		}
		ticker.Stop()
		conn.Close()
	}
}
