# gRPC Service Templates

this directory contains gRPC service templates in two layouts: `flat` and `standard`.
They are intended to accelerate building gRPC-based services in Go with protobuf-first development and a small set
of pragmatic conventions.

## what this layout uses

- protobuf for message and service definitions (`.proto` files under `proto/`).
- `protoc` with Go plugins (`protoc-gen-go` and `protoc-gen-go-grpc`) to generate Go code into the `gen/` directory.
- `google.golang.org/grpc` for the server and client runtime.
- small shared helper library `grpclib` (used in templates) that wraps server bootstrapping, common interceptors and options.
- Docker/Docker Compose and a `Makefile` for local development tasks (build, run, generate protos, migrations, mocks).

## Prerequisites

- Go 1.21+ (project uses Go modules)
- protoc (Protocol Buffers compiler).
- go protoc plugins (install to $GOBIN or $(go env GOPATH)/bin):

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## How protos are generated here

both templates include a `proto/` directory and a `Makefile` target `make proto` that:

- creates `gen/` if needed
- runs `protoc` with `--go_out=paths=source_relative:gen` and `--go-grpc_out=paths=source_relative:gen`

this produces `*.pb.go` files under `grpc/<layout>/gen/<pkg>` preserving source-relative imports. Proto files in the
template should include a `go_package` option to control the generated Go package path.

contracts should be defined inside `/proto` directory.

```bash
make proto
```

## quick start (development)

1. set `.env`.
2. generate protobuf code
3. start the local dev stack after adjusting dependencies and docker files.

```bash
docker-compose up --build
```

4. run the service locally:

```bash
make run
```

## Makefile targets (grpc specific)

- `make proto` — generate Go code from `proto/` contract files into `gen/` (requires protoc + plugins)

## project layout (what to expect)

```
cmd/                # application entrypoints (cmd/api/main.go)
proto/              # .proto definitions for services and messages
gen/                # generated protobuf Go code (created by make proto)
api/                # template API wiring that registers services and starts server
service/            # service interfaces (eg. KafkaProducer, MailSender)
models/             # domain models (if present)
Dockerfile          # build definition
docker-compose.yml  # development orchestration (optional)
Makefile            # developer tasks, including `proto` target
```

`flat` vs `standard`:

- `flat` keeps a shallower layout (good for small services).
- `standard` groups code under `internal/` for better encapsulation in larger projects.

## Dependencies used by the templates

The templates typically depend on:

- `google.golang.org/grpc` — gRPC runtime
- `google.golang.org/protobuf` — protobuf runtime
- `github.com/mwdev22/grpclib` — internal helper library used by the template to create and configure servers
- `github.com/mwdev22/gocfg` — configuration helper (used in examples)

# Note

Personally im a big fan of flat layouts, and use standard only when the project is expected to grow significantly.
