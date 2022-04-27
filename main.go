package main

import (
	"context"
	"fmt"
	"time"

	counterpb "github.com/ghosx/tinyurl/gen/go/proto/counter"
)

type counterServer struct {
	counterpb.UnimplementedCounterServer
}

func NewServer() *counterServer {
	return &counterServer{}
}

func (s *counterServer) GetCount(ctx context.Context, in *counterpb.CounterRequest) (*counterpb.CounterResponse, error) {
	return &counterpb.CounterResponse{
		Start: in.Current + 1,
		End:   in.Current + 1 + int64(in.Count),
	}, nil
}

func main() {
	fmt.Println(time.Now().Local())
	a := 1
	fmt.Println(a)
	fmt.Println(time.Now().Location())
}

// 	// Create a client connection to the gRPC server we just started
// 	// This is where the gRPC-Gateway proxies the requests
// 	conn, err := grpc.DialContext(
// 		context.Background(),
// 		"0.0.0.0:8080",
// 		grpc.WithBlock(),
// 		grpc.WithInsecure(),
// 	)
// 	if err != nil {
// 		log.Fatalln("Failed to dial server:", err)
// 	}

// 	gwmux := runtime.NewServeMux()
// 	// Register Greeter
// 	err = externalpb.RegisterExternalHandler(context.Background(), gwmux, conn)
// 	if err != nil {
// 		log.Fatalln("Failed to register gateway:", err)
// 	}

// 	gwServer := &http.Server{
// 		Addr:    ":8090",
// 		Handler: gwmux,
// 	}

// 	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
// 	log.Fatalln(gwServer.ListenAndServe())
// }
