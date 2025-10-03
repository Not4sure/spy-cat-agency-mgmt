
.PHONY: start
start:
	@if podman-compose up -d>/dev/null; then \
		: ; \
	else \
		docker-compose up -d; \
	fi

.PHONY: run
run:
	@go run ./cmd/app/main.go

# Run unit tests
.PHONY: test
test:
	@echo "Testing..."
	@go test ./... -v

migrateup:
	@migrate -database "postgres://postgres:example@localhost:5432/test?sslmode=disable" -path ./internal/common/db/migrations up

migratedown:
	@migrate -database "postgres://postgres:example@localhost:5432/test?sslmode=disable" -path ./internal/common/db/migrations down
