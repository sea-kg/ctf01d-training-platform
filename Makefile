.PHONY: lint install build run-server run-db attach-db fuzz-api

# Lint the code with golangci-lint
lint:
	make fmt; \
	if ! [ -x "$$(command -v golangci-lint)" ]; then \
		docker run --rm -v $(PWD):/app -w /app golangci/golangci-lint:latest golangci-lint run -v; \
	else \
		golangci-lint run; \
	fi

# Install requirements
install:
	go mod download
	go mod tidy

# Build the server executable
build:
	go build cmd/ctf01d/main.go

# Build the server executable in docker
build-in-docker:
	docker run --rm -v $(PWD):/app -w /app golang:1.22-bookworm go build cmd/ctf01d/main.go

# format go files
fmt:
	go fmt ./internal/...; \
	go fmt ./cmd/...;

# Run the local development server
run-server:
	go run cmd/ctf01d/main.go

# Run PostgreSQL container for local development
run-db:
	@if [ $$(docker ps -a -q -f name=ctf01d-postgres) ]; then \
		echo "Container ctf01d-postgres already exists. Restarting..."; \
		docker start ctf01d-postgres; \
	else \
		echo "Creating and starting container ctf01d-postgres..."; \
		docker run --rm -d \
			-v $(PWD)/docker_tmp/pg_data:/var/lib/postgresql/data/ \
			--name ctf01d-postgres \
			-e POSTGRES_DB=ctf01d_training_platform \
			-e POSTGRES_USER=postgres \
			-e POSTGRES_PASSWORD=postgres \
			-e PGPORT=4112 \
			-p 4112:4112 postgres:16.3; \
	fi

# Stop PostgreSQL container
stop-db:
	@if [ $$(docker ps -q -f name=ctf01d-postgres) ]; then \
		echo "Stopping container ctf01d-postgres..."; \
		docker stop ctf01d-postgres; \
	else \
		echo "Container ctf01d-postgres is not running."; \
	fi

# cleanup db and restart db and rebuild main app
reset-db:
	make stop-db; \
	sudo rm -rf docker_tmp/pg_data; \
	make run-db; \
	make build;

# Setup test database and run tests
test-db:
	# Ensure the PostgreSQL container is running
	@if ! [ $$(docker ps -q -f name=ctf01d-postgres) ]; then \
		echo "Starting PostgreSQL container..."; \
		docker start ctf01d-postgres; \
	fi
	# Remove test database if it exists
	@docker exec -it ctf01d-postgres psql -U postgres -c "DROP DATABASE IF EXISTS ctf01d_training_platform_test;"
	# Create a new test database
	@docker exec -it ctf01d-postgres psql -U postgres -c "CREATE DATABASE ctf01d_training_platform_test;"
	# Run the tests
	@go test -v ./test/server_integration_test.go

# Revome PostgreSQL container
remove-db:
	@if [ $$(docker ps -a -q -f name=ctf01d-postgres) ]; then \
		echo "Removing container ctf01d-postgres..."; \
		docker rm -f ctf01d-postgres; \
	else \
		echo "Container ctf01d-postgres does not exist."; \
	fi

# Attach to the running PostgreSQL container
attach-db:
	docker exec -it ctf01d-postgres psql -U postgres -d ctf01d_training_platform

# Generate Go server boilerplate from OpenAPI 3
codegen:
	oapi-codegen -generate models,chi -o internal/app/server/server.gen.go --package server api/openapi.yaml
