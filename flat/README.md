# Golang REST API Boilerplate

a production-ready, rest api boilerplate for Golang backend app. Built to accelerate modern backend development with robust architectural patterns, great developer experience, Docker-native workflows, and batteries-included tooling.

## Features

- **flat rest api structure** with idiomatic Go.
- **PostgreSQL** database and **Redis** cache integration. (easy to swap the storage, if you need to use for example mongo)
- **Database migrations** with `migrate`.
- **Swagger/OpenAPI** documentation tooling.
- **Docker \& Docker Compose** for easy, containerized development.
- **Environment-driven configuration** using `.env`.
- **Makefile** for common dev and CI tasks.
- **Ready for cloud-native deployment** on Kubernetes.
- **Testing** setup for unit/integration testing.

## directory structure

```
bin/                 # Compiled binaries
cmd/                 # Main application entrypoints
docker-compose.yml   # Multi-service dev orchestration
Dockerfile           # Main build definition
docs/                # Documentation (OpenAPI, etc.)
go.mod, go.sum       # Dependency definitions
internal/            # App logic: services, repos, handlers
migrations/          # DB migration SQL files
Makefile             # Developer tasks and CI
rest_init.sh         # Project bootstrap script
```

### 1. Prerequisites

- Go 1.21+ installed

````markdown
# golang rest api boilerplate — flat layout

This `flat` folder contains a minimal, flat variant of the REST API boilerplate. It's designed for small
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

3. Run the API locally (uses `cmd/api`):

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
http.go             # http requests and response models
docker-compose.yml  # dev orchestration
Dockerfile          # build definition
Makefile            # developer tasks
```

Notes

- If your project grows in complexity, consider migrating to the `standard` layout which places packages
  under `internal/` for better encapsulation.
- The `rest_init.sh` script in the repository can bootstrap a new project from this template.

License

See the repository `LICENSE` file for licensing details.

```
| `make migrate-down`   | Revert last applied migration                       |
```
