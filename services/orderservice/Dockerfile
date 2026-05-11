
ARG GO_VERSION=1.26.3
ARG ALPINE_VERSION=3.23

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS builder

ARG GOPROXY=https://go.devneeds.ir/,https://go.devneeds.ir/
ENV GOPROXY=$GOPROXY
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /src

#RUN #apk add --no-cache git

# Copy the whole monorepo so api/ and pkg/ from the root are available
COPY . .

# Build your gateway binary
# Adjust this path if your main package lives somewhere else
RUN go build -mod=vendor -o /out/orderservice ./services/orderservice/cmd


FROM alpine:3.23

#RUN #apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /out/orderservice /app/orderservice

COPY --from=builder /src/services/orderservice/internal/repository/mysql/migrations /app/migrations


COPY ./services/orderservice/cmd/config.yml /app/cmd/config.yml

EXPOSE 50062

ENTRYPOINT ["/app/orderservice"]