# Agent Workflow

## Purpose

This document defines how to use agentic coding systems such as Codex on this repository.

The objective is to keep the agent bounded, verifiable, and aligned with the assignment.

## Core rules

- One phase at a time.
- One branch or worktree per implementation task.
- Research before version-sensitive decisions.
- Tests before implementation where practical.
- Minimal code changes.
- Verify before claiming completion.
- Update docs when behavior changes.

## Recommended branch model

```text
main
phase-0-research
phase-1-docs
phase-2-runtime-skeleton
phase-3-metadata-rpc
phase-4-config-rpc
phase-5-private-rpc
phase-6-final-polish
```

Each branch should be small and reviewable.

## Starting a new agent task

Use this sequence:

1. Tell the agent to read `AGENTS.md`.
2. Tell the agent to read `docs/05-milestones.md`.
3. Specify the target phase.
4. Ask for a short implementation plan.
5. Require tests or verification plan before code.
6. Require a final summary with command output.

## Agent task template

```text
Read AGENTS.md and docs/05-milestones.md.

Current target phase: <PHASE NAME>.

Work only on this phase.

Before editing:
- summarize the phase goal;
- identify required outputs;
- identify Definition of Done;
- list tests or verification commands you will run.

Then implement the smallest changes needed.

After editing:
- run required tests/verification where possible;
- update README/docs if behavior changed;
- summarize files changed;
- paste relevant command output;
- list unresolved risks or blockers.
```

## Good agent behavior

A good agent should:

- stop to research official docs before pinning versions;
- add pure functions to make logic testable;
- keep Nakama-specific code thin;
- use table-driven tests;
- avoid unnecessary libraries;
- leave clear scripts/curl examples;
- explain failures honestly.

## Bad agent behavior

Reject or correct behavior where the agent:

- implements features before Phase 0 research;
- uses Lua/TypeScript runtime without justification;
- assumes latest Go automatically works with Nakama plugin build;
- adds a custom API server;
- adds game systems not requested;
- hardcodes production-looking secrets;
- claims Docker works without running it;
- creates a large architecture unrelated to the assignment.

## Suggested task order

### Task 0 - Project-control docs

Create this documentation baseline.

### Task 1 - Research and version baseline

Complete `docs/06-research-log.md` with current official-source findings.

### Task 2 - Runtime skeleton

Make Docker Compose boot Nakama and PostgreSQL and load a minimal Go runtime module.

### Task 3 - Metadata RPC

Implement and verify `update_user_metadata`.

### Task 4 - Config RPC

Implement and verify `get_game_config`.

### Task 5 - Private RPC

Implement and verify `private_health_check`.

### Task 6 - Final polish

Run full verification and prepare reviewer-facing README.

## Agent output expectations

Every agent completion message must include:

```text
Phase:
Files changed:
Commands run:
Verification result:
Open risks/blockers:
Next recommended phase:
```
