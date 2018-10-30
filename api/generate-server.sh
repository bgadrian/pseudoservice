#!/usr/bin/env bash

docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate \
    -i /local/swagger.yaml \
    -g go-server \
    -o /local/openapi && \
sudo chown -R ${USER} openapi
#rm -rf openapi/api/ && \
#mv openapi/ ../