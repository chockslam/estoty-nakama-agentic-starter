# Prompt 02 - Phase 2 Docker/Nakama/PostgreSQL Runtime Skeleton

```text
Read AGENTS.md, docs/05-milestones.md, docs/06-research-log.md, and docs/03-verification-plan.md.

Current target phase: Phase 2 - Docker/Nakama/PostgreSQL runtime skeleton.

Work only on Phase 2.

Goal:
Make the project boot with Docker Compose and load a minimal Nakama Go runtime module.

Before editing:
- summarize selected versions from docs/06-research-log.md;
- confirm the expected Docker/plugin-builder strategy;
- list files you will create or modify;
- list verification commands.

Required implementation:
- Dockerfile for building the Go runtime module using compatible Nakama/plugin-builder approach.
- docker-compose.yml for PostgreSQL + Nakama.
- local.yml or equivalent Nakama local config.
- go.mod and go.sum.
- main.go with minimal InitModule that logs module startup.
- README updates with startup instructions.

Do not implement RPC business logic yet.

Verification:
Run or prepare to run:
- docker compose up --build
- docker compose logs check proving Nakama started and the Go runtime loaded

Constraints:
- Keep implementation minimal.
- Use compatible versions from Phase 0.
- Do not add custom web server.
- Do not add gameplay features.
- Do not commit production secrets.

Definition of Done:
- Docker stack starts.
- PostgreSQL starts.
- Nakama starts.
- Go runtime module loads.
- README contains exact startup commands.
- Any local-only keys are documented as local-only.

After completion, report:
- files changed;
- commands run;
- relevant output summary;
- blockers if any;
- whether Phase 2 is complete.
```
