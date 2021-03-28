#!/bin/sh

APP="mim"

go run assets/assets_generate.go

# os
GOOS=linux GOARCH=amd64 go build -ldflags "-X main.SHA1VER=`git rev-parse HEAD` -X main.BUILDTIME=`date -u +%Y-%m-%dT%H:%M:%S`" -o ./bin/linux/${APP} *.go
GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.SHA1VER=`git rev-parse HEAD` -X main.BUILDTIME=`date -u +%Y-%m-%dT%H:%M:%S`" -o ./bin/darwin/${APP} *.go


