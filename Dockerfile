FROM iron/base
EXPOSE 8080

ADD interestcalculator-linux-amd64 /

ENTRYPOINT ["./interestcalculator-linux-amd64"]