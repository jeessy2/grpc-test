// Package main implements a server for Greeter service.
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "grpc-test/proto"
)

const (
	port = ":50051"
)

type server struct {
	pb.SimpleServiceServer
}

func (s *server) Route(ctx context.Context, in *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	log.Printf("Received: %v", in.GetData())
	return &pb.SimpleResponse{Code: 0, Value: "From server"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSimpleServiceServer(s, &server{})
	log.Println("starting server with" + port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
