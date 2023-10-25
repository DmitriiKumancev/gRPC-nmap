package grpc

import (
	"context"

	pb "github.com/DmitriiKumancev/gRPC-nmap/internal/delivery/grpc/netvuln"
	"github.com/DmitriiKumancev/gRPC-nmap/internal/usecase"
	"github.com/DmitriiKumancev/gRPC-nmap/pkg/logger"
)

type vulnServer struct {
	usecase usecase.NetVulnService
	log     *logger.Logger
	pb.UnimplementedNetVulnServiceServer
}

func NewVulnGrpcServerHandler(usecase usecase.NetVulnService, log *logger.Logger) pb.NetVulnServiceServer {
	return &vulnServer{
		usecase: usecase,
		log:     log,
	}
}

func (v *vulnServer) CheckVuln(ctx context.Context, req *pb.CheckVulnRequest) (*pb.CheckVulnResponse, error) {

	response, err := v.usecase.CheckVuln(ctx, req)
	if err != nil {
		v.log.Error("failed with method in usecase: %v", err)
	}

	return response, nil
}
