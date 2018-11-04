#!/usr/bin/env bash

#deterministic (specific seed)
curl --compressed -sH 'Accept-encoding: gzip' -X GET "http://localhost:8080/users/3?token=SECRET42&seed=66"

#random
curl --compressed -sH 'Accept-encoding: gzip' -X GET "http://localhost:8080/users/3?token=SECRET42"