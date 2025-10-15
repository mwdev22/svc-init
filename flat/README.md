# Golang REST API Boilerplate — Flat layout

A compact variant of the REST API boilerplate designed for small projects or as a minimal starting point.
It provides the same core features as the `standard` layout but with a shallower directory structure.

## Features

- Flat REST API structure with idiomatic Go
- Swagger/OpenAPI documentation tooling
- Docker & Docker Compose for easy, containerized development
- Makefile for common dev and CI tasks
- Ready for cloud-native deployment on Kubernetes
- Testing setup for unit/integration testing

## Prerequisites

- Go 1.21+ installed

## Flat layout for a Golang REST API

`flat` is a minimal layout. Use it when you prefer fewer nested directories or when starting a small service.

## Quick start

1. Copy `.env.example` to `.env` (if present) and set database/redis credentials and ports.
2. Start services with Docker Compose:

```bash
docker-compose up --build
```

3. Run the API locally:

```bash
make run
```

## Makefile shortcuts

- `make build` — compile the binary
- `make run` — run the server
- `make test` — run unit tests
- `make fmt` — format code
- `make vet` — static checks
- `make migrate-create` / `migrate-up` / `migrate-down` / `migrate-drop` — DB migrations
- `make swag` — generate swagger docs (requires `swag`)

## Layout (important files)

```
cmd/                # application entrypoints (cmd/api/main.go)
api/                # http handlers and routing helpers
service/            # business logic
models.go           # domain models
service.go          # service and external dependencies interfaces
http.go             # http helpers / request/response models
docker-compose.yml  # dev orchestration
Dockerfile          # build definition
Makefile            # developer tasks
```

## Notes

- If your project grows in complexity, consider migrating to the `standard` layout which places packages
  under `internal/` for better encapsulation.
- The `rest_init` script in the repository can bootstrap a new project from this template.

## License

See the repository `LICENSE` file for licensing details.
