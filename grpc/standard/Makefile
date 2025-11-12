# Include .env file
include ./.env
export $(shell sed 's/=.*//' .env)

# Project variables
PROJECT_NAME := PROJECT_NAME
PKG := ./...
MAIN := ./cmd/api/
MIGRATIONS_DIR = migrations
PROTO_DIR = proto

MOCKS_DIR := ./mocks
MOCKERY_CONFIG_PATH := ~/.mockery.yml

FMT := go fmt
VET := go vet
TEST := go test
MIGRATE := migrate

# Targets
.PHONY: all build fmt vet test run migrate-create migrate-up migrate-down migrate-drop

all: fmt vet test build

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(PROJECT_NAME).exe $(MAIN)
fmt:
	 go fmt $(PKG)

vet:
	 go vet $(PKG)

test:
	go test -v $(PKG)

run:
	go run $(MAIN)

migrate-create:
	read -p "Enter migration name: " name; \
	$(MIGRATE) create -ext sql -dir $(MIGRATIONS_DIR) -seq $(DATABASE_URI)

migrate-up:
	$(MIGRATE) -database $(DATABASE_URI) -path $(MIGRATIONS_DIR) up

migrate-down:
	$(MIGRATE) -database $(DATABASE_URI) -path $(MIGRATIONS_DIR) down 1;

migrate-force:
	$(MIGRATE) -database $(DATABASE_URI) -path $(MIGRATIONS_DIR) force $(VERSION);

migrate-drop:
	$(MIGRATE) -database $(DATABASE_URI) -path $(MIGRATIONS_DIR) drop

proto:
	@mkdir -p gen
	@protoc -I=$(PROTO_DIR) \
		--go_out=paths=source_relative:internal/gen \
		--go-grpc_out=paths=source_relative:internal/gen \
		$(PROTO_DIR)/*.proto
	@echo "--- PB GENERATED ---"

mocks:
	mockery --config $(MOCKERY_CONFIG_PATH)