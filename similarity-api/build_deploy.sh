#! /bin/bash

docker build -t similarity-api --build-arg API_HOST=0.0.0.0 --build-arg API_PORT=8000 --build-arg LOADBALANCER_API_KEY='key' --build-arg LOADBALANCER_REGISTER_URL='http://0.0.0.0/api/register' .
docker rm similarity-api-instance
docker run -p 8080:8000 --name similarity-api-instance similarity-api