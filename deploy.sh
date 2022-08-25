#!/bin/bash

docker compose build
docker compose down
source .env
docker compose up