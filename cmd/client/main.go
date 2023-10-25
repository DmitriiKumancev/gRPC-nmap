package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/DmitriiKumancev/gRPC-nmap/internal/config"
	pb "github.com/DmitriiKumancev/gRPC-nmap/internal/delivery/grpc/netvuln"
	"github.com/DmitriiKumancev/gRPC-nmap/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	configPath := "configs/config.json"

	log := logger.New()
	cfg, err := config.New(configPath)
	if err != nil {
		log.Error("failed to load config: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)

	defer cancel()

	conn, err := grpc.Dial(cfg.ClientGrpc.Port, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewNetVulnServiceClient(conn)

	request := &pb.CheckVulnRequest{
		Targets: []string{"scanme.nmap.org"},
		TcpPort: []int32{53, 80},
	}

	r, err := c.CheckVuln(ctx, request)
	if err != nil {
		log.Error("failed to get response from server: %v", err)
	}

	result := r.GetResults()

	resultByte, err := json.MarshalIndent(&result, " ", " ")
	if err != nil {
		log.Error("unable to marshaling response struct: %v", err)
	}

	log.Info("Results: %s", resultByte)

}
