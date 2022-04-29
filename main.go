package main

import (
	"fmt"
	"strconv"
	"time"
	"math/rand"
)

// test

func main() {
	fmt.Println(time.Now().Local())
	var b int64 = 8
	s := fmt.Sprintf("%b", b)
	fmt.Println("b=", s)
	fmt.Println(strconv.FormatInt(b, 2))
	l := make([]int, 10)
	for i := 0; i < 10; i++ {
		l[i] = i
	}
	fmt.Println(l)
	rand.Shuffle(len(l), func(i, j int) {
		l[i], l[j] = l[j], l[i]
	})
	fmt.Println(l)
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
