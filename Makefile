build:
	go build ./cmd/server/main.go

lint:
	golangci-lint run

test:
	go test -v ./internal/usecase/netvuln_test.go
server:
	go run ./cmd/server/main.go
	
client:
	go run ./cmd/client/main.go