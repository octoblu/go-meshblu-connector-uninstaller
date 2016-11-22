FROM golang:1.6
MAINTAINER Octoblu, Inc. <docker@octoblu.com>

WORKDIR /go/src/github.com/octoblu/go-meshblu-connector-installer
COPY . /go/src/github.com/octoblu/go-meshblu-connector-installer

RUN env CGO_ENABLED=0 go build -o go-meshblu-connector-installer -a -ldflags '-s' .

CMD ["./go-meshblu-connector-installer"]
