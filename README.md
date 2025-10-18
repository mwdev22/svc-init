# svc-init

a pair of Go microservice boilerplates (flat and standard) to jumpstart backend services.

currently only for api's exposed via rest, im going to add grpc support too.

repository contains two alternative layouts you can use as a starting point:

- `flat/` — a compact, shallow layout suitable for small projects or when you prefer fewer directories.
- `standard/` — a conventional layout that groups application code under `internal/` for better encapsulation in
  larger codebases.

which one to use?

- `flat/` when you want a minimal structure and faster bootstrapping.
- `standard/` when you expect growth, need stricter package boundaries

Getting started (choose layout)

1. install `svc-init` script in your path

2. cd to your desired project directory

3. init git directory in the desired dir (automatically adjust go mod paths based on repo remote url)

4. run svc-init <flat|standard>

5. be happy because you dont need to write it all again

common tasks

- Build the binary: `make build`
- Run tests: `make test`
- Apply DB migrations (requires `migrate`): `make migrate-up`
- Generate Swagger docs (requires `swag`): `make swag`

- `rest_init.sh` (present in the layouts) will copy template files into the current directory to initialize a
  new project. Edit it as needed before running.

Notes

- Both layouts aim to provide the same functionality; choose based on your organizational preferences.
- Both layouts use my library and chi for creating http services, since it is tool mainly created for myself.
