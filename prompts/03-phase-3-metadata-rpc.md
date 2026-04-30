# Prompt 03 - Phase 3 Authenticated Metadata RPC

```text
Read AGENTS.md, docs/02-rpc-contracts.md, docs/03-verification-plan.md, docs/05-milestones.md, and docs/06-research-log.md.

Current target phase: Phase 3 - Authenticated metadata update RPC.

Work only on Phase 3.

Prompt/doc synchronization rule:
- Before editing, check whether Phase 0 or earlier phase decisions changed assumptions in this prompt or downstream docs/prompts.
- Update only affected docs/prompts.
- Do not rewrite unrelated prompt files.
- Keep this prompt aligned with AGENTS.md, docs/05-milestones.md, docs/06-research-log.md, and README.md.

Goal:
Implement update_user_metadata RPC.

Before editing:
- restate the RPC contract;
- confirm metadata merge/replace strategy and baseline assumptions from docs/06-research-log.md;
- list pure helper functions to test;
- list table-driven test cases;
- list integration verification steps.

TDD requirement:
Write unit tests before or alongside implementation for pure logic, including:
- valid JSON object accepted;
- invalid JSON rejected;
- non-object JSON rejected if object-only contract is chosen;
- metadata merge behavior deterministic;
- payload cannot override target user ID.

Implementation requirements:
- Register update_user_metadata in Nakama runtime.
- Require authenticated user context.
- Update only caller's own account metadata.
- Return clear JSON success response.
- Return useful errors for invalid input and unauthenticated calls.
- Add or update scripts/curl examples.
- Update README and docs if behavior changed.

Verification:
Run:
- go test ./...
- docker compose up --build, if needed
- authenticated RPC success flow
- unauthenticated failure flow
- invalid JSON failure flow

Constraints:
- Do not implement unrelated RPCs.
- Do not add custom auth outside Nakama.
- Prefer Go standard library.
- Keep handler small by extracting pure functions.

After completion, report:
- files changed;
- tests added;
- commands run and output summary;
- integration result;
- blockers if any;
- whether Phase 3 is complete.
```
