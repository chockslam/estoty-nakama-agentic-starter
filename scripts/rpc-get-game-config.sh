#!/usr/bin/env bash

set -euo pipefail

NAKAMA_URL="${NAKAMA_URL:-http://127.0.0.1:7350}"
NAKAMA_HTTP_KEY="${NAKAMA_HTTP_KEY:-${NAKAMA_RUNTIME_HTTP_KEY:-defaulthttpkey}}"

curl -sS --fail --retry 10 --retry-connrefused --retry-delay 1 \
	-X POST "${NAKAMA_URL}/v2/rpc/get_game_config?http_key=${NAKAMA_HTTP_KEY}&unwrap=true" \
	-H 'Content-Type: application/json' \
	-H 'Accept: application/json' \
	--data '{}'
