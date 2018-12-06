# Pseudo service ![travis-mater](https://travis-ci.com/bgadrian/pseudoservice.svg?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/bgadrian/pseudoservice)](https://goreportcard.com/report/github.com/bgadrian/pseudoservice)  [![contributions](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/bgadrian/pseudoservice/issues)

Pseudo random (deterministic) data generator as a (micro) service.
It is a wrapper for the [FastFaker Go package](https://github.com/bgadrian/fastfaker).

#### The problem
The need to generate random (user or other type) data, to mockup, test or stress  your own systems. 

#### The solution
I've built this project to be used as an internal service, to generate large amounts of fake data.

### Performance 
To run your own benchmarks see [call_benchmark.sh](./call_benchmark.sh). For the best performance I recommend the following:
* Count (batch) value: **100**
* concurrent requests: CPUCores * 1.5

#### Benchmarks
Tests done on Pseudoservice 2.0.0 using ApacheBench 2.3 `ab -n 8000 -c 6 -kdq -m GET  http://localhost:8080/users/100?token=SECRET42`. The tests ignore any network latency, does not unmarshall the response, reuses HTTP connections (keep-alive) and were made from the local machine (where pseudoservice was running), meaning you will have **at least a 20% lower performance** in a real-world scenario.

* **t2.micro** (~1 CPU), AWS Linux
```bash
  Time taken for tests:   7.649 seconds
  Complete requests:      8000
  Requests per second:    1045.86 [#/sec] (mean)
  Time per request:       5.737 [ms] (mean)
  Time per request:       0.956 [ms] (mean, across all concurrent requests)
  Transfer rate:          45236.55 [Kbytes/sec] received
  Translation: 45Mb/s fake data = 104500 fake users/s
```
* **t2.xlarge** (~4 CPU), AWS Linux
```bash
    Time taken for tests:   2.825 seconds
    Requests per second:    2832.02 [#/sec] (mean)
    Time per request:       2.119 [ms] (mean)
    Time per request:       0.353 [ms] (mean, across all concurrent requests)
    Transfer rate:          122314.75 [Kbytes/sec] received
    Translation: 122Mb/s fake data =  283200 users/s
```
* **c5.2xlarge**	(~8 CPU), AWS Linux
```bash
    Time taken for tests:   2.035 seconds
    Total transferred:      353697150 bytes
    HTML transferred:       352137150 bytes
    Requests per second:    3931.71 [#/sec] (mean)
    Time per request:       1.526 [ms] (mean)
    Time per request:       0.254 [ms] (mean, across all concurrent requests)
    Transfer rate:          169755.12 [Kbytes/sec] received
    Translation: 170Mb/s fake data = 393100 users/s
```

### How it works

You run it as a simple HTTP server, in a private network, and the clients can make requests to get random data, based on their custom templates/needs.

Global optional parameters:
* `?seed=42` - if given, the result will be deterministic, as in for all calls the same data will be returned.
* `token=SECRET42` - the APIKEY
* `{count}` - how many results should be generated, integer [1,500]

### Endpoints (data types)

##### /docs 
Contains the OpenID/swagger documentation for this API.

##### /health
Returns `200` if the service is available.

##### /custom/{count}?template="Hello &#126;name&#126;!"
Generate random data based on a given template. Supports any string, including JSON schema.

The template can contains variables like `~name~` or `~email~`. For a full list see the [FastFaker variables](https://github.com/bgadrian/fastfaker/blob/master/TEMPLATE_VARIABLES.md), but instead of `{}` you must use the `~` delimiter (it is more URL friendly).

```bash
#the template is URL encoded: Hello ~name~!
curl -X GET "http://localhost:8080/custom/3?token=SECRET42&seed=42&template=Hello%20~name~%21"
{"results":[
    "Hello Jeromy Schmeler!",
    "Hello Kim Steuber!",
    "Hello Jacky Borer!"
],"seed":42}

#template: {name:"~name~",age:~digit~~digit~}
curl -X GET "http://localhost:8080/custom/3?token=SECRET42&seed=42&template=%22%7Bname%3A%22~name~%22%2Cage%3A~digit~%7D%22"
{"results":[
    "\"{name:\"Jeromy Schmeler\",age:53}\"",
    "\"{name:\"Dustin Jones\",age:62}\"",
    "\"{name:\"Keely Hartmann\",age:12}\""
],"seed":42}

#template: ~country~
curl -X GET "http://localhost:8080/custom/6?token=SECRET42&seed=42&template=~country~"
{"results":[
    "Tajikistan",
    "Cameroon",
    "Cote Divoire",
    "Turkmenistan",
    "Ethiopia",
    "Afghanistan"
],"seed":42}
``` 

For more examples see the FastFaker [examples](https://github.com/bgadrian/fastfaker/tree/master/example) and [GoDoc](https://godoc.org/github.com/bgadrian/fastfaker/faker#pkg-examples).

##### /users/{count}
Generate random users based on a preconfigured set of properties, with a friend relationship between them.

This endpoint was created to simulate a Social network and test a Graph database: [davidescus/10minFor300Gb](https://github.com/davidescus/10minFor300Gb).

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

#slower but deterministic call (seed=42)
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
```bash
wget https://github.com/bgadrian/pseudoservice/releases/download/v2.0.0/pseudoservice.tar.gz
tar -xzf pseudoservice.tar.gz
./pseudoservice/linux/pseudoservice --api-key=SECRET42 --read-timeout=1s --write-timeout=1s --keep-alive=15s --listen-limit=1024 --max-header-size=3KiB --host=0.0.0.0 --port=8080
```

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

The server accepts `gzip` and it was generated using swagger (open API), to see a full documentation of the service access the [127.0.0.1:8080/docs](http://127.0.0.1:8080/docs) endpoint.

### TODO
All the [issues](https://github.com/bgadrian/pseudoservice/issues)

### Copyright
Bledea Georgescu Adrian https://coder.today

