FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY main.go .
COPY /data /data

# Explicitly build for linux/amd64
RUN GOOS=linux GOARCH=amd64 go build -o server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .
COPY /data /data

EXPOSE 8080
CMD ["./server"]