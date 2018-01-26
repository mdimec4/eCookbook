#!/bin/bash

go get -u github.com/golang/dep/cmd/dep
dep ensure -v
if ! go test eCoockbook/cmd/chef
then
    exit 1
fi
