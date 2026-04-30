# Verification Plan

## Purpose

This document defines how agents and reviewers verify the project.

No phase is complete until its relevant verification steps pass or a blocker is clearly documented.

## Verification levels

The project uses two levels of verification:

1. Unit-level tests for pure Go logic.
2. Integration-level verification against Docker/Nakama/PostgreSQL using curl or scripts.

## Required final commands

Before final submission, run:

```bash
go test ./...
docker compose up --build
```

If either command cannot run, document the exact reason.

## Phase 0 verification

Research/version baseline is complete when:

- `docs/06-research-log.md` contains official-source entries for Nakama Go runtime.
- `docs/06-research-log.md` contains official-source entries for Nakama Docker/PostgreSQL setup.
- `docs/06-research-log.md` contains official-source entries for server-to-server/private RPC behavior.
- selected versions are recorded.
- unresolved uncertainties are documented.

## Phase 1 verification

Documentation baseline is complete when all required docs exist:

```bash
ls AGENTS.md README.md docs/00-assignment-brief.md docs/01-technical-design.md docs/02-rpc-contracts.md docs/03-verification-plan.md docs/04-agent-workflow.md docs/05-milestones.md docs/06-research-log.md
```

## Phase 2 verification - runtime skeleton

Expected command:

```bash
docker compose up --build
```

Expected checks:

- PostgreSQL container starts.
- Nakama container starts.
- Nakama connects to PostgreSQL.
- Go runtime module is built.
- Go runtime module is loaded by Nakama.
- Logs contain a clear module startup message.

Suggested log check placeholder:

```bash
docker compose logs nakama | grep -i "runtime"
```

Exact service name may differ and must be updated after implementation.

## Phase 3 verification - `update_user_metadata`

### Unit tests

Expected command:

```bash
go test ./...
```

Expected test cases:

- valid JSON object accepted;
- invalid JSON rejected;
- array/string/number payload rejected if object-only contract is chosen;
- shallow merge behavior deterministic;
- caller cannot supply target user ID;
- unauthenticated user rejected by authorization helper if extracted.

### Integration tests

Required successful flow:

1. Start Docker stack.
2. Authenticate a test user.
3. Call `update_user_metadata` with valid JSON.
4. Verify success response.
5. Optionally fetch account and verify metadata was updated.

Required negative flows:

- call without valid user session must fail;
- call with invalid JSON must fail cleanly;
- call with non-object JSON must fail if object-only contract is chosen.

Verification commands:

```bash
SESSION_TOKEN="$(./scripts/auth-device.sh)"
SESSION_TOKEN="$SESSION_TOKEN" ./scripts/rpc-update-metadata.sh
```

## Phase 4 verification - `get_game_config`

### Unit tests

Expected command:

```bash
go test ./...
```

Expected test cases:

- valid config accepted;
- missing `welcomeMessage` rejected;
- missing or invalid `xpRate` rejected;
- empty rarity list rejected;
- invalid rarity option rejected;
- config marshals to valid JSON.

### Integration tests

Required flow:

1. Start Docker stack.
2. Call `get_game_config` without authenticating first.
3. Verify returned JSON contains:
   - `welcomeMessage` string;
   - `xpRate` number;
   - `rarityOptions` string list.

Placeholder command:

```bash
./scripts/rpc-get-game-config.sh
```

## Phase 5 verification - `private_health_check`

### Unit tests

Expected command:

```bash
go test ./...
```

Expected test cases:

- empty user ID is allowed;
- non-empty user ID is rejected.

### Integration tests

Required positive flow:

1. Start Docker stack.
2. Call private RPC using runtime HTTP key / server-to-server path.
3. Verify successful status.

Required negative flow:

1. Authenticate a normal user.
2. Call private RPC using user session.
3. Verify failure.

Placeholder commands:

```bash
./scripts/rpc-private-health.sh
# documented user-session curl example
```

## Final reviewer verification checklist

Before submission, confirm:

- `go test ./...` passes.
- `docker compose up --build` starts the full stack.
- README startup instructions work from a clean checkout.
- Authenticated metadata RPC succeeds.
- Metadata RPC fails without auth.
- Invalid metadata JSON fails cleanly.
- Game config RPC returns expected JSON.
- Private RPC fails with normal user session.
- Private RPC succeeds with runtime/server key.
- No production secrets are committed.
- Local demo keys are clearly marked as local-only.
- Assignment requirements are mapped to implementation files.

## Evidence expectation for agents

When an agent reports completion, it must include:

- command run;
- relevant output summary;
- files changed;
- unresolved risks or skipped checks.

Do not claim success without command output or explicit blocker explanation.
