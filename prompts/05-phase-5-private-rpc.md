# Prompt 05 - Phase 5 Private Server-to-Server RPC

```text
Read AGENTS.md, docs/02-rpc-contracts.md, docs/03-verification-plan.md, docs/05-milestones.md, and docs/06-research-log.md.

Current target phase: Phase 5 - Private server-to-server RPC.

Work only on Phase 5.

Prompt/doc synchronization rule:
- Before editing, check whether Phase 0 or earlier phase decisions changed assumptions in this prompt or downstream docs/prompts.
- Update only affected docs/prompts.
- Do not rewrite unrelated prompt files.
- Keep this prompt aligned with AGENTS.md, docs/05-milestones.md, docs/06-research-log.md, and README.md.

Goal:
Implement private_health_check RPC that succeeds only through server-to-server/runtime-key invocation and fails for normal user-session calls.

Before editing:
- restate the private RPC contract;
- cite the documented server-to-server decision from docs/06-research-log.md;
- list pure authorization helper tests;
- list positive and negative integration checks.

TDD requirement:
Write unit tests for authorization decision logic, such as:
- empty user ID allowed;
- non-empty user ID rejected.

Implementation requirements:
- Register private_health_check in Nakama runtime.
- Reject calls with normal authenticated user context.
- Succeed for runtime HTTP key/server-to-server path.
- Return empty or minimal success response.
- Add positive and negative scripts/curl examples.
- Update README and docs if behavior changed.

Verification:
Run:
- go test ./...
- docker compose up --build, if needed
- private RPC call with runtime/server key succeeds
- private RPC call with normal user session fails

Constraints:
- Do not add admin UI.
- Do not add custom auth system.
- Do not expose private RPC to players.
- Keep local-only keys clearly marked.

After completion, report:
- files changed;
- tests added;
- commands run and output summary;
- integration result;
- blockers if any;
- whether Phase 5 is complete.
```
