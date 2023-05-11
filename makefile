run:
	@go run cmd/main.go

lint:
	@golangci-lint run

coverage:
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html