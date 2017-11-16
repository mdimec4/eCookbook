#!/bin/bash

source ./bin/activate 
source bin/go-depends.sh requirements.txt
mkdir -p ./build
if ! go build -o build/chef chef
then
    source ./bin/deactivate
    exit 1
fi
source ./bin/deactivate
