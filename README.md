# Interest Calculator
Simple microservice for interest rate calculations written in Go.

## How to run and test

* Under MacOS
```
$ ./dockerize.macos.sh
```

* or under another OS
```
$ ./dockerize.sh
```

* In your browser type
```
http://localhost:8081/calculate/verso?amount=30000&interest_rate=1.65&term_in_months=36
```

***
If you want to run locally
```
$ go get
$ go run *.go

```
and use port **8080** to access the service via your browser