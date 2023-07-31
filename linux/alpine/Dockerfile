#FROM alpine:latest
FROM golang:1.20.6-alpine3.18
WORKDIR /opt
RUN apk update
#RUN apk add wget git
#RUN apk add autoconf build-base binutils cmake make curl file gcc g++ git libgcc libtool linux-headers musl-dev ninja tar unzip wget gcompat
#RUN wget https://golang.org/dl/go1.20.6.linux-amd64.tar.gz
#RUN tar -C /usr/local/ -xzf go1.20.6.linux-amd64.tar.gz
#ENV PATH=$PATH:/usr/local/go/bin
RUN go version
COPY shell.go .
#RUN CGO_ENABLED=0 go build shell.go
RUN go build shell.go
