#!/usr/bin/env bash

make build && \
env APIKEY=MYSECRET PORT=8080 ./build/server