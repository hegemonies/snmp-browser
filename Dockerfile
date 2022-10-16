FROM node:18 AS frontend-builder
WORKDIR /frontend
COPY ./frontend .
RUN make build

FROM golang:1.19-alpine as server-builder
WORKDIR /server
COPY ./server/golang .
RUN apk add make
RUN make build

FROM alpine
WORKDIR /app
COPY --from=frontend-builder /frontend/build /app/frontend
COPY --from=server-builder /server/bin/snmp-browser /app/server/snmp-browser
ENTRYPOINT ["/app/server/snmp-browser"]
