# Technical Design

## Design objective

Build a minimal Nakama backend that is easy to run, verify, and review.

The final design should favor:

- small surface area;
- explicit version pinning;
- idiomatic Go;
- testable pure functions;
- clear RPC contracts;
- simple Docker-based local development;
- practical reviewer instructions.

## Technology baseline

The project should use:

- Nakama server;
- PostgreSQL;
- Go runtime module compiled as a plugin/shared object;
- Docker Compose for local orchestration;
- Go standard `testing` package for unit tests;
- shell scripts or documented curl examples for integration verification.

Phase 0 baseline selected from official docs and Phase 2 build validation:

- Nakama: `3.37.0`
- `github.com/heroiclabs/nakama-common/runtime`: `v1.44.2`
- `heroiclabs/nakama-pluginbuilder`: `3.37.0`
- Local Go tooling: `go1.25.5`

The Docker build path remains the source of truth for the runtime plugin build, because Nakama Go plugins must match the server binary's Go/dependency ABI. The selected plugin-builder image currently uses Go `1.25.5`, so `go.mod` should target that version.

Do not assume the newest local Go version is compatible with the selected Nakama plugin build. Nakama, plugin-builder, Go, and `nakama-common` compatibility must be confirmed from official documentation.

## Proposed final file layout

```text
.
├── AGENTS.md
├── README.md
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── local.yml
├── main.go
├── config/
│   └── game_config.json
├── scripts/
│   ├── auth-device.sh
│   ├── rpc-update-metadata.sh
│   ├── rpc-get-game-config.sh
│   └── rpc-private-health.sh
└── docs/
    ├── 00-assignment-brief.md
    ├── 01-technical-design.md
    ├── 02-rpc-contracts.md
    ├── 03-verification-plan.md
    ├── 04-agent-workflow.md
    ├── 05-milestones.md
    └── 06-research-log.md
```

## Nakama runtime approach

Use a Go runtime module with an `InitModule` entry point.

Expected responsibilities of `InitModule`:

- log module startup;
- register `update_user_metadata` RPC;
- register `get_game_config` RPC;
- register `private_health_check` RPC;
- load or prepare game config if needed.

Keep runtime initialization deterministic and small.

Avoid background goroutines unless there is a clear need. This assignment does not require them.

## Docker approach

Use Docker Compose to run:

- PostgreSQL;
- Nakama server;
- the Go plugin build process as needed.

Preferred approach:

- multi-stage Docker build if suitable;
- plugin-builder image compatible with selected Nakama version;
- Nakama image version pinned;
- local config file mounted or copied predictably.

Document all ports in README.

Typical Nakama local ports may include:

- API/HTTP port;
- gRPC port;
- console port.

Exact ports must be confirmed during Phase 0 research.

## PostgreSQL approach

Use PostgreSQL only as the Nakama database.

No custom application tables are required for this assignment unless implementation proves otherwise.

Use local development credentials only and document that they are not production secrets.

## RPC implementation approach

### `update_user_metadata`

Core logic:

1. Read caller user ID from Nakama runtime context.
2. Reject if missing.
3. Parse payload as JSON object.
4. Reject invalid JSON or unsupported shape.
5. Prevent payload from overriding target user ID.
6. Update the caller's Nakama account metadata.
7. Return success JSON.

Recommended pure functions:

```go
parseMetadataPayload(payload string) (map[string]any, error)
mergeMetadata(existing map[string]any, incoming map[string]any) map[string]any
```

Merge behavior is defined in `docs/02-rpc-contracts.md` as a shallow merge.

### `get_game_config`

Core logic:

1. Load config from `config/game_config.json`.
2. Treat the RPC as public; no user session is required.
3. Validate required fields.
4. Return JSON.

Recommended pure functions:

```go
loadGameConfig(path string) (GameConfig, error)
validateGameConfig(cfg GameConfig) error
```

Static config in `config/game_config.json` is the chosen source of truth for this repository. The runtime should load and validate that file, then return the JSON contents.

### `private_health_check`

Core logic:

1. Inspect runtime context.
2. Reject if a normal user ID is present.
3. Succeed for server-to-server/runtime-key invocation.
4. Return empty body or minimal success JSON depending on Nakama RPC conventions.

Recommended pure function:

```go
isPrivateRPCAllowed(userID string) bool
```

## Error handling strategy

Errors should be:

- explicit;
- useful for reviewers;
- not verbose with internal secrets;
- mapped to appropriate Nakama runtime errors.

Required error cases:

- unauthenticated metadata update;
- invalid JSON metadata payload;
- non-object metadata payload if object is required;
- metadata update failure;
- config loading/validation failure;
- user-session call to private RPC.

## Testing strategy

Use TDD-oriented development.

Unit tests should focus on pure logic:

- JSON payload parsing;
- config validation;
- metadata merge behavior;
- private RPC authorization decision.

Integration tests should be documented as curl/script flows against the running Docker stack.

Do not attempt to fully mock Nakama internals unless it is simple and clearly useful.

## Design constraints

- Keep the solution small.
- Prefer standard library.
- Avoid global mutable state.
- Avoid unnecessary abstractions.
- Avoid adding an application database schema.
- Avoid extra services.
- Avoid feature creep.
