# Prompt 01 - Phase 0 Research and Version Baseline (Historical)

Use this only if Nakama, Go, Docker, or dependency versions need to be revalidated.

```text
Read AGENTS.md, docs/05-milestones.md, and docs/06-research-log.md.

Current target phase: Phase 0 - Research and version baseline.

Work only on Phase 0.

Prompt/doc synchronization rule:
- Before editing, check whether prior phase decisions changed assumptions in current or downstream docs/prompts.
- Update only affected docs/prompts.
- Do not rewrite unrelated prompt files.
- Keep this prompt aligned with AGENTS.md, docs/05-milestones.md, docs/06-research-log.md, and README.md.

Goal:
Research current official documentation and establish the technical baseline for the Nakama + PostgreSQL + Docker + Go runtime project.

Research official sources for:
- Nakama Go runtime module setup;
- Nakama Docker/PostgreSQL local setup;
- Nakama RPC registration from Go runtime;
- Nakama server-to-server/private RPC behavior;
- Nakama, nakama-common, plugin-builder, and Go version compatibility;
- current Go testing/style best practices relevant to this project.

Required output:
Update docs/06-research-log.md with structured entries.

For each entry include:
- source title;
- URL;
- publisher/owner;
- date checked;
- relevant finding;
- decision;
- risk/uncertainty;
- follow-up.

Also update docs/01-technical-design.md and docs/03-verification-plan.md if research changes the expected implementation or verification approach.

Constraints:
- Prefer official docs over blog posts.
- Do not implement project code.
- Do not assume the newest local Go version is compatible with Nakama plugin building.
- Do not pin versions without documenting why.

Definition of Done:
- Official Nakama docs checked.
- Official Go docs checked.
- Version choices documented.
- Plugin-builder strategy documented.
- Server-to-server/private RPC verification strategy documented.
- Open uncertainties listed.

After completion, report:
- selected versions;
- changed files;
- key decisions;
- unresolved risks;
- whether Phase 0 is complete.
```
