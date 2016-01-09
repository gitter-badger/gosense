#!/bin/bash
set -ex
if [ ! -f config.toml ]; then
    cp config.toml.dist config.toml
fi
if [ ! -z $1 ]; then
DNSSERVERS=" --dns=208.67.222.222 --dns=208.67.220.220 --dns=8.8.8.8 --dns=8.8.4.4 "
docker run ${DNSSERVERS} --rm --name go-build -v $(pwd):/www golang /bin/bash /www/build.sh
fi
PIDGS=$(docker exec gosense pidof gosense)
echo "Pid of gosense $PIDGS"
docker exec gosense kill -s HUP $PIDGS
docker ps -a
docker logs -f -t --tail 30  gosense

