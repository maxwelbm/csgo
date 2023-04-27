#!/bin/sh

set -u
set +x
set +e

GOOS=windows
GOARCH=amd64
CGO_ENABLED=1

VERSION=$(git rev-parse --short HEAD)
DATE=$(date -u +.%Y%m%d.%H%M%S)

go build -v -ldflags "-X github.com/maxwelbm/csboost.Version=$VERSION -X github.com/maxwelbm/csboost.Date=$DATE" -o csboost.exe cmd/csboost/main.go