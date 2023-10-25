package usecase

import (
	"context"

	pb "github.com/DmitriiKumancev/gRPC-nmap/internal/delivery/grpc/netvuln"
)

type NetVulnService interface {
	CheckVuln(context.Context, *pb.CheckVulnRequest) (*pb.CheckVulnResponse, error)
}
