#!/bin/bash

pushd .
cd src/web/spa
if ! npm install
then
	exit 1
fi
if ! npm run build
then
	exit 1
fi
popd
mkdir -p build
cp -r src/web/spa/dist ./build
