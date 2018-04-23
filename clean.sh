#!/usr/bin/env bash

set -x

# set GOPATH
export GOPATH=$(pwd $(dirname $0))

# build http_server
go clean http_server

# rm all collected files
TARGET_PATH=${GOPATH}/target
rm -rf ${TARGET_PATH}


