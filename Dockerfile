# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY api/*.go ./

RUN go build -o /go-server

EXPOSE 8090

CMD [ "/go-server" ]