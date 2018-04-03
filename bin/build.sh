#!/bin/bash

go get github.com/golang/dep/cmd/dep
dep ensure -v
mkdir -p ./build
if ! go build -o build/eCookbook eCookbook
then
    exit 1
fi
