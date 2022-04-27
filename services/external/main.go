package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/ghosx/tinyurl/gen/go/proto/external"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedExternalServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) GetCount(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	return &pb.CreateResponse{
		Time: in.Time.AsTime().Format(time.RFC1123),
	}, nil
}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	pb.RegisterExternalServer(s, &server{})
	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	// go func() {
	// 	log.Fatalln(s.Serve(lis))
	// }()
	log.Fatalln(s.Serve(lis))

}
