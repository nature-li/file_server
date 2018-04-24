#!/usr/bin/env bash

export GOPATH=$(pwd $(dirname $0))
go run ${GOPATH}/src/http_server/*.go --conf=${GOPATH}/config/conf.yaml
