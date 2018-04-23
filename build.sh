#!/usr/bin/env bash

export GOPATH=$(pwd $(dirname $0))

# build all
go install server