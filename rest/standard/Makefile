# Include .env file
include ./.env
export $(shell sed 's/=.*//' .env)

# Project variables
PROJECT_NAME := PROJECT_NAME
PKG := ./...
MAIN := ./cmd/api/
MIGRATIONS_DIR = migrations


# Go commands
BUILD := go build
CLEAN := go clean
FMT := go fmt
VET := go vet
TEST := go test
RUN := go run
MIGRATE := migrate

# Targets
.PHONY: all build clean fmt vet test run migrate-create migrate-up migrate-down migrate-drop

all: fmt vet test build

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(BUILD) -o $(PROJECT_NAME).exe $(MAIN)

clean:
	$(CLEAN)

fmt:
	$(FMT) $(PKG)

vet:
	$(VET) $(PKG)

test:
	$(TEST) $(PKG)

run:
	$(RUN) $(MAIN)
migrate-create:
	@export DB_SOURCE=$$(if [ "$(DB_TYPE)" = "sqlite" ]; then \
		echo "sqlite3://$(DB_NAME).sqlite"; \
	else \
		echo "sqlserver://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):1433?database=$(DB_NAME)&sslmode=disable"; \
	fi); \
	read -p "Enter migration name: " name; \
	$(MIGRATE) create -ext sql -dir $(MIGRATIONS_DIR) -seq $$name

migrate-up:
	@export DB_SOURCE=$$(if [ "$(DB_TYPE)" = "sqlite" ]; then \
		echo "sqlite3://$(DB_NAME).sqlite"; \
	else \
		echo "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):5432?database=$(DB_NAME)&sslmode=disable"; \
	fi); \
	$(MIGRATE) -database $$DB_SOURCE -path $(MIGRATIONS_DIR) up


migrate-down:
	@export DB_SOURCE=$$(if [ "$(DB_TYPE)" = "sqlite" ]; then \
		echo "sqlite3://$(DB_NAME).sqlite"; \
	else \
		 echo "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):5432?database=$(DB_NAME)&sslmode=disable"; \
	fi); \
	$(MIGRATE) -database $$DB_SOURCE -path $(MIGRATIONS_DIR) down 1;

migrate-force:
	@export DB_SOURCE=$$(if [ "$(DB_TYPE)" = "sqlite" ]; then \
		echo "sqlite3://$(DB_NAME).sqlite"; \
	else \
		echo "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):5432?database=$(DB_NAME)&sslmode=disable"; \
	fi); \
	$(MIGRATE) -database $$DB_SOURCE -path $(MIGRATIONS_DIR) force $(VERSION);

migrate-drop:
	@export DB_SOURCE=$$(if [ "$(DB_TYPE)" = "sqlite" ]; then \
		echo "sqlite3://$(DB_NAME).sqlite"; \
	else \
		echo "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):5432?database=$(DB_NAME)&sslmode=disable"; \
	fi); \
	$(MIGRATE) -database $$DB_SOURCE -path $(MIGRATIONS_DIR) drop

swag:
	@swag init --generalInfo ./cmd/main.go
