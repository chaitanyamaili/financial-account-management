# Variables
COMPOSE_FILE := docker-compose.yaml
PROJECT_NAME := financial_system

# Default target: Show help
.PHONY: help
help:
	@echo "Makefile for managing Docker Compose commands"
	@echo "Usage:"
	@echo "  make up            Start all services (backend, database, frontend)"
	@echo "  make down          Stop all services"
	@echo "  make restart       Restart all services"
	@echo "  make logs          Show logs for all services"
	@echo "  make build         Build or rebuild services"
	@echo "  make api-logs      Show logs for the backend service"
	@echo "  make db-logs       Show logs for the database service"
	@echo "  make frontend-logs Show logs for the frontend service"
	@echo "  make clean         Stop and remove containers, networks, and volumes"

# Start all services
.PHONY: up
up:
	@docker-compose -f $(COMPOSE_FILE) -p $(PROJECT_NAME) up -d

# Stop all services
.PHONY: down
down:
	@docker-compose -f $(COMPOSE_FILE) -p $(PROJECT_NAME) down

# Restart all services
.PHONY: restart
restart:
	@docker-compose -f $(COMPOSE_FILE) -p $(PROJECT_NAME)

# Show logs for all services
.PHONY: logs
logs:
	@docker-compose -f $(COMPOSE_FILE) -p $(PROJECT_NAME) logs -f

# Build or rebuild services
.PHONY: build
build:
	@docker-compose -f $(COMPOSE_FILE) -p $(PROJECT_NAME) build

# Show logs for the backend service
.PHONY: api-logs
backend-logs:
	@docker-compose -f $(COMPOSE_FILE) -p $(PROJECT_NAME) logs -f api

# Show logs for the database service
.PHONY: db-logs
db-logs:
	@docker-compose -f $(COMPOSE_FILE) -p $(PROJECT_NAME) logs -f db

# Show logs for the frontend service
.PHONY: clean
clean:
	@docker-compose -f $(COMPOSE_FILE) -p $(PROJECT_NAME) down -v
