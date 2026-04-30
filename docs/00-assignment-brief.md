# Assignment Brief

## Source assignment

Create a Docker, PostgreSQL-based Nakama project.

Using Nakama as the base, develop three RPC methods:

1. RPC method that allows updating user/account metadata by adding any information from the caller.
2. RPC method that returns the game configuration file. The configuration file is a free-form JSON structure. It must include:
   - welcome message: string;
   - `xp rate`: float;
   - list of rarity options.
3. Create a private RPC method, where private means it can be called only with the server-to-server key. It should return a successful status code and does not need to return a response message.

Delivery:

- Submit the project as a public GitHub repository link or invite `janis@estoty.com`.
- Add instructions on:
  - additional technologies and libraries used;
  - how to start the project;
  - how to call the relevant RPC methods.

## Engineering interpretation

This is a backend/game-server assignment. The expected output is not a complete game. The expected output is a small, runnable, documented Nakama backend that proves the candidate can work with:

- Go runtime code;
- Nakama RPC registration;
- Docker-based local development;
- PostgreSQL-backed Nakama;
- authenticated player context;
- server-to-server/private call boundaries;
- JSON payload parsing and validation;
- practical README instructions.

## What the employer is likely testing

The assignment tests whether the candidate can:

- understand a short technical specification;
- research an unfamiliar framework;
- set up a working backend service stack;
- write idiomatic Go in a server/runtime context;
- expose RPCs for game-client and backend usage;
- handle authentication and private call boundaries correctly;
- write simple tests around risky logic;
- document how another developer can run and verify the project.

## In scope

- Docker Compose stack for Nakama + PostgreSQL.
- Go runtime module for Nakama.
- Three assignment RPCs.
- Unit tests for pure logic.
- Integration verification using Docker and curl/scripts.
- README for reviewer usage.
- Local-only development configuration.

## Out of scope

Unless all required functionality is complete and there is time left, do not add:

- matchmaking;
- leaderboards;
- real-time multiplayer simulation;
- custom web server;
- custom admin UI;
- Kubernetes;
- cloud deployment;
- payment systems;
- complex migrations;
- external authentication providers;
- non-essential third-party dependencies.

## Success standard

A reviewer should be able to clone the repository, run the Docker stack, authenticate a test user, call all three RPCs, and understand the implementation decisions from the README.
