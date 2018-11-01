#!/usr/bin/env bash

docker stop pseudoservice
docker rm pseudoservice
docker rmi bgadrian/pseudoservice

docker build -t bgadrian/pseudoservice:latest .
docker run -d -p 8080:8080 --name pseudoservice bgadrian/pseudoservice:latest