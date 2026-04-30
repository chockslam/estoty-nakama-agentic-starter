# Prompt 04 - Phase 4 Game Configuration RPC

```text
Read AGENTS.md, docs/02-rpc-contracts.md, docs/03-verification-plan.md, docs/05-milestones.md, and docs/06-research-log.md.

Current target phase: Phase 4 - Game configuration RPC.

Work only on Phase 4.

Prompt/doc synchronization rule:
- Before editing, check whether Phase 0 or earlier phase decisions changed assumptions in this prompt or downstream docs/prompts.
- Update only affected docs/prompts.
- Do not rewrite unrelated prompt files.
- Keep this prompt aligned with AGENTS.md, docs/05-milestones.md, docs/06-research-log.md, and README.md.

Goal:
Implement get_game_config RPC.

Before editing:
- restate the RPC contract;
- decide whether config is file-based or embedded, using docs/01-technical-design.md and baseline assumptions from docs/06-research-log.md;
- list pure config validation functions;
- list table-driven tests;
- list integration verification steps.

TDD requirement:
Write unit tests for:
- valid config accepted;
- missing welcomeMessage rejected;
- invalid xpRate rejected;
- empty rarityOptions rejected;
- invalid rarity option rejected;
- returned config marshals to valid JSON.

Implementation requirements:
- Register get_game_config in Nakama runtime.
- Return JSON containing:
  - welcomeMessage: string;
  - xpRate: float/number;
  - rarityOptions: list of strings.
- Validate config.
- Add or update scripts/curl examples.
- Update README and docs if behavior changed.

Verification:
Run:
- go test ./...
- docker compose up --build, if needed
- get_game_config RPC call
- verify response field names and types

Constraints:
- Do not introduce a database table for static config unless justified.
- Do not add unrelated gameplay/economy systems.
- Keep config small.

After completion, report:
- files changed;
- tests added;
- commands run and output summary;
- integration result;
- blockers if any;
- whether Phase 4 is complete.
```
