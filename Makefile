.PHONY: lint install build run-server run-db attach-db fuzz-api

# Lint the code with golangci-lint
lint:
	go fmt ./...
	@if ! [ -x "$$(command -v golangci-lint)" ]; then \
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

build-in-docker:
	docker run --rm -v $(PWD):/app -w /app golang:1.22-bookworm go build cmd/ctf01d/main.go

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
		docker run -d --name ctf01d-postgres -e POSTGRES_DB=ctf01d -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 postgres; \
	fi

# Stop PostgreSQL container
stop-db:
	@if [ $$(docker ps -q -f name=ctf01d-postgres) ]; then \
		echo "Stopping container ctf01d-postgres..."; \
		docker stop ctf01d-postgres; \
	else \
		echo "Container ctf01d-postgres is not running."; \
	fi

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
	docker exec -it ctf01d-postgres psql -U postgres

# Generate Go server boilerplate from OpenAPI 3
codegen:
	oapi-codegen -generate models -o internal/app/apimodels/models.gen.go --package models api/swagger.yaml
