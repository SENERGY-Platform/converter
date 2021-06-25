#!/bin/bash

arch="$(go env GOARCH)"

if [[ $arch == "amd64" ]]
then
  apt-get update
  apt-get install -y gcc-arm*
  CGO_ENABLED=1 GOOS=linux go build -o converter.so -buildmode=c-shared ./cmd/sharedlib
  CGO_ENABLED=1 CC=arm-linux-gnueabi-gcc GOOS=linux GOARCH=arm GOARM=7 go build -o converter_arm.so -buildmode=c-shared ./cmd/sharedlib
fi

if [[ $arch == "arm" ]]
then
  CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/rest
  CGO_ENABLED=1 GOOS=linux go build -o converter.so -buildmode=c-shared ./cmd/sharedlib
  CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=7 go build -o converter_arm.so -buildmode=c-shared ./cmd/sharedlib
fi