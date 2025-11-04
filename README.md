# svc-init

A pair of Go microservice templates (flat and standard) to jumpstart backend microservice projects.
The repository supports both REST and gRPC service templates.

# which one to use?

- `flat/` - a compact, shallow layout suitable for small/mid projects or when you prefer fewer directories. Use when you want a minimal structure and faster bootstrapping.
- `standard/` - a conventional layout that groups application code under `internal/` for better encapsulation in
  larger codebases. Use when you expect growth, need stricter package boundaries.

Personally im a big fan of flat layouts, and use standard only when the project is expected to grow significantly.
Flat uses group by dependency pattern from `https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1`, shot out to the author for that article.

# Getting started (choose layout)

1. install `svc-init` script in your path

```bash
chmod +x svc-init
cp svc-init /usr/local/bin/
```

2. cd to your desired project directory

```bash
cd /path/to/your/project
```

3. init git repo in the desired dir (automatically adjust go mod paths based on repo remote url)

```bash
git init
git remote add origin <your-repo-url>
```

4. run svc-init to copy the chosen template into the current directory. The script supports an
   optional `-n|--noupdate` flag to skip updating the template repository (useful for air-gapped
   environments or CI where templates are preinstalled in `/opt/svc-init`). Examples:

- `svc-init flat` — shorthand for `svc-init -rest flat` (default group is `rest`)
- `svc-init standard` — use `rest/standard`
- `svc-init -grpc flat` — use `grpc/flat`
- `svc-init -n -grpc standard` — use `grpc/standard` and do NOT update the local template clone

# USAGE

```bash
  svc-init flat
  svc-init standard
  svc-init -grpc standard
  svc-init -n -grpc flat
```

## Common Makefile targets

Each layout ships a `Makefile` with a common set of developer shortcuts. Exact flags and behavior may vary
slightly between `flat` and `standard` layouts — check the layout's `Makefile` for details — but the
following targets are present in most templates and are useful to know:

- `make build` — compile the service binary for local development (usually a platform-native build). Some
  templates include a cross-build invocation (GOOS/GOARCH) for producing a Linux binary for container images.

- `make run` — run the application locally (typically `go run ./cmd/api/main.go`). Use this for
  quick local testing.

- `make test` — run the unit test suite (`go test ./...`). Good to run before committing changes.

- `make fmt` — format Go sources (`go fmt ./...`). Keeps code consistent.

- `make vet` — run `go vet` static checks to catch suspicious code.

- Migration helpers (require the `golang-migrate` CLI):

  - `make migrate-create` — create a new migration file (usually prompts for a name).
  - `make migrate-up` — apply pending migrations to the database.
  - `make migrate-down` — roll back the last migration.
  - `make migrate-drop` — drop all objects managed by migrations.

- `make mocks` — run mock generation (uses `mockery` with a config file). See the `Mockery configuration` section
  below for details.

## Tooling & prerequisites

Make targets that interact with external tools require those CLI binaries to be installed and on your `PATH`.
Common prerequisites across layouts include:

- Go 1.21+
- `golang-migrate` CLI (for DB migration targets)
- `swag` (for OpenAPI generation)
- `protoc` and Go protobuf plugins (for gRPC templates / `make proto`)
- `mockery` — used to generate mocks for interfaces used in tests. Install mockery from `https://vektra.github.io/mockery/v3.5/installation/`.

If you prefer not to install it globally, you can run mockery via a container or point the Makefile at a
project-local config file (see below).

## Mockery configuration and usage

The templates provide a `mocks` target that runs `mockery` with a configuration file. The Makefile variable
used by the templates is `MOCKERY_CONFIG_PATH` and defaults to `~/.mockery.yml`. The layouts call:

```bash
mockery --config $(MOCKERY_CONFIG_PATH)
```

This repository includes a project-level `./.mockery.yml` which is a good starting point (you can copy it to
`~/.mockery.yml` or set `MOCKERY_CONFIG_PATH` to the local file). The included YAML contains the common options we use; here's
what each top-level option means:

```bash
MOCKERY_CONFIG_PATH=./.mockery.yml make mocks
```

The generated mocks appear under the configured `dir` (by default `./mocks`). Import them in tests as needed.

# Notes

- Many Makefile commands rely on environment variables (DB connection, ports, etc.). Copy `.env.example` to
  `.env` in your project and set the required values before running targets that touch networked services.
- Always inspect the `Makefile` in the layout you copied — the Makefile contains the authoritative list of
  commands and any layout-specific behavior.
- Both layouts aim to provide the same functionality; choose based on your organizational preferences.
- Both layouts use my library and chi for creating http services, since it is tool mainly created for myself.
