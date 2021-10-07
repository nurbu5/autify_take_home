# syntax=docker/dockerfile:1

FROM golang:1.17.1-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
RUN mkdir ./fetcher
ADD fetcher ./fetcher

RUN go build -o /fetch

ENTRYPOINT [ "/fetch" ]
