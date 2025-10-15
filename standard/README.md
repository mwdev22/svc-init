# Golang REST API Boilerplate

A production-ready, REST API boilerplate for Golang backend app. Built to accelerate modern backend development with robust architectural patterns, great developer experience, Docker-native workflows, and batteries-included tooling.

## Features

- **Modern RESTful API structure** with idiomatic Go.
- **PostgreSQL** database and **Redis** cache integration. (easy to swap the storage, if you need to use for example mongo)
- **Database migrations** with `migrate`.
- **Swagger/OpenAPI** documentation tooling.
- **Docker \& Docker Compose** for easy, containerized development.
- **Environment-driven configuration** using `.env`.
- **Simple Makefile** for common dev and CI tasks.
- **Clean internal architecture** (service, repository, handlers).

````markdown
# Golang REST API Boilerplate — Standard layout

This `standard` folder contains the opinionated, production-friendly layout that groups application code
under `internal/` for stronger encapsulation and clearer boundaries. Use this layout when you expect the
project to grow or when you want to follow common Go project conventions.

What changed recently

- Updated README to match the improved style used by the `flat` layout.
- Makefile targets and migration commands remain the same; see Makefile for DB driver specifics.

Highlights

- Entrypoint: `cmd/api/main.go`.
- Project layout follows Go best-practices (internal packages for app logic).
- Docker and Docker Compose examples included for local development.

Quick start

1. Copy `.env.example` to `.env` and set DB/Redis credentials.
2. Start the development stack with Docker Compose:

```bash
docker-compose up --build
```
````

3. Run the API locally:

```bash
make run
```

Makefile targets

See the `Makefile` for the precise list of targets. Typical commands include:

- `make build`, `make run`, `make test`, `make fmt`, `make vet`
- Migration helpers: `make migrate-create`, `make migrate-up`, `make migrate-down`, `make migrate-drop`
- `make swag` to generate OpenAPI docs (requires `swag` to be installed)

Project layout (important files)

```
cmd/                # application entrypoints (cmd/api/main.go)
internal/           # app logic: app, handlers, models, service, store
docker-compose.yml  # dev orchestration
Dockerfile          # build definition
Makefile            # developer tasks and migration helpers
```

Notes

- The `standard` layout is recommended for larger codebases because `internal/` prevents accidental imports
  from other modules and better communicates package boundaries.
- If you prefer the shallow layout, see the `flat/` folder which provides the same features with fewer
  directories.

Contributing

Please open PRs or issues on the repository for bugs, improvements, and feature requests.

License

See the repository `LICENSE` file for licensing details.

```

```
