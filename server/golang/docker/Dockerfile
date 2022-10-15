FROM golang:1.19-alpine

WORKDIR /app

COPY . /app

RUN apk add make
RUN make build

ENTRYPOINT ["/app/bin/snmp-browser"]
