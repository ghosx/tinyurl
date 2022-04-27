package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "github.com/ghosx/tinyurl/gen/go/proto/counter"
)

type server struct {
	pb.UnimplementedCounterServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) GetCount(ctx context.Context, in *pb.CounterRequest) (*pb.CounterResponse, error) {
	start, end := in.Current + 1, in.Current + int64(in.Count)
	log.Printf("generate count:%d,%d", start, end)
	return &pb.CounterResponse{
		Start: start,
		End:   end,
	}, nil
}


func main(){
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	pb.RegisterCounterServer(s, &server{})
	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	// go func() {
	// 	log.Fatalln(s.Serve(lis))
	// }()
	log.Fatalln(s.Serve(lis))

}

