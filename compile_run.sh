#!/usr/bin/env sh

set -e

go build -o build pkg/main.go
./build/main
