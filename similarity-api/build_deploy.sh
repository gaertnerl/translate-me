#! /bin/bash

docker build -t similarity-api .
docker rm similarity-api-instance
docker run -p 8080:8000 --name similarity-api-instance-0 similarity-api