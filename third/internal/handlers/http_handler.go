package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	pb "github.com/KasiditR/backend-challenge-third/api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type BeefResponse struct {
	Beef map[string]int `json:"beef"`
}

const grpcServerAddr = "localhost:50051"

func GetBeefHandler(w http.ResponseWriter, _ *http.Request) {
	creds := insecure.NewCredentials()
	conn, err := grpc.NewClient(grpcServerAddr, grpc.WithTransportCredentials(creds))
	if err != nil {
		http.Error(w, "Failed to connect to gRPC server", http.StatusInternalServerError)
		log.Fatalf("Failed to create client: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	client := pb.NewBeefClient(conn)
	res, err := client.CountBeef(ctx, &pb.Empty{})
	if err != nil {
		http.Error(w, "Error calling gRPC service", http.StatusInternalServerError)
		log.Printf("Error calling gRPC: %v", err)
	}

	beefMap := res.Beef.AsMap()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(beefMap)
}
