
.PHONY: run
run:
	@go run ./cmd/app/main.go

# Run unit tests
.PHONY: test
test:
	@echo "Testing..."
	@go test ./... -v
