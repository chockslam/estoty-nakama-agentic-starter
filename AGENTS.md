# AGENTS.md

## Repository purpose

This repository is for a take-home assignment for an Estoty Go/Golang Developer role.

The expected final project is a small, runnable Nakama backend using:

- Go runtime module
- Nakama
- PostgreSQL
- Docker / Docker Compose

The assignment requires three RPC methods:

1. `update_user_metadata` - authenticated RPC that updates the caller's account metadata by adding arbitrary JSON information from the caller.
2. `get_game_config` - RPC that returns game configuration JSON.
3. `private_health_check` - private server-to-server RPC callable only through the server/runtime key path.

This is a backend engineering test. Prioritize correctness, simple architecture, reliable local setup, clear docs, tests, and verifiable commands.

## Milestone discipline

This repository is developed phase-by-phase.

Agents must follow:

```text
docs/05-milestones.md
```

Do not skip phases.

Do not start implementation before Phase 0 research and Phase 1 documentation are complete.

Each phase has a Definition of Done. A phase is not complete until every Definition of Done item is satisfied or explicitly marked as not possible with a reason.

When asked to work on the project, first identify the current phase, then work only on the next incomplete phase.

## Research requirement

Before making version-sensitive technical choices, research current official documentation.
If a decision depends on versions, runtime behavior, or platform quirks, confirm it with the official docs first and record the source in `docs/06-research-log.md` before editing implementation code.

Required official sources:

- Heroic Labs Nakama documentation.
- Go official documentation.
- Docker documentation if Docker behavior is unclear.
- PostgreSQL documentation if database behavior is unclear.

Prefer official documentation over blog posts.

Record decisions in:

```text
docs/06-research-log.md
```

The research log must include:

- source title;
- URL;
- date checked;
- decision influenced by the source;
- uncertainty or follow-up if applicable.

Do not claim “latest best practices” without checking current sources.

## TDD requirement

Use a TDD-oriented workflow.

For each RPC:

1. Define the contract.
2. Write unit tests for pure logic before implementation where practical.
3. Implement the minimal code to pass tests.
4. Add integration verification using Docker/Nakama/curl.
5. Update README and RPC docs.

Prefer table-driven tests for payload parsing, validation, config loading, merge behavior, and authorization decision logic.

Use Go's standard `testing` package unless an additional dependency is clearly justified.

Do not make Nakama runtime handlers impossible to test. Extract pure functions where useful.

## Technical direction

Use a Nakama Go runtime module rather than Lua or TypeScript because the role is for Go development.

Nakama Go runtime code must expose `InitModule`, where RPCs and hooks are registered.

Build the Go runtime module as a Nakama plugin/shared object inside Docker. Use compatible Nakama, plugin-builder, Go, and `nakama-common` versions.

Use PostgreSQL as the Nakama database.

Do not introduce a separate HTTP API server unless there is a strong reason. Prefer Nakama RPCs.

## Expected project shape

Target final structure:

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

Adjust only if implementation proves this layout wrong.

## RPC requirements

### `update_user_metadata`

Purpose: update the authenticated caller's Nakama account metadata.

Rules:

- Must require authenticated user context.
- Must update only the caller's own account.
- Must reject invalid JSON.
- Must reject non-object JSON if the selected contract requires an object.
- Must not allow the caller to specify another user ID.
- Should return a clear success payload and/or updated metadata.

### `get_game_config`

Purpose: return game configuration JSON.

Required fields:

- `welcomeMessage`: string
- `xpRate`: float
- `rarityOptions`: list of strings

Rules:

- Keep the config small and readable.
- Prefer loading from `config/game_config.json` or embedding a simple typed config.
- Return valid JSON.

### `private_health_check`

Purpose: prove server-to-server/private RPC protection.

Rules:

- Must reject calls that include a user ID in the Nakama runtime context.
- Must succeed only for server-to-server/runtime-key style invocation.
- No meaningful response body is required.

## Security rules

Do not expose private RPCs to normal authenticated users.

Do not commit production secrets.

Local default Nakama keys may be used only for local testing and must be clearly documented as local-only.

Do not place runtime HTTP key, server key, database passwords, or other secrets in scripts without making it clear they are test defaults.

## Error handling

Return clear Nakama runtime errors for:

- unauthenticated user calling authenticated RPC;
- authenticated user calling private RPC;
- invalid JSON payload;
- internal account update failure;
- config loading/parsing failure.

Prefer explicit error messages that help the reviewer understand what failed.

## Verification requirements

Before claiming completion, run and capture results for:

```bash
go test ./...
docker compose up --build
```

Then verify:

- Nakama starts.
- PostgreSQL starts.
- Go runtime module is loaded.
- User authentication works.
- Metadata RPC works for authenticated user.
- Metadata RPC fails without authentication.
- Game config RPC returns expected JSON.
- Private RPC fails from user-authenticated context.
- Private RPC succeeds through server-to-server/runtime HTTP key path.
- Invalid JSON returns a clean error.

If any command cannot be run, state exactly why.

## Coding rules

- Keep the implementation small.
- Prefer the Go standard library.
- Avoid unnecessary packages.
- Avoid background goroutines.
- Avoid global mutable state unless read-only config is loaded safely.
- Keep RPC names stable once documented.
- Update docs when implementation changes behavior.
- Do not overengineer this into a framework.

## Do-not-overbuild rules

Do not add these unless explicitly required after all core assignment requirements are complete:

- Kubernetes.
- Cloud deployment.
- Custom admin UI.
- Custom web server.
- Matchmaking.
- Leaderboards.
- Real-time multiplayer simulation.
- Payment systems.
- External auth providers.
- Complex migrations.
- Large dependency stack.

## Agent workflow

For each task:

1. Read `AGENTS.md`.
2. Read relevant docs in `docs/`.
3. Identify the current milestone.
4. Produce a short plan before editing.
5. Make minimal changes.
6. Run relevant verification.
7. Update README/docs if behavior changed.
8. Summarize changed files and command results.

Never claim that a feature works without either command output or a clear explanation that it was not possible to run locally.
