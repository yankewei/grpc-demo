package main

import (
	"context"
	"log"
	"net"
	welcomeservice "grpc-go/proto/welcomeservice"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *welcomeservice.HelloRequest) (*welcomeservice.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &welcomeservice.HelloReply{Message: "hello " + in.Name}, nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	welcomeservice.RegisterWelcomeServiceServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
