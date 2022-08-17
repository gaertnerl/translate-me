#! /bin/bash

BASE_PATH=../webserver/src/static
npx webpack
rm -r $BASE_PATH
mkdir $BASE_PATH
cp -r ./src/html $BASE_PATH/html
cp -r ./src/compiled_js $BASE_PATH/compiled_js
cp -r ./src/style $BASE_PATH/style
cp -r ./src/favicon.ico $BASE_PATH/favicon.ico