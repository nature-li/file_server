#!/usr/bin/env bash

export GOPATH=$(pwd $(dirname $0))
go run ${GOPATH}/src/server/*.go