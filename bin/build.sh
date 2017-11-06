#!/bin/bash

source ./bin/activate
mkdir build
go build -o build/chef chef
source ./bin/deactivate
