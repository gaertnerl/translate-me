#! /bin/bash

docker build -t webapp --build-arg REGISTER_SIM_ENDPOINT_KEY="key" .
docker rm webapp-instance
docker run -p 8080:8080 --name webapp-instance webapp