FROM golang:1.6
MAINTAINER Octoblu, Inc. <docker@octoblu.com>

WORKDIR /go/src/github.com/octoblu/go-meshblu-connector-uninstaller
COPY . /go/src/github.com/octoblu/go-meshblu-connector-uninstaller

RUN env CGO_ENABLED=0 go build -o go-meshblu-connector-uninstaller -a -ldflags '-s' .

CMD ["./go-meshblu-connector-uninstaller"]
