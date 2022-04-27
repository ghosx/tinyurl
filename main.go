package main

import (
	"fmt"
	"time"
)


func main() {
	fmt.Println(time.Now().Local())
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
