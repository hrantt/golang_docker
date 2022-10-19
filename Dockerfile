# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR .
RUN mkdir -p /app
COPY /src/greetings.go /app
copy /src/go.mod /app
WORKDIR /app
RUN go build -o /greetings
CMD ["/greetings"]
