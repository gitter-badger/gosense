FROM golang:1.7

ADD . /go/src/gosense/


WORKDIR /go/src/gosense

RUN go get -v -d ... ; \
        go get -v github.com/jteeuwen/go-bindata/...; \
        go get -v github.com/elazarl/go-bindata-assetfs/...; \
        go-bindata-assetfs assets/... templates/...; \
        go build -ldflags "-linkmode external -extldflags -static" -v


CMD ["/go/src/gosense/gosense"]


