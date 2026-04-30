#!/usr/bin/env bash

set -euo pipefail

: "${SESSION_TOKEN:?set SESSION_TOKEN to a Nakama session token}"

NAKAMA_URL="${NAKAMA_URL:-http://127.0.0.1:7350}"
PAYLOAD="${PAYLOAD:-{\"favoriteHero\":\"warrior\",\"tutorialCompleted\":true,\"preferredDifficulty\":\"normal\"}}"

curl -sS -X POST "${NAKAMA_URL}/v2/rpc/update_user_metadata?unwrap=true" \
	--fail \
	--retry 10 \
	--retry-connrefused \
	--retry-delay 1 \
	-H "Authorization: Bearer ${SESSION_TOKEN}" \
	-H 'Content-Type: application/json' \
	-H 'Accept: application/json' \
	--data "${PAYLOAD}"
