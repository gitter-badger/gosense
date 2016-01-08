#!/bin/bash
set -ex
if [ ! -f config.toml ]; then
    cp config.toml.dist config.toml
fi
if [ ! -z $1 ]; then
    docker run ${DNSSERVERS} --rm --name go-build -v $HOME/go:/go -v $(pwd):/www golang sh -c "cd /www ;go get github.com/jteeuwen/go-bindata/...;go get github.com/elazarl/go-bindata-assetfs/...;go-bindata-assetfs assets/... templates/...;go get -v ; go build -o /www/gosense "
fi
PIDGS=$(docker exec gosense pidof gosense)
echo "Pid of gosense $PIDGS"
docker exec gosense kill -s HUP $PIDGS
docker ps -a
docker logs -f -t --tail 30  gosense

