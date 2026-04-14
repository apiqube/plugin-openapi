# plugin-openapi

> OpenAPI / Swagger test generator for [ApiQube](https://github.com/apiqube).

[![License](https://img.shields.io/badge/License-MIT-blue?style=flat-square)](LICENSE)
[![Status](https://img.shields.io/badge/Status-Scaffold-yellow?style=flat-square)]()

Generator plugin (not an executor). Invoked via `qube generate --from openapi.json`
to produce qube test files from an existing API specification.

Detects CRUD patterns and emits scenario files with proper save/extract chains.

## Usage

```bash
# From a local file
qube generate --from ./swagger.json --output ./tests/

# From a URL
qube generate --from http://localhost:8081/swagger/doc.json --output ./tests/
```

## What it generates

For each `paths` entry in the spec:

- CRUD groups (POST + GET/{id} + PUT/{id} + DELETE/{id}) → scenario file
- Schema fields → `fake.*` template values
- Response codes → `expect.status` assertions
- Required fields → test bodies

## Build

```bash
tinygo build -o plugin-openapi.wasm -target=wasi ./
```

## License

[MIT](LICENSE)
