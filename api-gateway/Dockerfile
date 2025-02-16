FROM golang:1.23.3-alpine AS builder

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o gateway-service ./cmd/gateway-service

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/gateway-service .

EXPOSE 8080

CMD ["./gateway-service"]