#!/bin/bash

# Example: ./build.sh linux

set -e

if [ -z "$1" ]; then
  echo "Provide build option (linux, darwin, windows)"
  exit 1
fi

GOOS="$1"
GOARCH=""
OUTPUT="coc"

# Determine GOARCH and output name based on GOOS
case "$GOOS" in
  linux)
    GOARCH="amd64"
    ;;
  darwin)
    GOARCH="arm64"
    ;;
  windows)
    GOARCH="amd64"
    OUTPUT="coc.exe"
    ;;
  *)
    echo "Unsupported GOOS: $GOOS"
    exit 1
    ;;
esac

echo "Building for $GOOS/$GOARCH..."

CGO_ENABLED=0 GOOS="$GOOS" GOARCH="$GOARCH" go build -o ../"$OUTPUT" ../main.go

echo "Build complete: $OUTPUT"
