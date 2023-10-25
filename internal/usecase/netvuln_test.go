package usecase

import (
	"context"
	"log"
	"testing"
	"time"

	pb "github.com/DmitriiKumancev/gRPC-nmap/internal/delivery/grpc/netvuln"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestCheckVuln(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewNetVulnServiceClient(conn)

	t.Run("testasp.vulnweb.com", func(t *testing.T) {

		reqTestV2 := &pb.CheckVulnRequest{
			Targets: []string{"testasp.vulnweb.com"},
			TcpPort: []int32{80},
		}

		got, err := c.CheckVuln(ctx, reqTestV2)
		if err != nil {
			t.Fatalf("CheckVuln returned an error: %v", err)
		}

		if got == nil {
			t.Fatal("Received response is nil")
		}

		if len(got.Results) == 0 {
			t.Fatal("Received response does not contain any results")
		}

		for _, result := range got.Results {
			if result.Target != "44.238.29.244" {
				t.Errorf("Expected target: '44.238.29.244', but got: %s", result.Target)
			}

			for _, service := range result.Services {
				if service.Name != "http" {
					t.Errorf("Expected service name: 'http', but got: %s", service.Name)
				}
				if service.Version != "8.5" {
					t.Errorf("Expected service version: '8.5', but got: %s", service.Version)
				}
				if service.TcpPort != 80 {
					t.Errorf("Expected service TCP port: 80, but got: %d", service.TcpPort)
				}
			}
		}
	})

	t.Run("scanme.nmap.org", func(t *testing.T) {
		reqTestV1 := &pb.CheckVulnRequest{
			Targets: []string{"scanme.nmap.org"},
			TcpPort: []int32{80},
		}

		got, err := c.CheckVuln(ctx, reqTestV1)
		if err != nil {
			t.Fatalf("CheckVuln returned an error: %v", err)
		}

		if got == nil {
			t.Fatal("Received response is nil")
		}

		if len(got.Results) == 0 {
			t.Fatal("Received response does not contain any results")
		}

		for _, result := range got.Results {
			if result.Target != "45.33.32.156" {
				t.Errorf("Expected target: '45.33.32.156', but got: %s", result.Target)
			}

			for _, service := range result.Services {
				if service.Name != "http" {
					t.Errorf("Expected service name: 'http', but got: %s", service.Name)
				}
				if service.Version != "2.4.7" {
					t.Errorf("Expected service version: '2.4.7', but got: %s", service.Version)
				}
				if service.TcpPort != 80 {
					t.Errorf("Expected service TCP port: 80, but got: %d", service.TcpPort)
				}
			}
		}
	})
}
