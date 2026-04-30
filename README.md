# Estoty Nakama Go Assignment - Agentic Starter

This repository starter contains the documentation, prompts, milestones, and workflow controls for building the Estoty Go/Golang Developer take-home assignment using agentic coding systems such as Codex.

It is **not yet the finished Nakama implementation**. It is the project-control layer that should be committed before implementation begins.

## Assignment target

Build a Docker + PostgreSQL-based Nakama project using Go runtime code.

Final project must implement:

1. `update_user_metadata` - authenticated RPC that updates the caller's account metadata by adding arbitrary JSON information from the caller.
2. `get_game_config` - RPC that returns a free-form JSON game configuration containing:
   - welcome message: string;
   - xp rate: float;
   - rarity options: list of strings.
3. `private_health_check` - private server-to-server RPC callable only with the server/runtime key, returning success status only.

## How to use this starter

1. Unzip this folder.
2. Initialize a Git repository.
3. Commit the starter docs.
4. Give your agent the prompt in:

```text
prompts/00-agentic-kickoff.md
```

5. Let the agent complete Phase 0 and Phase 1 first.
6. Only then start implementation phases.

Recommended first commands:

```bash
git init
git add .
git commit -m "Add agentic project-control docs"
```

## Documentation map

```text
AGENTS.md                         Agent operating instructions read by Codex-style tools.
docs/00-assignment-brief.md        Engineering restatement of the assignment.
docs/01-technical-design.md        Proposed architecture and implementation approach.
docs/02-rpc-contracts.md           RPC names, payloads, auth rules, errors, examples.
docs/03-verification-plan.md       Unit and integration verification checklist.
docs/04-agent-workflow.md          Agentic coding workflow and task sequencing.
docs/05-milestones.md              Phase plan with Definition of Done per phase.
docs/06-research-log.md            Required research log template.
prompts/00-agentic-kickoff.md      First prompt for Codex/agentic coding.
prompts/01-phase-0-research.md     Prompt for research/version baseline.
prompts/02-phase-2-runtime.md      Prompt for Docker/Nakama/PostgreSQL skeleton.
prompts/03-phase-3-metadata-rpc.md Prompt for authenticated metadata RPC.
prompts/04-phase-4-config-rpc.md   Prompt for config RPC.
prompts/05-phase-5-private-rpc.md  Prompt for private server-to-server RPC.
prompts/06-final-polish.md         Prompt for final verification and reviewer polish.
```

## Core development principle

This assignment should be built incrementally:

1. Research and pin compatible versions.
2. Create skeleton and docs.
3. Make Docker + Nakama + PostgreSQL boot.
4. Add one RPC at a time.
5. Use tests for pure logic.
6. Verify integration through Docker and curl.
7. Keep the README practical for the reviewer.

## Expected final verification

Before submission, the final repository should pass:

```bash
go test ./...
docker compose up --build
```

And the README should show how to verify:

- authenticated metadata update success;
- metadata update failure without auth;
- game config RPC success;
- private RPC failure from a user session;
- private RPC success through server-to-server/runtime key;
- invalid JSON clean failure.

## Current status

This starter is ready for agentic coding workflow setup. Implementation code should be added only after Phase 0 and Phase 1 are complete.
