# Rest API Service Templates

directory contains `Rest API` service templates in two layouts: `flat` (minimal) and `standard` (production-friendly).
These templates provide a ready-made starting point for HTTP/REST services in Go with sensible defaults for
configuration, migrations, swagger, and local development.

## What this layout uses

- idiomatic Go project layout (`flat` vs `standard`) and Go modules
- HTTP routing and handler using my internal `rest` library and `chi` router
- Swagger/OpenAPI for API documentation and code generation (using `swaggo/swag`)
- Docker & Docker Compose for local development
- A Makefile with common developer tasks (build, run, test, migrations, docs, mocks)

## Prerequisites

- Go 1.21+ (project uses Go modules)
- `swag` for OpenAPI generation (if you plan to use swagger):

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

## Quick start (development)

1. Copy `.env.example` to `.env` and set DB/Redis/port values.
2. Start the development stack with Docker Compose (if provided):

```bash
docker-compose up --build
```

3. Run the API locally:

```bash
make run
```

## Makefile targets (rest-specific)

- `make swag` — generate swagger/openapi docs (requires `swag`)

## project layout

```
cmd/                # application entrypoints
api/                # http handlers, routing and middleware
service             # business logic
models              # domain models
Dockerfile          # build definition
docker-compose.yml  # development orchestration
Makefile            # developer tasks including migrations and docs
```

`flat` vs `standard`:

- `flat` keeps a shallower layout (good for small services).
- `standard` groups code under `internal/` for better encapsulation in larger projects.

## Dependencies used by the templates

Typical dependencies (check `go.mod` inside the template you copied):

- `github.com/swaggo/swag` for OpenAPI generation
- `github.com/mwdev22/gocfg` — configuration helper used in examples
- `github.com/mwdev22/rest` - internal helper library used by the template to create and configure HTTP servers

# Note

Personally im a big fan of flat layouts, and use standard only when the project is expected to grow significantly.
