#!/bin/bash

source ./bin/activate 
source bin/go-depends.sh requirements.txt
mkdir -p ./build
if ! go build -o build/chef chef
    then
      exit 1
fi
source ./bin/deactivate
