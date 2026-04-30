# Estoty Nakama Go Assignment

Runnable Nakama backend for the Estoty Go/Golang take-home assignment.

Implemented RPCs:

- `update_user_metadata` - authenticated metadata merge for the caller's own account.
- `get_game_config` - public game config JSON loaded from [`config/game_config.json`](config/game_config.json).
- `private_health_check` - server-to-server RPC protected by Nakama's runtime HTTP key.

## Design Decisions

- The runtime module is written in Go because the assignment is for a Go Developer role.
- Nakama, `nakama-common`, and plugin-builder versions are pinned together to avoid Go plugin compatibility issues.
- RPC handlers keep Nakama-specific code thin and delegate validation, merge, and authorization decisions to small testable functions.
- `private_health_check` rejects user-session calls and succeeds only through the server-to-server/runtime HTTP key path.
- The project avoids unrelated gameplay systems so the assignment stays focused and reviewable.

## Baseline

Pinned versions from official docs and Docker validation:

- Nakama `3.37.0`
- `github.com/heroiclabs/nakama-common/runtime` `v1.44.2`
- `heroiclabs/nakama-pluginbuilder` `3.37.0`
- Go toolchain target `1.25.5`

Local-only values in the default Docker setup:

- PostgreSQL password: `localdb`
- Nakama server key: `defaultkey`
- Nakama runtime HTTP key: `NAKAMA_RUNTIME_HTTP_KEY` (`NAKAMA_HTTP_KEY` is also accepted by the review scripts): `defaulthttpkey`
- Nakama session encryption key: `defaultencryptionkey`
- Nakama refresh encryption key: `defaultrefreshencryptionkey`

Do not reuse those values outside local review.

## Quick Start

```bash
docker compose up --build
```

Ports:

- HTTP API: `http://localhost:7350`
- Console HTTP: `http://localhost:7351`
- gRPC API: `localhost:7349`
- Console gRPC: `localhost:7348`

## Verification

Unit tests:

```bash
go test ./...
```

Pinned-toolchain fallback:

```bash
docker run --rm -v "$PWD":/backend -w /backend heroiclabs/nakama-pluginbuilder:3.37.0 test ./...
```

Metadata RPC:

```bash
SESSION_TOKEN="$(./scripts/auth-device.sh)"
SESSION_TOKEN="$SESSION_TOKEN" ./scripts/rpc-update-metadata.sh
```

Expected success:

```json
{"success":true,"metadata":{...}}
```

Merge semantics: metadata updates use a deterministic shallow merge. Incoming top-level keys overwrite existing top-level keys, nested object deep-merge is out of scope, and the caller cannot choose the target user/account.

Negative metadata examples:

```bash
curl -i -X POST 'http://127.0.0.1:7350/v2/rpc/update_user_metadata?unwrap=true' \
  -H 'Content-Type: application/json' \
  --data '{"favoriteHero":"warrior"}'
```

Expected unauthenticated failure:

```text
HTTP 401
Auth token or HTTP key required
```

```bash
SESSION_TOKEN="$(./scripts/auth-device.sh)"
curl -i -X POST 'http://127.0.0.1:7350/v2/rpc/update_user_metadata?unwrap=true' \
  -H "Authorization: Bearer ${SESSION_TOKEN}" \
  -H 'Content-Type: application/json' \
  --data 'not-json'
```

Expected invalid JSON failure:

```text
HTTP 400
metadata payload must be valid JSON
```

Game config RPC:

`get_game_config` is intentionally not private and may be called as a normal RPC. The helper script uses the local runtime HTTP key only for simple reviewer verification; the private RPC is `private_health_check`.

```bash
./scripts/rpc-get-game-config.sh
```

Expected success:

```json
{"welcomeMessage":"Welcome to the game!","xpRate":1.5,"rarityOptions":["common","rare","epic","legendary"]}
```

Private RPC:

Server-to-server success:

```bash
./scripts/rpc-private-health.sh
```

User-session failure:

```bash
MODE=user SESSION_TOKEN="$(./scripts/auth-device.sh)" ./scripts/rpc-private-health.sh
```

Expected success:

```json
{"success":true}
```

Expected user-session failure:

```text
HTTP 403
private_health_check is only callable via server-to-server/runtime HTTP key
```

## Reviewer Map

- [`main.go`](main.go) registers all runtime RPCs.
- [`metadata_rpc.go`](metadata_rpc.go) and [`metadata_rpc_test.go`](metadata_rpc_test.go) implement and test `update_user_metadata`.
- [`config_rpc.go`](config_rpc.go) and [`config_rpc_test.go`](config_rpc_test.go) implement and test `get_game_config`.
- [`private_rpc.go`](private_rpc.go) and [`private_rpc_test.go`](private_rpc_test.go) implement and test `private_health_check`.
- [`Dockerfile`](Dockerfile), [`docker-compose.yml`](docker-compose.yml), and [`local.yml`](local.yml) boot Nakama + PostgreSQL.
- [`scripts/`](scripts) contains the exact review commands.
- [`docs/`](docs) contains the assignment brief, contracts, verification plan, milestones, and research log.

## Notes

The Docker build path is the source of truth for the Nakama Go plugin ABI. The host Go toolchain is only for local tooling and tests.
