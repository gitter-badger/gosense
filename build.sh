#!/bin/bash
cd /www ;
go get -v github.com/jteeuwen/go-bindata/...;
go get -v github.com/elazarl/go-bindata-assetfs/...;
go-bindata-assetfs assets/... templates/...;
go get -d -v ./...
go build -ldflags "-linkmode external -extldflags -static" -o /www/gosense