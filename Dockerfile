# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /nickscord-bot

##
## Deploy
##
FROM gcr.io/distroless/base-debian10 AS deploy

WORKDIR /

COPY --from=build /nickscord-bot /nickscord-bot

USER nonroot:nonroot

ENTRYPOINT ["/nickscord-bot"]