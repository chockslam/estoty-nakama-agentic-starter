# RPC Contracts

This file defines the intended public contract for the three assignment RPCs.

Contracts may be refined during implementation, but any change must be reflected here and in the README.

## RPC 1: `update_user_metadata`

### Purpose

Update the authenticated caller's Nakama account metadata by adding arbitrary JSON information from the caller.

### Auth requirement

Requires a valid user session.

Must update only the caller's own account.

Must reject calls without authenticated user context.

### Request body

The request body should be a JSON object.

Example:

```json
{
  "favoriteHero": "warrior",
  "tutorialCompleted": true,
  "preferredDifficulty": "normal"
}
```

### Payload rules

- Payload must be valid JSON.
- Payload must be an object.
- Payload is merged shallowly into the caller's existing metadata.
- Incoming values overwrite existing keys with the same name.
- Nested object deep-merge is intentionally out of scope.
- Payload must not be allowed to select a target user ID or account.
- Payload may contain arbitrary caller-supplied keys and values.

### Success response

Preferred example:

```json
{
  "success": true,
  "metadata": {
    "favoriteHero": "warrior",
    "tutorialCompleted": true,
    "preferredDifficulty": "normal"
  }
}
```

Exact response may differ if Nakama conventions make another shape cleaner.

### Error cases

- Missing authenticated user.
- Invalid JSON.
- JSON payload is not an object.
- Metadata update fails internally.

### Local verification flow

```bash
SESSION_TOKEN="$(./scripts/auth-device.sh)"
SESSION_TOKEN="$SESSION_TOKEN" ./scripts/rpc-update-metadata.sh
```

## RPC 2: `get_game_config`

### Purpose

Return free-form game configuration JSON loaded from `config/game_config.json`.

### Auth requirement

No user session is required.

This RPC is public because it returns non-sensitive game configuration.

It is intentionally not the private/server-to-server assignment RPC; that boundary is covered by `private_health_check`.

The HTTP helper for local verification uses Nakama's runtime HTTP key, since `RegisterRpc` requests still require either a session token or an HTTP key at the transport layer.

### Response body

Must include at least:

```json
{
  "welcomeMessage": "Welcome to the game!",
  "xpRate": 1.0,
  "rarityOptions": ["common", "rare", "epic", "legendary"]
}
```

### Validation rules

- `welcomeMessage` must be a non-empty string.
- `xpRate` must be a positive number.
- `rarityOptions` must be a non-empty list of non-empty strings.

### Error cases

- Config file missing, if file-based config is used.
- Invalid config JSON.
- Missing required fields.
- Invalid field types.

### Local verification flow

```bash
./scripts/rpc-get-game-config.sh
```

## RPC 3: `private_health_check`

### Purpose

Prove private server-to-server RPC authorization boundary.

The assignment does not require meaningful response content. It only requires successful status for valid private/server-to-server call.

### Auth requirement

Must be callable only through the server-to-server/runtime key path.

Must reject normal user-session calls.

### Request body

No request body required.

Acceptable payload:

```json
{}
```

### Success response

Preferred response:

```json
{}
```

Alternative acceptable response if easier to verify:

```json
{
  "success": true
}
```

If using a non-empty response, document why.

### Error cases

- Called by normal authenticated user.
- Called without server-to-server/runtime key.
- Missing or invalid runtime HTTP key.

### Local verification flow

```bash
./scripts/rpc-private-health.sh
MODE=user SESSION_TOKEN="$(./scripts/auth-device.sh)" ./scripts/rpc-private-health.sh
```

## Naming stability

Do not rename RPCs after implementation begins unless there is a strong reason.

Stable planned names:

```text
update_user_metadata
get_game_config
private_health_check
```
