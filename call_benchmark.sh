#!/usr/bin/env bash

#-n total count of request, -c concurrent calls ....... /users/X = batch size
echo "50 users * "
ab -n 2000 -c 6 -kdq -m GET  http://localhost:8080/api/v1/users/50?token=SECRET42 | grep "Requests per second"
echo "100 users * "
ab -n 2000 -c 6 -kdq -m GET  http://localhost:8080/api/v1/users/100?token=SECRET42 | grep "Requests per second"
echo "150 users * "
ab -n 2000 -c 6 -kdq -m GET  http://localhost:8080/api/v1/users/150?token=SECRET42 | grep "Requests per second"
echo "500 users * "
ab -n 2000 -c 6 -kdq -m GET  http://localhost:8080/api/v1/users/500?token=SECRET42 | grep "Requests per second"
