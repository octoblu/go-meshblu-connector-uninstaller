#!/bin/bash

APP_NAME=go-meshblu-connector-uninstaller
BUILD_DIR=$PWD/dist
IMAGE_NAME=local/$APP_NAME

build_on_docker() {
  docker build --tag $IMAGE_NAME:built .
}

build_on_local() {
  local goos="$1"
  local goarch="$2"

  local output_file="${BUILD_DIR}/${APP_NAME}-${goos}-${goarch}"
  if [ "$goos" == "windows" ]; then
    output_file="${output_file}.exe"
  fi

  env GOOS="$goos" GOARCH="$goarch" \
    go build \
      -a \
      -tags netgo \
      -installsuffix cgo \
      -ldflags '-w' \
      -o "$output_file" \
      .
}

build_osx_on_local() {
  build_on_local "darwin" "amd64"
}

copy() {
  cp $BUILD_DIR/$APP_NAME-linux-amd64 entrypoint/$APP_NAME
}

init() {
  rm -rf $BUILD_DIR/ \
  && mkdir -p $BUILD_DIR/
}

package() {
  docker build --tag $IMAGE_NAME:latest entrypoint
}

run() {
  docker run --rm \
    --volume $BUILD_DIR:/export/ \
    $IMAGE_NAME:built \
      /bin/bash -c "cp dist/* /export" # /bin/bash -c needed for '*' expansion
}

fatal() {
  local message=$1
  echo $message
  exit 1
}

cross_compile_build(){
  for goos in darwin linux windows; do
    for goarch in 386 amd64; do
      echo "building: ${goos}-${goarch}"
      build_on_local "$goos" "$goarch" > /dev/null
    done
  done
}

docker_build() {
  init            || fatal "init failed"
  build_on_docker || fatal "build_on_docker failed"
  run             || fatal "run failed"
  copy            || fatal "copy failed"
  package         || fatal "package failed"
}

osx_build() {
  init               || fatal "init failed"
  build_osx_on_local || fatal "build_osx_on_local failed"
}

release_osx_build() {
  mkdir -p dist \
  && osx_build \
  && tar -czf "${APP_NAME}-osx.tar.gz" "${APP_NAME}" \
  && mv "${APP_NAME}-osx.tar.gz" dist/
  echo "Wrote dist/${APP_NAME}-osx.tar.gz"
}

main() {
  local mode="$1"
  if [ "$mode" == "docker" ]; then
    echo "Docker Build"
    docker_build
    exit $?
  fi

  if [ "$mode" == "osx" ]; then
    echo "OSX Build"
    osx_build
    exit $?
  fi

  if [ "$mode" == "release-osx" ]; then
    echo "Release Build"
    release_osx_build
    exit $?
  fi

  if [ "$mode" == "cross-compile" ]; then
    echo "Release Build"
    cross_compile_build
    exit $?
  fi

  echo "Usage: ./build.sh docker/osx/release-osx/cross-compile"
  exit 1
}
main $@
