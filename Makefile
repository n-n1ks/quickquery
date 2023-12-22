install:
	go mod download

generate-mocks:
	mockgen -source=internal/database/storage/engine/engine.go -destination=internal/database/storage/engine/mock_engine/engine.go -package=mock_engine Engine

test:
	go test -v ./...

lint:
	golangci-lint run

run-server:
	go run cmd/quickquery/main.go
