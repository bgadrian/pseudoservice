# Pseudo service ![travis-mater](https://travis-ci.com/bgadrian/pseudoservice.svg?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/bgadrian/pseudoservice)](https://goreportcard.com/report/github.com/bgadrian/pseudoservice)

Pseudo random (deterministic) data generator as a (micro) service.
It is a wrapper for the [FastFaker Go package](https://github.com/bgadrian/fastfaker).

#### The problem
The need to generate random (user or other type) data, to mockup, test or stress  your own systems. 

#### The solution
I've built this project to be used as an internal service, to generate large amounts of fake data.

### Performance 
 
On my localhost (4 Core 3Ghz) I was able to generate over **14.000 users/second** (`/users/`) with 100 batches, on an 600Mhz B2 AppEngine (1 Core) can deliver 900 users/second. 
Future improvements will be done. To run your own benchmarks see [call_benchmark.sh](./call_benchmark.sh)

For the best performance: 
* Count (batch) value: **100**
* concurrent requests: CPUCores * 1.5
* it does not scale well on multiple cores, it uses the math.Rand that has a mutex, recommended 3-6 cores

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

