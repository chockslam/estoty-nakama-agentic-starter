#!/usr/bin/env bash

set -euo pipefail

NAKAMA_URL="${NAKAMA_URL:-http://127.0.0.1:7350}"
NAKAMA_HTTP_KEY="${NAKAMA_HTTP_KEY:-defaulthttpkey}"
MODE="${MODE:-private}"
response_file="${TMPDIR:-/tmp}/nakama-private-health.$$"

trap 'rm -f "$response_file"' EXIT

case "$MODE" in
private)
	curl -sS --fail --retry 10 --retry-connrefused --retry-delay 1 \
		-X POST "${NAKAMA_URL}/v2/rpc/private_health_check?http_key=${NAKAMA_HTTP_KEY}&unwrap=true" \
		-H 'Content-Type: application/json' \
		-H 'Accept: application/json' \
		--data '{}' \
		--output "${response_file}"
	cat "${response_file}"
	;;
user)
	: "${SESSION_TOKEN:?set SESSION_TOKEN to a Nakama session token}"
	status="$(curl -sS -o "${response_file}" -w '%{http_code}' \
		-X POST "${NAKAMA_URL}/v2/rpc/private_health_check?unwrap=true" \
		-H "Authorization: Bearer ${SESSION_TOKEN}" \
		-H 'Content-Type: application/json' \
		-H 'Accept: application/json' \
		--data '{}')"
	if [ "$status" -lt 400 ]; then
		printf 'expected private RPC to fail for user session, got HTTP %s\n' "$status" >&2
		cat "${response_file}" >&2
		exit 1
	fi
	printf 'HTTP %s\n' "$status"
	cat "${response_file}" >&2
	exit 1
	;;
*)
	printf 'unknown MODE: %s\n' "$MODE" >&2
	exit 2
	;;
esac
