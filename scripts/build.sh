#!/usr/bin/env bash

gofmt -w .

GOOS=linux GOARCH=amd64 \
go build -buildmode=pie \
    -ldflags="-linkmode=external -s -w -bindnow" \
    -o ./bin/orbital-elements ./cmd/orbital-elements/main.go

