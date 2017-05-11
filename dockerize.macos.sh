#!/bin/bash

export GOOS=linux

go get; go build -o interestcalculator-linux-amd64

export GOOS=darwin

docker-compose up
