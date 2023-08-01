FROM golang:1.20.6-alpine3.18
WORKDIR /opt
RUN apk update
RUN go version
COPY shell.go .
RUN go build shell.go
