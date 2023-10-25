package config

import (
	"encoding/json"
	"os"
)

type (
	Config struct {
		ServerGrpc ServerGrpc `json:"grpc_server"`
		ClientGrpc ClientGrpc `json:"grpc_client"`
	}

	ServerGrpc struct {
		Port string `json:"port"`
	}

	ClientGrpc struct {
		Port string `json:"port"`
	}
)

func New(path string) (*Config, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &Config{}
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}
