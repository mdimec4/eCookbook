#!/bin/bash

source ./bin/activate
source bin/go-depends.sh requirements.txt
if ! go test chef
then
    source ./bin/deactivate
    exit 1
fi
source ./bin/deactivate
