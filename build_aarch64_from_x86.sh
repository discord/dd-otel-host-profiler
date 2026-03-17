#!/usr/bin/env bash

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$ROOT_DIR"

VERSION="${VERSION:-${1:-test}}"
CC_BIN="${CC_BIN:-aarch64-linux-gnu-gcc}"

require_cmd() {
  if ! command -v "$1" >/dev/null 2>&1; then
    echo "missing required command: $1" >&2
    exit 1
  fi
}

require_cmd "$CC_BIN"
require_cmd make
require_cmd file

echo "Building dd-otel-host-profiler for linux/arm64"
echo "version: $VERSION"
echo "cc: $CC_BIN"

CC="$CC_BIN" \
GOOS=linux \
GOARCH=arm64 \
CGO_ENABLED=1 \
make VERSION="$VERSION"

echo
file ./dd-otel-host-profiler
