#!/bin/bash
cd /www ;
go get -u -v github.com/jteeuwen/go-bindata/...;
go get -u -v github.com/elazarl/go-bindata-assetfs/...;
go-bindata-assetfs assets/... templates/...;
go get -u -v ;
go build -o /www/gosense