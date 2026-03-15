#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
ROOT_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
RELEASE_DIR="$ROOT_DIR/release"

VERSION="${RELEASE_VERSION:?RELEASE_VERSION must be set}"

rm -rf "$RELEASE_DIR"
mkdir -p "$RELEASE_DIR"

platforms=(
  "linux/amd64"
  "linux/arm64"
  "darwin/amd64"
  "darwin/arm64"
  "windows/amd64"
)

for platform in "${platforms[@]}"; do
  GOOS="${platform%/*}"
  GOARCH="${platform#*/}"

  output="devpod-provider-sprites-${GOOS}-${GOARCH}"
  if [ "$GOOS" = "windows" ]; then
    output="${output}.exe"
  fi

  echo "Building ${output}..."
  CGO_ENABLED=0 GOOS="$GOOS" GOARCH="$GOARCH" go build -ldflags="-s -w" -o "$RELEASE_DIR/$output" "$ROOT_DIR"

  (cd "$RELEASE_DIR" && shasum -a 256 "$output" > "${output}.sha256")
done

echo "Generating provider.yaml..."
go run "$SCRIPT_DIR/provider/main.go" "$VERSION" "$RELEASE_DIR" "$SCRIPT_DIR/provider/provider.yaml"

echo "Release artifacts in $RELEASE_DIR:"
ls -la "$RELEASE_DIR"
