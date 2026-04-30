# Prompt 06 - Phase 6 Final Verification and Reviewer Polish

```text
Read AGENTS.md, docs/03-verification-plan.md, docs/05-milestones.md, README.md, and the implementation files.

Current target phase: Phase 6 - Full verification and reviewer polish.

Work only on Phase 6.

Goal:
Prepare the repository for Estoty review.

Before editing:
- summarize implemented features;
- map assignment requirements to files;
- list verification commands;
- identify docs/scripts needing cleanup.

Required work:
- Make README practical for a reviewer who clones the repo.
- Include exact startup commands.
- Include exact RPC call examples or scripts.
- Include expected successful responses.
- Include negative test examples.
- Clearly mark local-only keys.
- Remove irrelevant generated files or abandoned experiments.
- Ensure .gitignore is correct.

Final verification:
Run:
- go test ./...
- docker compose up --build
- authenticated metadata RPC success
- metadata RPC unauthenticated failure
- invalid JSON failure
- game config RPC success
- private RPC user-session failure
- private RPC server-to-server success

Constraints:
- Do not add optional improvements until required verification passes.
- Do not obscure the assignment with extra systems.
- Keep README concise.

After completion, report:
- files changed;
- all commands run;
- pass/fail status for each verification item;
- final submission readiness;
- any known limitations.
```
