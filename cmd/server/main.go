package main

import (
	"github.com/DmitriiKumancev/gRPC-nmap/internal/config"
	"github.com/DmitriiKumancev/gRPC-nmap/internal/server"
	"github.com/DmitriiKumancev/gRPC-nmap/pkg/logger"
)

func main() {

	log := logger.New()

	configPath := "configs/config.json"

	cfg, err := config.New(configPath)
	if err != nil {
		log.Error("failed to load config: %v", err)
	}

	if err := server.Run(log, cfg); err != nil {
		log.Fatal("failed to run server: %v", err)
	}
}
