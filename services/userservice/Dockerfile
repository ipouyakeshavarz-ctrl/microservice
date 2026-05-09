FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY services/userservice/go.mod .
COPY services/userservice/go.sum .
COPY services/userservice/vendor ./vendor

COPY services/userservice/cmd ./cmd
COPY services/userservice/internal ./internal
COPY pkg ./pkg

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -mod=vendor -o app ./cmd

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/app .
CMD ["./app"]

