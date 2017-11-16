#!/bin/bash

source ./bin/activate
source bin/go-depends.sh requirements.txt
go test chef
source ./bin/deactivate
