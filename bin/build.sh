#!/bin/bash

go get -u github.com/golang/dep/cmd/dep
dep ensure -v
mkdir -p ./build
if ! go build -o build/chef eCookbook/cmd/chef
then
    exit 1
fi
