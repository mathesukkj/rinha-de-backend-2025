# build
FROM golang:1.24.1-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./src/main.go

# run
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main /app/main
CMD ["./main"]