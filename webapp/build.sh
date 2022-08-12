#! bin/bash

npx webpack
rm -r ../webserver/static
mkdir ../webserver/static
cp -r ./src/html ../webserver/static/html
cp -r ./src/compiled_js ../webserver/static/compiled_js
cp -r ./src/style ../webserver/static/style