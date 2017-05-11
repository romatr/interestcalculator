#!/bin/bash

go get; go build -o interestcalculator-linux-amd64
docker-compose up --build
