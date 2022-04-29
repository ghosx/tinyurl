package main

import (
	"context"
	"log"
	"net"
	"sync/atomic"

	pb "github.com/ghosx/tinyurl/gen/go/proto/counter"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCounterServer
}

func NewServer() *server {
	return &server{}
}

var globalCount uint64 = 0

func (s *server) GetCount(ctx context.Context, in *pb.CounterRequest) (*pb.CounterResponse, error) {
	start, end := globalCount, globalCount+in.Count
	atomic.AddUint64(&globalCount, in.Count)
	log.Printf("generate count:%d,%d", start, end)
	return &pb.CounterResponse{
		Start: start,
		End:   end,
	}, nil
}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	pb.RegisterCounterServer(s, &server{})
	// Serve gRPC server
	log.Println("Serving gRPC on :9000")
	// go func() {
	// 	log.Fatalln(s.Serve(lis))
	// }()
	log.Fatalln(s.Serve(lis))

}
