#!/bin/bash
set -ex
if [ ! -f config.toml ]; then
    cp config.toml.dist config.toml
fi
if [ ! -z $1 ]; then
    go get github.com/tools/godep
    godep get -v
    godep go build
fi
PIDGS=$(docker exec gosense pidof gosense)
echo "Pid of gosense $PIDGS"
docker exec gosense kill -s HUP $PIDGS
docker ps -a
docker logs -f -t --tail 30  gosense

