#!/bin/bash
# This script reads Go source external dependencies from
# the file requrements.txt and installs them in $GOPATH.
while read line          
do
    echo "  '$line'"
    if ! go get $line
    then
	    exit 1
    fi
done < $1
