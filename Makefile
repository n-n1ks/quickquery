install:
	go mod download

generate-mocks:
	mockgen -source=internal/database/storage/engine/engine.go -destination=internal/database/storage/engine/mock_engine/engine.go -package=mock_engine Engine && \
	mockgen -source=internal/database/database.go -destination=internal/database/mock/database_mocks.go -package=database_mock ComputerLayer StorageLayer

test:
	go test -v -race -parallel=4 ./...

lint:
	golangci-lint run

run-server:
	go run cmd/quickquery/main.go
