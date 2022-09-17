#/bin/env bash

BUILD_TIME=$(date +%Y%m%d%H%M%S)
BUILD_VERSION="v0.0.1-beta"
GO_VERSION=$(go version | { read _ _ v _; echo ${v#go}; })


go build -a -tags=xtunnel -ldflags "-w -s -X main.BUILD_TIME=$BUILD_TIME -X main.VERSION=$BUILD_VERSION -X main.GO_VERSION=$GO_VERSION" ./.