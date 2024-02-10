FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY cmd/main.go .

RUN go mod download
RUN go build -o bin .

ENTRYPOINT [ "/app/bin" ]