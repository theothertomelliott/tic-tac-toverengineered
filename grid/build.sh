#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

pushd $DIR > /dev/null

CGO_ENABLED=0 GOOS=linux go build \
    -ldflags "-X github.com/theothertomelliott/tic-tac-toverengineered/common/version.Version=$VERSION" \
    -o .output/app ./cmd/grid

popd > /dev/null