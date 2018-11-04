# Pseudo service

Pseudo random (deterministic) data generator as a (micro) service.

#### The problem
The need to generate random (user or other type) data, to (stress) test your own system. 

#### The solution
I've built this project to be used as an internal service, to generate large amounts of fake data.

It was originally built for this project: [davidescus/10minFor300Gb](https://github.com/davidescus/10minFor300Gb) as a wrapper on [brianvoe/gofakeit](https://github.com/brianvoe/gofakeit).

### Performance 
 
On my localhost (4 Core 3Ghz) I was able to generate over **14.000 users/second** (`/users/`) with 100 batches, on an 600Mhz B2 AppEngine (1 Core) can deliver 900 users/second. 
Future improvements will be done. To run your own benchmarks see [call_benchmark.sh](./call_benchmark.sh)

For the best performance: 
* Count (batch) value: **100**
* concurrent requests: CPUCores * 1.5
* it does not scale well on multiple cores, it uses the math.Rand that has a mutex, recommended 3-6 cores

### How it works

You run it as a simple HTTP server, in a private network, and clients can make requests to get random data. 

 All the endpoints have 2 modes:
* deterministic - for each `?seed=42` received as input, the same data will be generated.
* random - if a `seed` is not given, random data will be generated. This is the most optimal method (performance).

### Endpoints (data types)

###### /users/{count}
`count` = how many users should generate, an integer between 1 and 500.


```bash
#fast and random user
curl "http://localhost:8080/api/v1/users/1?token=SECRET42"
{"seed":6124420007038740542,"nextseed":0,"users":[
    {
        "id":"a4419f4f-f7f9-4c99-a746-5799b281a5d6","
        age":209,
        "name":"Jonas Borer",
        "company":"core frictionless Inc",
        "position":"Senior Paradigm Associate",
        "email":"JonasBorer@leadsolutions.org",
        "country":"Yemen"
    }
]}

#slow and deterministic call (seed=42)
curl "http://localhost:8080/api/v1/users/1?token=SECRET42&seed=42"

{"seed":42,"nextseed":43,"users":[
    {"
        id":"538c7f96-b164-4f1b-97bb-9f4bb472e89f",
        "age":62,
        "name":"Jacky Borer",
        "company":"Ameliorated mindshare Inc",
        "position":"International Assurance Orchestrator",
        "email":"JackyBorer@leadsolutions.org",
        "country":"Greece"
    }
]}
```

### How to install

a). Download the [binaries from a Release](https://github.com/bgadrian/pseudoservice/releases)

OR b). Get the [docker image from bgadrian/pseudoservice/](https://hub.docker.com/r/bgadrian/pseudoservice/)
```bash
docker run --name pseudoservice -p 8080:8080 -d -e APIKEY=MYSECRET bgadrian/pseudoservice
curl "http://localhost:8080/api/v1/users/1?token=MYSECRET&seed=42"

```

OR c). Hard way, requires Go 1.11, make and git
```bash
git clone git@github.com:bgadrian/pseudoservice.git
cd pseudoservice
make run
#OR
make build
env PORT=8080 ./build/pseudoservice
```

### How to use

The binary has the following env variables:
* `PORT` - http listening port (8080)
* `APIKEY` - secret string to be served at `?token=SECRET42` 
* for more see `pseudoservice --help`

All the endpoint have the following query params:
* `token` - required, the APIKEY
* `seed` - optional, for deterministic responses


The server accepts `gzip`.

The server was generated using swagger (open API), to see a full documentation of the service access the [127.0.0.1:8080/docs](http://127.0.0.1:8080/docs) endpoint.

### TODO
* make the docker multi-stage build work (lower the container from 700mb to 7mb)
* add a general/custom endpoint, where the payload is an object with the types of data it requires to be generated (each client decide what objects to be generated)

### Copyright
Bledea Georgescu Adrian 

Free to use but not for commercial purposes.


 

