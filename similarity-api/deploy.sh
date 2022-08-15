#! /bin/bash

docker rm similarity-api-instance
docker run -p 80:8000 --name similarity-api-instance similarity-api
