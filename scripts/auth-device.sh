#!/usr/bin/env bash

set -euo pipefail

NAKAMA_URL="${NAKAMA_URL:-http://127.0.0.1:7350}"
NAKAMA_SERVER_KEY="${NAKAMA_SERVER_KEY:-defaultkey}"
DEVICE_ID="${DEVICE_ID:-estoty-local-device}"
USERNAME="${USERNAME:-estoty-local-device}"
response_file="${TMPDIR:-/tmp}/nakama-auth-device.$$"

trap 'rm -f "$response_file"' EXIT

curl -sS "${NAKAMA_URL}/v2/account/authenticate/device?create=true&username=${USERNAME}" \
	-u "${NAKAMA_SERVER_KEY}:" \
	-H 'Content-Type: application/json' \
	-H 'Accept: application/json' \
	--data "{\"id\":\"${DEVICE_ID}\"}" \
	--output "${response_file}"

token="$(sed -n 's/.*"token":"\([^"]*\)".*/\1/p' "${response_file}")"

if [ -z "$token" ]; then
	cat "${response_file}" >&2
	exit 1
fi

printf '%s\n' "$token"
