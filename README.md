# 🛒 Go Microservices Platform

A production-ready e-commerce backend built with Go, using a microservices architecture. Services communicate internally via **gRPC** and expose a single **REST API** through an API Gateway.

---

## 📋 Table of Contents

- [Architecture Overview](#architecture-overview)
- [Services](#services)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [API Endpoints](#api-endpoints)
- [Getting Started](#getting-started)
- [Configuration](#configuration)
- [Observability](#observability)

---

## Architecture Overview

```
                        ┌─────────────────────────────────────────┐
                        │              Client (HTTP)               │
                        └──────────────────┬──────────────────────┘
                                           │ :8080
                        ┌──────────────────▼──────────────────────┐
                        │              API Gateway                 │
                        │         (Echo · REST · JWT Auth)         │
                        └──┬──────┬───────┬──────┬────────────────┘
                     gRPC  │      │       │      │  gRPC
          ┌────────────────┘      │       │      └──────────────────┐
          │                       │       │                         │
   :50052 ▼                :50053 ▼  :50054 ▼                :50060 ▼
┌──────────────┐    ┌─────────────┐ ┌────────────┐    ┌──────────────────┐
│ userservice  │    │storeservice │ │productservice   │  cartservice      │
│   (MySQL)    │    │(MySQL+Redis)│ │(MySQL+Redis)│   │    (Redis)        │
└──────┬───────┘    └─────────────┘ └────────────┘    └────────┬─────────┘
       │                                                        │ publish
       │ gRPC :50051                                            │ RabbitMQ
       ▼                                                        ▼
┌──────────────┐                                    ┌──────────────────────┐
│ authservice  │                                    │    orderservice      │
│  (JWT only)  │                                    │  (MySQL · consumer)  │
└──────────────┘                                    └──────────────────────┘
```

**Flow:** All HTTP traffic enters through the Gateway. The Gateway validates JWT tokens via `authservice` and routes requests to the appropriate backend service over gRPC. When a user checks out, `cartservice` publishes a `CartCheckedOutEvent` to RabbitMQ, and `orderservice` consumes it asynchronously to create the order.

---

## Services

| Service | Port | Protocol | Responsibility |
|---|---|---|---|
| **gateway** | `8080` | HTTP/REST | Single entry point; JWT validation; routes to backends |
| **authservice** | `50051` | gRPC | Issues and validates JWT access/refresh tokens |
| **userservice** | `50052` | gRPC | User registration, login, profile; stores in MySQL |
| **storeservice** | `50053` | gRPC | CRUD for stores; MySQL + Redis cache |
| **productservice** | `50054` | gRPC | CRUD for products; MySQL + Redis cache |
| **cartservice** | `50060` | gRPC | Manages user carts in Redis; publishes checkout events |
| **orderservice** | `50062` | gRPC + MQ | Consumes checkout events from RabbitMQ; stores orders |

---

## Tech Stack

| Category | Technology |
|---|---|
| Language | Go 1.22+ |
| HTTP Framework | [Echo v4](https://echo.labstack.com/) |
| Service Communication | gRPC + Protocol Buffers |
| Primary Database | MySQL 8 |
| Cache | Redis 7 |
| Message Broker | RabbitMQ 3.13 |
| Authentication | JWT (HS256) |
| Logger | [Zap](https://github.com/uber-go/zap) + Lumberjack (log rotation) |
| Config | [Koanf](https://github.com/knadh/koanf) (YAML + ENV override) |
| Metrics | Prometheus + Grafana |
| Containerization | Docker + Docker Compose |

---

## Project Structure

```
.
├── api/
│   ├── proto/          # Protobuf definitions (.proto files)
│   └── gen/            # Auto-generated gRPC code
├── pkg/
│   ├── config/         # Config loader (koanf)
│   ├── richerror/      # Domain error type (Kind, Op, Message)
│   ├── grpcerror/      # gRPC ↔ RichError mapping
│   ├── httpmsg/        # HTTP error response mapper
│   ├── interceptor/    # gRPC server error interceptor
│   └── logger/         # Global zap logger setup
├── services/
│   ├── gateway/        # HTTP API Gateway (Echo)
│   ├── authservice/    # JWT token service (gRPC)
│   ├── userservice/    # User management (gRPC + MySQL)
│   ├── productservice/ # Product catalog (gRPC + MySQL + Redis)
│   ├── storeservice/   # Store management (gRPC + MySQL + Redis)
│   ├── cartservice/    # Shopping cart (gRPC + Redis + RabbitMQ)
│   └── orderservice/   # Order processing (RabbitMQ consumer + MySQL)
├── docker-compose.yml
└── prometheus.yml
```

Each service follows a **clean architecture** layout:

```
services/<name>/
├── cmd/
│   ├── main.go         # Entry point
│   └── config.yml      # Service config
└── internal/
    ├── config/         # Config struct
    ├── domain/         # Core entities (no dependencies)
    ├── param/          # Request/Response structs
    ├── validator/      # Input validation (ozzo-validation)
    ├── service/        # Business logic
    ├── repository/     # Data access (MySQL / Redis)
    └── delivery/       # Transport layer (gRPC / HTTP / broker)
```

---

## API Endpoints

Base URL: `http://localhost:8080`

> Endpoints marked with 🔒 require an `Authorization: Bearer <token>` header.

### Auth & Users

| Method | Path | Auth | Description |
|---|---|---|---|
| `POST` | `/users/register` | — | Register a new user |
| `POST` | `/users/login` | — | Login and receive JWT tokens |
| `GET` | `/users/profile` | 🔒 | Get current user profile |

**Register request body:**
```json
{
  "name": "Ali Rezaei",
  "phone_number": "09121234567",
  "password": "secret123"
}
```

**Login request body:**
```json
{
  "phone_number": "09121234567",
  "password": "secret123"
}
```

---

### Products

| Method | Path | Auth | Description |
|---|---|---|---|
| `POST` | `/product/create` | 🔒 | Create a new product |
| `GET` | `/product/get` | 🔒 | Get product by ID |
| `POST` | `/product/update` | 🔒 | Update product details |
| `POST` | `/product/delete` | 🔒 | Delete a product |

---

### Stores

| Method | Path | Auth | Description |
|---|---|---|---|
| `POST` | `/store/create` | 🔒 | Create a new store |
| `GET` | `/store/list` | 🔒 | List stores for current user |
| `POST` | `/store/update` | 🔒 | Update store details |
| `POST` | `/store/delete` | 🔒 | Delete a store |

---

### Cart

| Method | Path | Auth | Description |
|---|---|---|---|
| `POST` | `/cart/add_item` | 🔒 | Add item to cart |
| `GET` | `/cart/get_cart` | 🔒 | View current cart |
| `GET` | `/cart/check_out` | 🔒 | Checkout — publishes event to RabbitMQ |

---

### System

| Method | Path | Description |
|---|---|---|
| `GET` | `/health-check` | Gateway health status |
| `GET` | `/metrics` | Prometheus metrics (gateway) |

---

## Getting Started

### Prerequisites

- [Docker](https://www.docker.com/) ≥ 24
- [Docker Compose](https://docs.docker.com/compose/) v2

### 1. Clone the repository

```bash
git clone https://github.com/ipouyakeshavarz-ctrl/microservice.git
cd microservice
```

### 2. Start all services

```bash
docker compose up --build
```

This command starts:
- All 7 Go microservices
- MySQL 8 (port `3311` on host)
- Redis 7 (port `6379`)
- RabbitMQ 3.13 (port `5672`, management UI on `15672`)
- Prometheus (port `9090`)
- Grafana (port `3000`)

### 3. Verify everything is running

```bash
curl http://localhost:8080/health-check
```

### 4. Run a quick test

```bash
# Register a user
curl -X POST http://localhost:8080/users/register \
  -H "Content-Type: application/json" \
  -d '{"name":"Test User","phone_number":"09121234567","password":"test1234"}'

# Login
curl -X POST http://localhost:8080/users/login \
  -H "Content-Type: application/json" \
  -d '{"phone_number":"09121234567","password":"test1234"}'
```

### Stop services

```bash
docker compose down          # stop containers
docker compose down -v       # stop and remove all volumes (wipes data)
```

---

## Configuration

Each service reads its config from `cmd/config.yml`. Values can be overridden with environment variables (prefix per service).

**Example — `services/gateway/cmd/config.yml`:**

```yaml
http_server:
  address: "0.0.0.0:8080"

grpc_client:
  auth_address:    "authservice:50051"
  user_address:    "userservice:50052"
  store_address:   "storeservice:50053"
  product_address: "productservice:50054"
  cart_address:    "cartservice:50060"

logger:
  development: true
  service_name: "gateway"
  file_path: "./logs/gateway.log"
```

> **Note:** In production, move secrets (DB passwords, JWT secret key) to environment variables or a secrets manager. Never commit credentials to git.

---

## Observability

### Prometheus

Metrics are exposed on port `9100` of every service. Prometheus scrapes all of them.

Access: [http://localhost:9090](http://localhost:9090)

### Grafana

Pre-configured to use Prometheus as a data source.

Access: [http://localhost:3000](http://localhost:3000)
Default credentials: `admin` / `admin`

### RabbitMQ Management UI

Monitor queues, exchanges, and message rates.

Access: [http://localhost:15672](http://localhost:15672)
Default credentials: `user` / `password`

### Structured Logging

All services use **Zap** for structured JSON logging with log rotation via Lumberjack. Logs are written to the path defined in each service's `config.yml`.

---

## Internal Design Patterns

### RichError

All services use a custom `RichError` type (`pkg/richerror`) that carries:
- `Op` — the operation where the error occurred
- `Kind` — error category (`Invalid`, `Forbidden`, `NotFound`, `Unexpected`)
- `Message` — human-readable message
- `Meta` / `Fields` — extra context

This allows the gRPC interceptor and HTTP mapper to automatically convert domain errors into the correct HTTP status codes and gRPC status codes without any manual mapping in handlers.

### Async Checkout Flow

```
cartservice               RabbitMQ             orderservice
     │                       │                      │
     │── checkout request ──▶│                      │
     │── CartCheckedOutEvent ▶│                      │
     │                       │── consume event ─────▶│
     │                       │                      │── CreateOrder in MySQL
```

---

## License

MIT
