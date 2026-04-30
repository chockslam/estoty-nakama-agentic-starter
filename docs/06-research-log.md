# Research Log

## Purpose

This file records technical research that affects implementation decisions.

Agents must update this file when making decisions about:

- Go version;
- Nakama version;
- `nakama-common` version;
- Docker/plugin-builder strategy;
- RPC calling conventions;
- server-to-server/private RPC authorization;
- testing strategy;
- PostgreSQL/Nakama configuration.

Prefer official sources.

## Entry template

### YYYY-MM-DD - Topic

Source:

- Title:
- URL:
- Publisher/owner:
- Date checked:

Relevant finding:

Decision:

Risk / uncertainty:

Follow-up:

---

## Initial required research items

### Nakama Go runtime

Research:

- How Go runtime modules are initialized.
- How RPCs are registered.
- Required function signatures.
- Plugin builder requirements.
- Version compatibility requirements.

Decision to record:

- selected Nakama version;
- selected `nakama-common` version;
- selected plugin-builder image;
- whether Go runtime code is compiled as a shared object.

---

### Nakama Docker/PostgreSQL setup

Research:

- official Docker Compose pattern;
- required ports;
- PostgreSQL connection settings;
- local config file approach.

Decision to record:

- `docker-compose.yml` layout;
- local credentials;
- exposed ports;
- whether local defaults are safe to commit.

---

### Nakama RPC registration

Research:

- how to register Go runtime RPCs;
- RPC handler signatures;
- how to return JSON;
- how to return errors.

Decision to record:

- final RPC registration approach;
- handler function organization;
- helper function extraction strategy.

---

### Server-to-server/private RPC behavior

Research:

- how Nakama distinguishes server-to-server calls from user-session calls;
- how runtime HTTP key calls are made;
- how to reject user-session calls.

Decision to record:

- private RPC protection strategy;
- expected positive and negative tests.

---

### Go testing and style

Research:

- current stable Go release;
- Go formatting/style guidance;
- table-driven testing guidance;
- whether fuzz tests are useful for JSON payload parsing;
- compatibility with Nakama plugin build.

Decision to record:

- Go version used locally;
- Go version expected inside Nakama/plugin build;
- test strategy;
- commands required before completion.

---

## Research entries

Add entries below this line during Phase 0.
