FROM golang:1.23.3-alpine AS builder

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o worker-service ./cmd/worker

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/worker-service .

CMD ["./worker-service"]