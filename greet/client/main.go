package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-go-demo/greet/greetpb"
	"log"
	"os"
	"time"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)
func main() {
	// Set up a connection to the server.

	// To call service methods, we first need to create a gRPC channel to communicate with the server. We create this by passing the
	// server address and port number to grpc.Dial()
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := greetpb.NewGreeterClient(conn)
	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &greetpb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
