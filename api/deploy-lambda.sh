#!/bin/bash

set -ex

install=$1

set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0
cd src/lambda
go build -o ../../main lambda.go 
cd ../..
zip function.zip ./main

if [ "$install" = "install" ]
then 
  aws lambda create-function --function-name lbcfb --runtime go1.x \
    --zip-file fileb://function.zip --handler main \
    --role arn:aws:iam::793918886560:role/lbcfblambda
else
  aws lambda update-function-code --function-name lbcfb --zip-file fileb://function.zip 
fi

rm function.zip
rm ./main

