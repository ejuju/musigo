#!/bin/bash

set -e # show errors
set -x # show commands

go mod tidy
go mod verify
go vet ./...

CGO_ENABLED=1 go test ./... \
    -cover \
    -race \
    -timeout=60s