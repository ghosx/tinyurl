package main

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
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
	L   []uint64
}

func NewCounter(cur, max uint64) *Counter {
	l := make([]uint64, max-cur)
	for i := 0; i < int(max-cur); i++ {
		l[i] = cur + uint64(i)
	}
	rand.Shuffle(len(l), func(i, j int) {
		l[i], l[j] = l[j], l[i]
	})
	fmt.Println("l", l)
	return &Counter{Cur: cur, Max: max, L: l}
}

func (c *Counter) getCounter() uint64 {
	r := c.L[0]
	c.L = c.L[1:]
	c.Cur++
	return r
}

var Client counterpb.CounterClient
var counter *Counter
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
	counter = NewCounter(r.GetStart(), r.GetEnd())
}

func NewServer() *server {
	return &server{}
}

func (s *server) CreateUrl(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	if counter.Cur >= counter.Max {
		r, _ := getCount(&counterpb.CounterRequest{Current: counter.Cur, Count: CounterRange})
		counter = NewCounter(r.GetStart(), r.GetEnd())
	}
	r := counter.getCounter()
	fmt.Println("r", r)
	rr := strconv.FormatUint(r, 2)
	m := md5.Sum([]byte(rr))
	fmt.Println("m", m)
	fmt.Printf("%x\n", m)
	u := base64.URLEncoding.EncodeToString(m[:])
	u = u[:8]
	fmt.Println("create url:", u)
	return &pb.CreateResponse{
		ShortUrl: "http://tinyurl.com/" + u,
	}, nil
}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	pb.RegisterExternalServer(s, &server{})
	// Serve gRPC server
	log.Println("Serving gRPC on :8000")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8000",
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
		Addr:    ":80",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:80")
	log.Fatalln(gwServer.ListenAndServe())
}
