#!/bin/bash
# This script reads Go source external dependencies from
# the file requrements.txt and installs them in $GOPATH.
while read line          
do
    echo "  '$line'"
    go get $line
done < $1
