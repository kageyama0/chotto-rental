# Build Stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY .  /app

RUN CGO_ENABLED=0 GOOS=linux go build -v -o main ./cmd/api

# Runtime Stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/app

COPY --from=0 /app ./

EXPOSE 8080

CMD ["./main"]
