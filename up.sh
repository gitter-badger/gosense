#!/bin/bash
set -ex
DNSSERVERS=" --dns=208.67.222.222 --dns=208.67.220.220 --dns=8.8.8.8 --dns=8.8.4.4 "
 if [ ! -f config.toml ]; then
    cp config.toml.dist config.toml
fi
if [ $(docker network ls | grep gosense-network | wc -l ) -eq 0 ]; then
    docker network create -d bridge gosense-network
fi
docker run ${DNSSERVERS} --rm --name go-build -v $(pwd):/www golang /bin/bash /www/build.sh
if [ $(docker ps -a | grep gs_db | wc -l) -le 0 ]; then
    docker run --restart=always --net=gosense-network -d --name gs_db  netroby/docker-mysql
    while true; do
        if [ $(docker logs gs_db 2>&1 | grep "ready for connections" | wc -l)  -ge 2 ]; then
            break;
        else
            echo "not ready, waiting"
            sleep 2
        fi
    done
else
    docker network disconnect gosense-network gs_db
    docker network connect gosense-network gs_db
fi


# If database not found, then we import database 
if [ $(docker exec gs_db mysql -u root -h 127.0.0.1 -e "use gosense;" | grep -i unkown  | wc -l) -eq 1 ]; then
    docker cp sql/bak.sql gs_db:/root/
    docker exec gs_db sh -c "mysql < /root/bak.sql"
fi

if [ $(docker ps -a | grep gosense | wc -l) -eq 1 ]; then
    docker rm -vf gosense
fi
docker run ${DNSSERVERS}  --restart=always --net=gosense-network -d -p 8080:8080  -v $(pwd):/www --name gosense golang  sh -c "cd /www && /www/gosense"
./catlog.sh -f
