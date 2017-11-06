#!/bin/bash

pushd .
cd src/web/spa
npm install
npm run build
popd
mkdir build
cp -r src/web/spa/dist ./build
