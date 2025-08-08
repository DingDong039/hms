.PHONY: help env tidy build run test test-handlers docker-up docker-build docker-down docker-logs docker-restart

help:
	@echo "Available targets:"
	@echo "  env             Create .env from .env.example if missing"
	@echo "  tidy            Run go mod tidy"
	@echo "  build           Build the app (Docker)"
	@echo "  run             Run locally: go run cmd/main/main.go"
	@echo "  test            Run all tests"
	@echo "  test-handlers   Run handler tests with -v"
	@echo "  docker-up       Start services (detached)"
	@echo "  docker-build    Build and start services"
	@echo "  docker-down     Stop and remove services"
	@echo "  docker-logs     Tail docker logs"
	@echo "  docker-restart  Restart services"

env:
	@test -f .env || cp .env.example .env
	@# Ensure NGINX_PORT avoids 80 by default
	@sed -i '' -e 's/^NGINX_PORT=.*/NGINX_PORT=8081/' .env || true

 tidy:
	go mod tidy

 build:
	docker-compose build

 run:
	go run cmd/main/main.go

 test:
	go test ./...

 test-handlers:
	go test -v ./tests/handlers

 docker-up:
	docker-compose up -d

 docker-build:
	docker-compose up -d --build

 docker-down:
	docker-compose down

 docker-logs:
	docker-compose logs -f

 docker-restart:
	docker-compose down && docker-compose up -d --build
