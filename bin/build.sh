#!/bin/bash

source ./bin/activate 
source bin/go-depends.sh requirements.txt
mkdir -p ./build
go build -o build/chef chef
source ./bin/deactivate
