#!/usr/bin/env bash

export GOPATH=$(pwd $(dirname $0))

# build all
go build -ldflags "-X mtlog.CodeRoot=${GOPATH}" server
go install server