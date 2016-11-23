FROM golang:1.7
MAINTAINER Octoblu, Inc. <docker@octoblu.com>

WORKDIR /go/src/github.com/octoblu/go-meshblu-connector-uninstaller
COPY . /go/src/github.com/octoblu/go-meshblu-connector-uninstaller

RUN ./build.sh cross-compile

CMD ["./dist/go-meshblu-connector-uninstaller-linux-amd64"]
