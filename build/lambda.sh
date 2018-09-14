#!/usr/bin/env bash

mkdir -p out
GOOS=linux go build -o out/lambda_handler cmd/go-call-me-maybe/go-call-me-maybe.go

# create the lambda archive
cd out
zip lambda.zip lambda_handler
rm lambda_handler

echo "complete - your lambda is called 'lambda_handler'"
