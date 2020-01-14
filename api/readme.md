# LeBonCoin FizzBuzz (LBCFB) - API

Wanna play right now ? A demo has been deployed on Amazon Web Services :

https://lbcfb.antoineviau.com/api/?string1=fizz&string2=buzz&int1=3&int2=5&limit=20

The requests are sent to AWS API Gateway, then processed by a Lambda function (see `src/lambda/lambda.go`).

## Dev, what you need

Docker and Docker Compose.  
Optional : Go.  
Optional : an Amazon Web Services account and AWS CLI installed/configured.

## About the code

The fizzbuzz code business is in the src/fizzbuzz package.  
The api package and the server (`main.go`) are based on [Chi](https://github.com/go-chi/chi)  
Lambda code is in... `src/lambda`.  
**CORS is managed at the proxy level. So, there will be CORS for the AWS/Lambda version (managed by AWS API Gateway), but not for the Docker/native version.**

## Build and run

    git clone https://github.com/AntoineViau/lbcfb.git

### With Go on your computer

Build :

    cd src
    go build main.go

Core business unit tests :

    cd src/fizzbuzz
    go test

Run :

    cd src
    go build
    ./main

Manual testing with http://localhost:8080/api/v1/fizzbuzz?string1=fizz&string2=buzz&int1=3&int2=5&limit=20

Integration tests :

    cd src/api
    go test

### With Docker

**Theses containers are to be used in a CI/CD process.**  
Building container will include unit tests.

    cd docker
    docker-compose build --no-cache app
    docker-compose run app

Running integration tests :

    cd docker
    docker-compose build --no-cache integration
    docker-compose run integration

## Deploy on AWS

A shell script is here for the heavy lifting : `deploy-lambda.sh`  
The first time use :

    ./deploy-lambda.sh install

To update just launch :

    ./deploy-lambda.sh
