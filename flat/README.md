# Golang REST API Boilerplate

a production-ready, rest api boilerplate for Golang backend app. Built to accelerate modern backend development with robust architectural patterns, great developer experience, Docker-native workflows, and batteries-included tooling.

## Features

- **flat rest api structure** with idiomatic Go.
- **Swagger/OpenAPI** documentation tooling.
- **Docker \& Docker Compose** for easy, containerized development.
- **Makefile** for common dev and CI tasks.
- **Ready for cloud-native deployment** on Kubernetes.
- **Testing** setup for unit/integration testing.

### 1. Prerequisites

- Go 1.21+ installed

````markdown
# flat layout for golang rest api

`Flat`, minimal variant of the REST API boilerplate. It's designed for small
projects or to act as a compact starting point when you prefer a shallower directory layout compared to the
`standard` layout.

What you get

- A working REST API scaffold in Go (entry: `cmd/api/main.go`).
- Docker + Docker Compose development stack (Postgres + Redis examples).
- DB migration integration via `migrate`.
- Swagger/OpenAPI generation with `swag` (optional).
- A simple `Makefile` with common dev tasks.

1. Copy `.env.example` (if present) to `.env` and set database/redis credentials and ports.
2. Start services with Docker Compose:

```bash
docker-compose up --build
```
````

3. Run the API locally:

```bash
make run
```

Makefile shortcuts

- `make build` — compile the binary
- `make run` — run the server
- `make test` — run unit tests
- `make fmt` — format code
- `make vet` — static checks
- `make migrate-create` / `migrate-up` / `migrate-down` / `migrate-drop` — DB migrations
- `make swag` — generate swagger docs (requires `swag`)

Layout (important files)

```
cmd/                # application entrypoints (cmd/api/main.go)
api/                # http handlers and routing helpers
service/            # business logic
models.go           # domain models
service.go          # service and external dependecies interfaces
http.go             # http requests and response models
docker-compose.yml  # dev orchestration
Dockerfile          # build definition
Makefile            # developer tasks
```

Notes

- if your project grows in complexity, consider migrating to the `standard` layout which places packages
  under `internal/` for better encapsulation.
- the `rest_init` script in the repository can bootstrap a new project from this template.

License

see the repository `LICENSE` file for licensing details.

```
| `make migrate-down`   | Revert last applied migration                       |
```
