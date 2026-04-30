FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

COPY vendor ./vendor


COPY . .

RUN go build -mod=vendor -o main ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

COPY --from=builder /app/cmd/config.yml ./cmd/config.yml

CMD ["./main"]
