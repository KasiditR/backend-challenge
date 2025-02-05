package main

import (
	"log"
	"net"

	pb "github.com/KasiditR/backend-challenge-third/api/proto"
	"github.com/KasiditR/backend-challenge-third/internal/services"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBeefServer(grpcServer, &services.BeefServiceServer{})

	log.Printf("GRPC server is running on port %v", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start GRPC server: %v", err)
	}
}
