# Variables
APP_NAME := hr-system
DOCKER_COMPOSE := docker-compose
GO := go
MIGRATE := cd migrate; atlas migrate diff --env gorm; cd ..;

# Build the Go binary
.PHONY: build
build:
	$(GO) build -o bin/$(APP_NAME) main.go

# Run the server locally
.PHONY: run
run:
	$(GO) run main.go

# Run seeds locally
.PHONY: seed
run:
	$(GO) run HRSystem/cmd/seed/main.go

# Run migrations
.PHONY: migrate
migrate:
	$(MIGRATE)

# Run tests
.PHONY: test
test:
	$(GO) test -tags=unit ./... -v

# Clean the build
.PHONY: clean
clean:
	rm -rf bin/

# Build and run the application using Docker Compose
.PHONY: docker-up
docker-up:
	$(DOCKER_COMPOSE) up --build

# Stop and remove Docker containers
.PHONY: docker-down
docker-down:
	$(DOCKER_COMPOSE) down

# Restart Docker Compose
.PHONY: docker-restart
docker-restart: docker-down docker-up

# Lint
.PHONY: lint
lint:
	golangci-lint run

# Build swagger docs
.PHONY: swagger
swagger:
	swag init --parseDependency --overridesFile ./docs/.swaggo

.PHONY: mockery
mockery:
	mockery --config mockery.yaml