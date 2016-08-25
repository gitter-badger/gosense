#!/bin/bash
cd /www ;
go get -d -v ./...
go build -ldflags "-linkmode external -extldflags -static" -o /www/gosense