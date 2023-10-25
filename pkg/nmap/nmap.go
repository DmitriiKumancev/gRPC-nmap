package nmap

import (
	"context"
	"strconv"

	pb "github.com/DmitriiKumancev/gRPC-nmap/internal/delivery/grpc/netvuln"
	"github.com/DmitriiKumancev/gRPC-nmap/pkg/logger"
	"github.com/Ullaakut/nmap/v3"
)

type Nmap struct {
	log *logger.Logger
}

func NewNmapScanner(log *logger.Logger) *Nmap {
	return &Nmap{
		log: log,
	}
}

func (n *Nmap) Scanner(ctx context.Context, req *pb.CheckVulnRequest) (*nmap.Run, error) {
	n.log.Info("Start nmap scanner")
	scanner, err := nmap.NewScanner(
		ctx,
		nmap.WithTargets(req.Targets...),
		nmap.WithPorts(intToString(req.TcpPort)...),
		nmap.WithServiceInfo(),
		nmap.WithScripts("vulners"),
	)
	if err != nil {
		n.log.Fatal("unable to create nmap scanner: %v", err)
	}

	result, warnings, err := scanner.Run()
	if len(*warnings) > 0 {
		n.log.Error("run finished with warnings: %s\n", *warnings)
	}
	if err != nil {
		n.log.Error("unable to run nmap scan: %v", err)
	}

	n.log.Info("The scanner has finished working")
	return result, err
}

func intToString(arr []int32) []string {
	var arrStr []string

	for _, el := range arr {
		strEl := strconv.Itoa(int(el))
		arrStr = append(arrStr, strEl)
	}

	return arrStr
}
