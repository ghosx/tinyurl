package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"

	counterpb "github.com/ghosx/tinyurl/gen/go/proto/counter"
	pb "github.com/ghosx/tinyurl/gen/go/proto/external"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	pb.UnimplementedExternalServer
}

type Counter struct {
	Cur uint64
	Max uint64
}

var Client counterpb.CounterClient
var counter = Counter{Cur: 0, Max: 0}
var CounterRange uint64 = 10

func getCount(in *counterpb.CounterRequest) (*counterpb.CounterResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := Client.GetCount(ctx, in)
	if err != nil {
		log.Fatalf("could not get count: %v", err)
	}
	log.Printf("get count from counter service:%d,%d", r.GetStart(), r.GetEnd())
	return r, nil
}

func init() {
	addr := "0.0.0.0:8080"
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//defer conn.Close()
	Client = counterpb.NewCounterClient(conn)
	r, _ := getCount(&counterpb.CounterRequest{Current: 0, Count: CounterRange})
	counter = Counter{Cur: r.GetStart(), Max: r.GetEnd()}
}

func NewServer() *server {
	return &server{}
}

func (s *server) CreateUrl(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	if counter.Cur >= counter.Max {
		r, _ := getCount(&counterpb.CounterRequest{Current: counter.Cur, Count: CounterRange})
		counter = Counter{Cur: r.GetStart(), Max: r.GetEnd()}
	}
	counter.Cur++
	return &pb.CreateResponse{
		ShortUrl: "http://tinyurl.com/" + in.Url + "/" + strconv.Itoa(int(counter.Cur)),
	}, nil
}
// test

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	pb.RegisterExternalServer(s, &server{})
	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:9090")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()
	

	conn, err := grpc.DialContext(
				context.Background(),
				"0.0.0.0:9090",
				grpc.WithBlock(),
				grpc.WithInsecure(),
			)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	gwmux := runtime.NewServeMux()
// 	// Register Greeter
	err = pb.RegisterExternalHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":9091",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:9091")
	log.Fatalln(gwServer.ListenAndServe())


}
