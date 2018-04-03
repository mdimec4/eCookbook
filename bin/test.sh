#!/bin/bash

go get github.com/golang/dep/cmd/dep
dep ensure -v
if ! go test eCookbook
then
    exit 1
fi
