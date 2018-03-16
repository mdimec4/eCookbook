#!/bin/bash

go get github.com/golang/dep/cmd/dep
dep ensure -v
if ! go test eCookbook/cmd/chef
then
    exit 1
fi
