# ================================================================
#  Makefile  — place this at the REPO ROOT
#  Usage:
#    make build-auth       build authservice image
#    make build-gateway    build gateway image
#    make build-all        build both
#    make up               start infra + both services (test compose)
#    make down             stop everything
#    make logs             follow all logs
# ================================================================

IMAGE_PREFIX  := go-platform
TAG           := test

AUTH_IMAGE    := $(IMAGE_PREFIX)/authservice:$(TAG)
GATEWAY_IMAGE := $(IMAGE_PREFIX)/gateway:$(TAG)

COMPOSE_FILE  := docker-compose.test.yml

.PHONY: build-auth build-gateway build-all up down logs

## Build authservice — context is always repo root
build-auth:
	docker build \
		-f services/authservice/Dockerfile \
		-t $(AUTH_IMAGE) \
		.

## Build gateway — context is always repo root
build-gateway:
	docker build \
		-f services/gateway/Dockerfile \
		-t $(GATEWAY_IMAGE) \
		.

build-all: build-auth build-gateway

## Start infra + services (builds images if not present)
up:
	docker compose -f $(COMPOSE_FILE) up --build -d

## Stop and remove containers
down:
	docker compose -f $(COMPOSE_FILE) down

## Follow logs for all running containers
logs:
	docker compose -f $(COMPOSE_FILE) logs -f
