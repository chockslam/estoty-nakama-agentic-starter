# Milestones

## Purpose

This document defines the project phases required to complete the Estoty Nakama Go assignment.

Each phase must be completed independently and must satisfy its Definition of Done before the next phase begins.

This project uses a TDD-oriented workflow:

1. Research current official guidance.
2. Define expected behavior.
3. Write tests or executable verification steps before implementation where practical.
4. Implement the smallest code needed.
5. Run verification.
6. Update documentation.

The goal is not to create a large framework. The goal is to produce a small, correct, runnable, reviewable Nakama + PostgreSQL + Docker project.

---

## Phase 0 - Research and version baseline

### Goal

Establish the technical baseline before writing implementation code.

The agent must research current official documentation for:

- Nakama Go runtime;
- Nakama Docker/PostgreSQL setup;
- Nakama RPC registration;
- Nakama server-to-server/private RPC behavior;
- Nakama and `nakama-common` version compatibility;
- current Go version and relevant Go testing/style best practices.

### Required outputs

- Completed entries in `docs/06-research-log.md`.
- Confirmed Nakama image version.
- Confirmed `nakama-common` version.
- Confirmed Go version/toolchain strategy.
- Confirmed Docker/plugin-builder strategy.
- Confirmed TDD strategy.

### TDD / verification-first requirement

No application code should be written in this phase.

The agent should identify what will be tested in later phases and update `docs/03-verification-plan.md` if research changes the expected verification method.

### Definition of Done

- Official Nakama docs have been checked.
- Official Go docs have been checked.
- Version choices are documented.
- Any uncertainty is explicitly recorded.
- No implementation code has been written unless needed for discovery.
- The agent has not assumed that the newest local Go version is automatically compatible with the Nakama plugin build.

---

## Phase 1 - Project skeleton and documentation baseline

### Goal

Create the repository structure and documentation required for agentic coding.

### Required outputs

- `AGENTS.md`
- `README.md`
- `docs/00-assignment-brief.md`
- `docs/01-technical-design.md`
- `docs/02-rpc-contracts.md`
- `docs/03-verification-plan.md`
- `docs/04-agent-workflow.md`
- `docs/05-milestones.md`
- `docs/06-research-log.md`

### TDD / verification-first requirement

Validate that all required documentation exists before implementation begins.

### Definition of Done

- Assignment is restated in engineering language.
- Scope and out-of-scope items are clear.
- RPC contracts are named and documented.
- Verification plan includes success and failure cases.
- README explains intended local workflow, even if commands are placeholders.
- Agent workflow instructs future agents to work phase-by-phase.
- No unnecessary architecture has been introduced.

---

## Phase 2 - Docker/Nakama/PostgreSQL runtime skeleton

### Goal

Make the project boot successfully with Docker Compose and load a minimal Go runtime module.

### Required behavior

- PostgreSQL starts.
- Nakama starts.
- Go runtime module builds.
- Nakama loads the Go runtime module.
- A minimal runtime log message confirms module loading.

### Required outputs

Expected files after this phase:

- `Dockerfile`
- `docker-compose.yml`
- `local.yml`
- `go.mod`
- `go.sum`
- `main.go`
- updated `README.md`
- updated `docs/06-research-log.md`

### TDD / verification-first requirement

Before implementation, define the expected verification commands in `docs/03-verification-plan.md`.

At minimum, the verification must include:

```bash
docker compose up --build
```

And a log check proving the Go runtime loaded.

### Definition of Done

- `docker compose up --build` succeeds.
- Nakama connects to PostgreSQL.
- Nakama API/console ports are documented.
- Runtime module is loaded.
- No RPC business logic is required yet.
- README contains exact startup commands.
- Version pinning is documented.
- Generated/build artifacts are ignored where appropriate.

---

## Phase 3 - Authenticated metadata update RPC

### Goal

Implement the authenticated RPC that updates the caller's account metadata.

### Required behavior

RPC name:

```text
update_user_metadata
```

Rules:

- Must require an authenticated Nakama user context.
- Must update only the caller's own account.
- Must accept arbitrary JSON object metadata from the caller.
- Must reject invalid JSON.
- Must reject non-object JSON if the contract requires an object.
- Must not allow user ID override.
- Must return a clear JSON success response.

### Required outputs

- RPC handler implementation.
- Pure helper functions where useful.
- Unit tests.
- Script or curl instructions.
- Updated README.
- Updated RPC contract docs if behavior changed.

### TDD requirement

Before implementing the RPC handler, create tests for pure logic where possible:

- valid JSON object accepted;
- invalid JSON rejected;
- array/string/number payload rejected if only objects are supported;
- metadata merge behavior is deterministic;
- user ID cannot be supplied by payload.

If direct Nakama runtime testing is difficult, isolate pure functions such as:

```go
parseMetadataPayload(payload string) (map[string]any, error)
mergeMetadata(existing map[string]any, incoming map[string]any) map[string]any
```

Use table-driven tests.

### Integration verification

Add a script or documented curl flow proving:

- authenticated user can update own metadata;
- unauthenticated call fails;
- invalid JSON fails cleanly.

### Definition of Done

- Unit tests pass.
- Integration verification passes.
- README includes exact call examples.
- RPC contract documentation is updated.
- Error responses are understandable.
- No unrelated RPCs or features are added.

---

## Phase 4 - Game configuration RPC

### Goal

Implement an RPC that returns game configuration JSON.

### Required behavior

RPC name:

```text
get_game_config
```

Returned JSON must include:

```json
{
  "welcomeMessage": "Welcome to the game!",
  "xpRate": 1.0,
  "rarityOptions": ["common", "rare", "epic", "legendary"]
}
```

The exact values may differ, but the required field types must be preserved.

### Required outputs

- RPC handler implementation.
- Config source, preferably `config/game_config.json`, unless embedded config is justified.
- Config validation helper.
- Unit tests.
- Script or curl instructions.
- Updated README.

### TDD requirement

Before implementation, create tests for:

- config loads successfully;
- missing required fields fail validation;
- invalid `xpRate` fails validation;
- empty rarity list fails validation;
- returned JSON is valid.

Prefer a small config loader/validator function that can be tested without Docker.

### Integration verification

Add a script or documented curl flow proving:

- RPC returns valid JSON without requiring a user session;
- required fields are present;
- field types are correct.

### Definition of Done

- Unit tests pass.
- Integration verification passes.
- Config source is documented.
- README includes exact call example.
- No database dependency is introduced for static config unless explicitly justified.

---

## Phase 5 - Private server-to-server RPC

### Goal

Implement the private RPC that can be called only through the server-to-server/runtime-key path.

### Required behavior

RPC name:

```text
private_health_check
```

Rules:

- User-authenticated calls must fail.
- Server-to-server/runtime-key calls must succeed.
- No meaningful response body is required.
- The purpose is to prove correct authorization boundary, not business logic.

### Required outputs

- RPC handler implementation.
- Authorization helper if useful.
- Unit tests.
- Positive and negative integration scripts/curl instructions.
- Updated README.

### TDD requirement

Before implementation, isolate the authorization decision if practical:

```go
isPrivateRPCAllowed(userID string) bool
```

Expected behavior:

- empty user ID: allowed;
- non-empty user ID: rejected.

### Integration verification

Add a script or documented curl flow proving:

- normal authenticated session cannot call the private RPC successfully;
- runtime HTTP key/server-to-server call succeeds.

### Definition of Done

- Unit tests pass.
- Integration verification passes.
- Private RPC cannot be called as a normal player.
- README clearly distinguishes user-session calls from server-to-server calls.
- Local-only keys are documented as local-only.

---

## Phase 6 - Full verification and reviewer polish

### Goal

Prepare the repository for Estoty review.

### Required behavior

A reviewer must be able to:

1. Clone the repository.
2. Start the stack.
3. Authenticate a test user.
4. Call metadata update RPC.
5. Call game config RPC.
6. Call private RPC using server-to-server key.
7. See negative tests documented or scripted.

### Required outputs

- Final README.
- Final verification evidence.
- Clean scripts.
- Clean `.gitignore`.
- No irrelevant files.
- Final assignment mapping section.

### Required commands

At minimum:

```bash
go test ./...
docker compose up --build
```

And documented curl/script calls for all RPCs.

### Definition of Done

- `go test ./...` passes.
- Docker stack starts from a clean checkout.
- All three RPCs work.
- Required negative tests are documented and verified.
- README is complete and concise.
- Assignment requirements are explicitly mapped to implemented files.
- Repo contains no irrelevant generated files, secrets, or abandoned experiments.
- Final summary explains:
  - what was built;
  - how to run it;
  - how to call each RPC;
  - what trade-offs were made.

---

## Phase 7 - Optional improvements only if time remains

### Allowed optional improvements

Only after all required phases are complete:

- GitHub Actions running `go test ./...`.
- Small `Makefile` for common commands.
- Basic lint command.
- More polished shell scripts.
- Additional config validation tests.
- Short architecture diagram in README.

### Not allowed unless explicitly justified

- Kubernetes.
- Cloud deployment.
- Custom admin UI.
- Custom web server.
- Complex database migrations.
- Extra gameplay systems.
- Matchmaking.
- Leaderboards.
- Real-time multiplayer simulation.
- Authentication systems outside Nakama.
- Large third-party dependency stack.

### Definition of Done

- Optional improvements do not obscure the core assignment.
- README remains simple.
- Reviewer can still understand the project in under 5 minutes.
