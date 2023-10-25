package usecase

import (
	"context"
	"strconv"

	pb "github.com/DmitriiKumancev/gRPC-nmap/internal/delivery/grpc/netvuln"
	"github.com/DmitriiKumancev/gRPC-nmap/pkg/logger"
	"github.com/DmitriiKumancev/gRPC-nmap/pkg/nmap"
	vulvmap "github.com/Ullaakut/nmap/v3"
)

type netVulnSerice struct {
	nmap *nmap.Nmap
	log  *logger.Logger
}

func NewNetVulnGrpcService(nmap *nmap.Nmap, log *logger.Logger) NetVulnService {
	return &netVulnSerice{
		nmap: nmap,
		log:  log,
	}
}

func (n *netVulnSerice) CheckVuln(ctx context.Context, req *pb.CheckVulnRequest) (*pb.CheckVulnResponse, error) {

	result, err := n.nmap.Scanner(ctx, req)
	if err != nil {
		n.log.Error("error during nmap scan: %v", err)
		return nil, err
	}

	response := &pb.CheckVulnResponse{
		Results: make([]*pb.TargetResult, 0, len(result.Hosts)),
	}

	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		targetResult := &pb.TargetResult{
			Target:   host.Addresses[0].String(),
			Services: make([]*pb.Service, 0),
		}

		for _, port := range host.Ports {

			service := &pb.Service{
				Name:    port.Service.Name,
				Version: port.Service.Version,
				TcpPort: int32(port.ID),
				Vulns:   make([]*pb.Vulnerability, 0),
			}

			vulns := n.createCvssAndId(&port)
			service.Vulns = append(service.Vulns, vulns...)

			targetResult.Services = append(targetResult.Services, service)
		}
		response.Results = append(response.Results, targetResult)
	}

	return response, nil
}

func (n *netVulnSerice) createCvssAndId(port *vulvmap.Port) []*pb.Vulnerability {
	vulns := make([]*pb.Vulnerability, 0)

	for _, script := range port.Scripts {

		if script.ID == "vulners" {

			for _, table := range script.Tables[0].Tables {

				vuln := &pb.Vulnerability{}
				for _, el := range table.Elements {

					switch el.Key {

					case "id":
						vuln.Identifier = el.Value
					case "cvss":
						cvssFloat, err := strconv.ParseFloat(el.Value, 32)
						if err != nil {
							n.log.Error("failed to convert string to float64: %v", err)
						}
						vuln.CvssScore = float32(cvssFloat)
					}
				}
				vulns = append(vulns, vuln)
			}
		}
	}

	n.log.Info("Created vulnerabilities for port: %d", port.ID)
	return vulns
}
