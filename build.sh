#!/usr/bin/env bash

export GOPATH=$(pwd $(dirname $0))

# build all
go build -ldflags "-X mtlog.CodeRoot=${GOPATH}" server
go install server
cp -rf ${GOPATH}/bin/server /Users/liyanguo/tmp/http_server