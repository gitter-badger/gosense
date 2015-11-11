#!/bin/bash
set -ex
if [ ! -f config.toml ]; then
    cp config.toml.dist config.toml
fi
go get github.com/tools/godep
godep get -v
godep go build
docker kill -s HUP gosense
