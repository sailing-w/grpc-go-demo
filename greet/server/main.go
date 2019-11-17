package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-go-demo/greet/greetpb"
	"log"
	"net"
)
const (
	port = ":50051"
)
// server is used to implement greetpb.GreeterServer.
type server struct{}

func (s *server) SayHello(ctx context.Context, in *greetpb.HelloRequest) (*greetpb.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &greetpb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port) // Specify the port we want to use to listen for client requests
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {   // Call Serve() on the server with our port details to do a blocking wait until the process is killed or Stop() is called.
		log.Fatalf("failed to serve: %v", err)
	}
}
