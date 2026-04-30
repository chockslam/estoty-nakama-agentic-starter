# Research Log

## Purpose

This file records technical research that affects implementation decisions.

Agents must update this file when making decisions about:

- Go version;
- Nakama version;
- `nakama-common` version;
- Docker/plugin-builder strategy;
- RPC calling conventions;
- server-to-server/private RPC authorization;
- testing strategy;
- PostgreSQL/Nakama configuration.

Prefer official sources.

## Entry template

### YYYY-MM-DD - Topic

Source:

- Title:
- URL:
- Publisher/owner:
- Date checked:

Relevant finding:

Decision:

Risk / uncertainty:

Follow-up:

---

## Initial required research items

### Nakama Go runtime

Research:

- How Go runtime modules are initialized.
- How RPCs are registered.
- Required function signatures.
- Plugin builder requirements.
- Version compatibility requirements.

Decision to record:

- selected Nakama version;
- selected `nakama-common` version;
- selected plugin-builder image;
- whether Go runtime code is compiled as a shared object.

---

### Nakama Docker/PostgreSQL setup

Research:

- official Docker Compose pattern;
- required ports;
- PostgreSQL connection settings;
- local config file approach.

Decision to record:

- `docker-compose.yml` layout;
- local credentials;
- exposed ports;
- whether local defaults are safe to commit.

---

### Nakama RPC registration

Research:

- how to register Go runtime RPCs;
- RPC handler signatures;
- how to return JSON;
- how to return errors.

Decision to record:

- final RPC registration approach;
- handler function organization;
- helper function extraction strategy.

---

### Server-to-server/private RPC behavior

Research:

- how Nakama distinguishes server-to-server calls from user-session calls;
- how runtime HTTP key calls are made;
- how to reject user-session calls.

Decision to record:

- private RPC protection strategy;
- expected positive and negative tests.

---

### Go testing and style

Research:

- current stable Go release;
- Go formatting/style guidance;
- table-driven testing guidance;
- whether fuzz tests are useful for JSON payload parsing;
- compatibility with Nakama plugin build.

Decision to record:

- Go version used locally;
- Go version expected inside Nakama/plugin build;
- test strategy;
- commands required before completion.

---

## Research entries

Add entries below this line during Phase 0.

### 2026-04-30 - Nakama Go runtime and compatibility baseline

Source:

- Title: Go Runtime - Heroic Labs Documentation
  URL: https://heroiclabs.com/docs/nakama/server-framework/go-runtime/
  Publisher/owner: Heroic Labs
  Date checked: 2026-04-30
- Title: Dependency Pinning - Heroic Labs Documentation
  URL: https://heroiclabs.com/docs/nakama/server-framework/go-runtime/go-dependencies/
  Publisher/owner: Heroic Labs
  Date checked: 2026-04-30
- Title: Release Notes - Heroic Labs Documentation
  URL: https://heroiclabs.com/docs/nakama/getting-started/release-notes/
  Publisher/owner: Heroic Labs
  Date checked: 2026-04-30

Relevant finding:

- Go runtime modules are loaded as shared objects from the runtime path.
- `InitModule` is the required entry point used to register RPCs, hooks, and other runtime functions.
- Heroic Labs documents `go build --trimpath --mod=vendor --buildmode=plugin` as the shared-object build path.
- Go runtime code should avoid goroutines and shared mutable state because multi-node concurrency is not supported.
- The official compatibility matrix maps Nakama 3.37.0 to `github.com/heroiclabs/nakama-common/runtime@v1.44.2`.
- The release notes for 3.37.0 confirm `nakama-common` v1.44.2 for this release.

Decision:

- Pin Nakama to `3.37.0`.
- Pin `github.com/heroiclabs/nakama-common/runtime` to `v1.44.2`.
- Use a Go runtime plugin compiled as a `.so` via the Nakama plugin-builder flow.
- Keep runtime code small, deterministic, and testable with pure helpers.

Risk / uncertainty:

- The exact Go toolchain used to build the 3.37.0 Nakama image is not stated in the release notes; confirm from startup logs in Phase 2 before relying on any host toolchain assumptions.

Follow-up:

- In Phase 2, verify the runtime startup log and record the Go version reported by Nakama.

---

### 2026-04-30 - Docker + PostgreSQL local setup

Source:

- Title: Docker Compose - Heroic Labs Documentation
  URL: https://heroiclabs.com/docs/nakama/getting-started/install/docker/
  Publisher/owner: Heroic Labs
  Date checked: 2026-04-30
- Title: Docker Configuration - Heroic Labs Documentation
  URL: https://heroiclabs.com/docs/nakama/getting-started/configuration/docker-configuration/
  Publisher/owner: Heroic Labs
  Date checked: 2026-04-30
- Title: Configuration - Heroic Labs Documentation
  URL: https://heroiclabs.com/docs/nakama/getting-started/configuration/
  Publisher/owner: Heroic Labs
  Date checked: 2026-04-30
- Title: postgres - Official Image | Docker Hub
  URL: https://hub.docker.com/_/postgres/
  Publisher/owner: PostgreSQL Docker Community
  Date checked: 2026-04-30
- Title: Compose file reference | Docker Docs
  URL: https://docs.docker.com/compose/compose-file/
  Publisher/owner: Docker
  Date checked: 2026-04-30

Relevant finding:

- Heroic Labs provides a Docker Compose pattern for Nakama with a PostgreSQL service and a Nakama service that reads a mounted YAML config file.
- Nakama config is supplied with the `--config` flag and `runtime.path` defaults to the modules directory under `data_dir`.
- The official Nakama docs list the local ports that matter here: HTTP 7350, console HTTP 7351, gRPC 7349, and console gRPC 7348.
- The official PostgreSQL image requires a password and supports a standard Compose setup with a dedicated `postgres` service and persisted volume.
- Docker Compose uses the Compose Specification; no legacy format feature is required for this project.

Decision:

- Use Docker Compose with separate Nakama and PostgreSQL services.
- Mount a local YAML config into the Nakama container and point to it with `--config`.
- Expose/document the Nakama ports 7350, 7351, 7349, and 7348.
- Use local-only credentials and keys in the developer setup; do not treat them as production values.

Risk / uncertainty:

- The exact compose implementation will be finalized in the runtime skeleton phase, but the official pattern is clear enough for Phase 0.

Follow-up:

- In Phase 2, verify the stack boots with PostgreSQL and Nakama using the documented ports and config mount approach.

---

### 2026-04-30 - Server-to-server / private RPC behavior

Source:

- Title: Introduction - Heroic Labs Documentation
  URL: https://heroiclabs.com/docs/nakama/server-framework/introduction/
  Publisher/owner: Heroic Labs
  Date checked: 2026-04-30
- Title: Runtime Context - Heroic Labs Documentation
  URL: https://heroiclabs.com/docs/nakama/server-framework/introduction/runtime-context/
  Publisher/owner: Heroic Labs
  Date checked: 2026-04-30
- Title: Configuration - Heroic Labs Documentation
  URL: https://heroiclabs.com/docs/nakama/getting-started/configuration/
  Publisher/owner: Heroic Labs
  Date checked: 2026-04-30

Relevant finding:

- Nakama runtime functions receive a context object.
- `RUNTIME_CTX_USER_ID` is not present in server-to-server calls.
- Heroic Labs explicitly recommends checking whether the context has a user ID to distinguish client calls from server-to-server calls.
- `runtime.http_key` protects server runtime HTTP invocations.

Decision:

- Implement `private_health_check` so it fails when a user ID is present in context and succeeds only through the server-to-server/runtime-key path.
- Use the runtime HTTP key path for the positive integration check and a user-session call for the negative check.

Risk / uncertainty:

- The exact request shape for the positive runtime-key curl call will be documented in a later implementation phase.

Follow-up:

- In the private-RPC phase, verify both the positive runtime-key path and the negative user-session path.

---

### 2026-04-30 - Go release and testing strategy

Source:

- Title: Release History - The Go Programming Language
  URL: https://go.dev/doc/devel/release
  Publisher/owner: Go team
  Date checked: 2026-04-30
- Title: testing package - testing - Go Packages
  URL: https://pkg.go.dev/testing@go1.26.2
  Publisher/owner: Go team
  Date checked: 2026-04-30
- Title: Go Wiki: TableDrivenTests - The Go Programming Language
  URL: https://go.dev/wiki/TableDrivenTests
  Publisher/owner: Go team
  Date checked: 2026-04-30

Relevant finding:

- The latest stable Go release found in the official release history is Go 1.26.2, released 2026-04-07.
- `go test` and the standard `testing` package are the official baseline for package tests.
- Table-driven tests are the preferred Go style when one test body needs to exercise multiple JSON, validation, or authorization cases.
- The `heroiclabs/nakama-pluginbuilder:3.37.0` image reports Go 1.25.5 during the Phase 2 Docker build, so that is the compatibility target for this repository.

Decision:

- Use Go 1.25.5 for the module `go` directive and plugin-builder compatibility.
- Use Go 1.25.5 for local tooling if a matching Go install is available.
- Use the standard library `testing` package for unit tests.
- Prefer table-driven tests for payload parsing, merge behavior, config validation, and authorization decisions.

Risk / uncertainty:

- Go 1.26.2 is newer than the plugin-builder runtime target, but it is not selected here because the Nakama 3.37.0 plugin-builder image currently builds against Go 1.25.5.

Follow-up:

- Keep `go.mod` and future local tooling aligned with Go 1.25.5 unless a later Nakama/plugin-builder upgrade changes the compatibility target.
