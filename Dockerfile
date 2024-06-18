FROM golang:1.19-alpine AS builder

RUN apk add --no-cache build-base

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /cmd/api/main ./cmd/api

FROM alpine:latest

COPY --from=builder /cmd/api/main /app/main

WORKDIR /app

EXPOSE 8080

CMD ["/app/main"]