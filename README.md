# 🚀 Boilerplate Go - Clean Layered Architecture API

Production-ready Go REST API boilerplate dengan Clean Architecture, Gin Framework, PostgreSQL, GORM, dan best practices modern.

---

## 📖 Table of Contents

- [Quick Start (5 Menit)](#-quick-start-5-menit)
- [Teknologi yang Digunakan](#-teknologi-yang-digunakan)
- [Prerequisites](#-prerequisites)
- [Installation & Setup](#-installation--setup)
- [Project Structure](#-project-structure)
- [Clean Layered Architecture](#-clean-layered-architecture)
- [Configuration](#-configuration)
- [API Endpoints](#-api-endpoints)
- [API Testing Guide](#-api-testing-guide)
- [Database Setup](#-database-setup)
- [Docker Setup](#-docker-setup)
- [Makefile Commands](#-makefile-commands)
- [Adding New Endpoints](#-adding-new-endpoints)
- [Architecture Decisions](#-architecture-decisions)
- [Production Deployment](#-production-deployment)
- [Contributing Guidelines](#-contributing-guidelines)
- [Best Practices](#-best-practices)
- [Troubleshooting](#-troubleshooting)

---

## ⚡ Quick Start (5 Menit)

### Step 1: Prerequisites Check
```bash
go version              # Go 1.22 or higher
docker --version       # Docker (optional)
```

### Step 2: Clone & Setup
```bash
git clone <repository-url>
cd boilerplate-go
cp .env.example .env
make setup
```

### Step 3: Start Services
```bash
# Option A: Docker (Recommended)
docker-compose up -d

# Option B: Local PostgreSQL
make migrate-up
```

### Step 4: Run Application
```bash
make dev                # With hot reload
# atau
make run               # Without hot reload
```

### Step 5: Test API
```bash
# Health check
curl http://localhost:8080/health

# Create user
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Test","email":"test@example.com","phone":"08123456789"}'
```

✅ **Aplikasi siap development!**

---

## 🛠 Teknologi yang Digunakan

| Technology | Version | Purpose |
|-----------|---------|---------|
| **Go** | 1.22+ | Language |
| **Gin** | v1.9.0 | HTTP Web Framework |
| **PostgreSQL** | 14+ | Database |
| **GORM** | v1.25.0 | ORM |
| **Zap** | v1.24.0 | Structured Logging |
| **UUID** | v1.3.0 | Unique Identifiers |
| **godotenv** | v1.5.1 | Environment Variables |
| **Docker** | Latest | Containerization |
| **golang-migrate** | Latest | Database Migrations |

---

## 📦 Prerequisites

### Required
- **Go 1.22** atau lebih baru
- **PostgreSQL 14** atau lebih baru
- **Git**

### Optional
- **Docker** & **Docker Compose** (untuk containerized development)
- **Make** (untuk automation commands)

### Install Go
```bash
# macOS dengan Homebrew
brew install go

# Linux (Ubuntu/Debian)
sudo apt-get install golang-go

# Windows
# Download dari https://golang.org/dl
```

### Verifikasi Instalasi
```bash
go version      # Should show go1.22 or higher
echo $GOPATH    # Should show your Go workspace
```

---

## 🔧 Installation & Setup

### Step 1: Clone Repository
```bash
git clone <repository-url>
cd boilerplate-go
```

### Step 2: Install Dependencies
```bash
go mod download
go mod tidy
```

Atau gunakan Makefile:
```bash
make setup
```

### Step 3: Environment Configuration
```bash
# Copy template
cp .env.example .env

# Edit dengan nilai sesuai environment Anda
nano .env  # atau gunakan editor pilihan Anda
```

### Step 4: Setup Database

**Option A: Docker Compose (Recommended)**
```bash
docker-compose up -d          # Start PostgreSQL
make migrate-up               # Run migrations
```

**Option B: Local PostgreSQL**
```bash
# Buat database
createdb boilerplate_go

# Run migrations
make migrate-up
```

### Step 5: Run Application
```bash
# Development dengan hot reload
make dev

# atau tanpa hot reload
make run

# atau build binary
make build
./main
```

Aplikasi berjalan di: **http://localhost:8080**

---

## 📁 Project Structure

```
boilerplate-go/
├── 📄 Documentation & Config
│   ├── README.md                    ← Dokumentasi lengkap (file ini)
│   ├── .env                         ← Environment variables
│   ├── .env.example                 ← Template environment
│   ├── go.mod & go.sum              ← Dependencies
│   └── .gitignore
│
├── 🔧 Build & Configuration
│   ├── Makefile                     ← Task automation
│   ├── Dockerfile                   ← Docker image
│   ├── docker-compose.yml           ← Development stack
│   ├── .air.toml                    ← Hot reload config
│   └── setup.sh / setup.bat         ← Auto setup scripts
│
├── 💻 Source Code
│   ├── cmd/api/
│   │   └── main.go                  ← Entry point (90 lines)
│   ├── internal/                    ← Private packages
│   │   ├── config/
│   │   │   ├── config.go            ← Load environment (50 lines)
│   │   │   └── database.go          ← GORM initialization (40 lines)
│   │   ├── handler/
│   │   │   └── user_handler.go      ← HTTP handlers (90 lines)
│   │   ├── middleware/
│   │   │   ├── cors.go              ← CORS config (20 lines)
│   │   │   ├── error_handler.go     ← Error handling (30 lines)
│   │   │   └── logger.go            ← Request logging (25 lines)
│   │   ├── model/
│   │   │   └── user.go              ← Domain models & DTOs (60 lines)
│   │   ├── repository/
│   │   │   └── user_repository.go   ← Database queries (100 lines)
│   │   └── service/
│   │       └── user_service.go      ← Business logic (90 lines)
│   └── pkg/                         ← Public packages
│       └── logger/
│           └── logger.go            ← Zap wrapper (60 lines)
│
├── 🗄️ Database
│   └── migrations/
│       ├── 000001_create_users_table.up.sql
│       └── 000001_create_users_table.down.sql
│
└── .git/                            ← Git repository
```

### File Statistics
- **Documentation**: 1 file (comprehensive)
- **Source Code**: 12 files (550+ lines)
- **Configuration**: 9 files
- **Setup Scripts**: 2 files
- **Migrations**: 2 files

---

## 🏗️ Clean Layered Architecture

### Arsitektur Overview

```
┌──────────────────────────────────────────────┐
│        HTTP Client (Browser, Mobile, etc)   │
└─────────────────────┬──────────────────────┘
                      │
┌─────────────────────▼──────────────────────┐
│    Handler Layer (HTTP Routing)             │ ← cmd/api/main.go
│  Validate → Call Service → Format Response │   internal/handler/*
└─────────────────────┬──────────────────────┘
                      │
┌─────────────────────▼──────────────────────┐
│   Service Layer (Business Logic)            │ ← internal/service/*
│  Validation → Processing → Rules           │
└─────────────────────┬──────────────────────┘
                      │
┌─────────────────────▼──────────────────────┐
│  Repository Layer (Data Access)             │ ← internal/repository/*
│         CRUD Queries via GORM               │
└─────────────────────┬──────────────────────┘
                      │
┌─────────────────────▼──────────────────────┐
│       Database Layer (PostgreSQL)           │
│    Tables, Indexes, Constraints             │
└──────────────────────────────────────────────┘
```

### Request Flow Example

```
1. Client Request
   POST /api/v1/users
   {"name": "John", "email": "john@example.com"}

2. Handler Layer (user_handler.go)
   ✓ Parse JSON
   ✓ Validate fields
   ✓ Call userService.CreateUser()

3. Service Layer (user_service.go)
   ✓ Check email uniqueness
   ✓ Generate UUID
   ✓ Call userRepository.Create()

4. Repository Layer (user_repository.go)
   ✓ Build GORM query
   ✓ Execute INSERT
   ✓ Return created user

5. Response
   201 Created
   {"id": "550e8400-...", "name": "John", ...}
```

### Layer Responsibilities

| Layer | File | Responsibility |
|-------|------|-----------------|
| **Handler** | `internal/handler/` | HTTP routing, request validation, response formatting |
| **Service** | `internal/service/` | Business logic, data validation, domain rules |
| **Repository** | `internal/repository/` | Database queries, entity mapping, CRUD operations |
| **Model** | `internal/model/` | Domain entities, request DTOs, response DTOs |
| **Middleware** | `internal/middleware/` | CORS, error handling, request logging |
| **Config** | `internal/config/` | Environment loading, database initialization |
| **Logger** | `pkg/logger/` | Structured logging (reusable) |

---

## ⚙️ Configuration

### Environment Variables (.env)

```env
# Application
APP_NAME=boilerplate-go              # Nama aplikasi
APP_ENV=development                 # Environment: development/production
APP_PORT=8080                       # HTTP port

# Database
DB_HOST=localhost                   # PostgreSQL host
DB_PORT=5432                        # PostgreSQL port
DB_USER=postgres                    # Database user
DB_PASSWORD=postgres                # Database password
DB_NAME=boilerplate_go              # Database name
DB_SSL_MODE=disable                 # SSL mode: disable/require/verify-ca/verify-full

# Server
SERVER_TIMEOUT=30                   # Request timeout (seconds)

# CORS
CORS_ALLOW_ORIGINS=*               # Allowed origins (* atau comma-separated)
```

### Database Connection String

GORM akan otomatis membuat connection string dari environment variables:

```go
dsn := fmt.Sprintf(
    "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
    cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode,
)
```

### Connection Pool Configuration

```go
// internal/config/database.go
sqlDB.SetMaxIdleConns(10)      // Min idle connections
sqlDB.SetMaxOpenConns(100)     // Max open connections
```

Untuk production, sesuaikan nilai berdasarkan traffic:
- Low traffic: MaxIdleConns=5, MaxOpenConns=20
- Medium traffic: MaxIdleConns=10, MaxOpenConns=50
- High traffic: MaxIdleConns=20, MaxOpenConns=100+

---

## 📡 API Endpoints

### Base URL
```
http://localhost:8080/api/v1
```

### Endpoints

#### Health Check
```http
GET /health
```
Response: `{"status":"OK","time":"2024-01-15T10:30:45Z"}`

#### Users - Create
```http
POST /api/v1/users
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "phone": "08123456789"
}
```
Response (201 Created):
```json
{
  "message": "User created successfully",
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "John Doe",
    "email": "john@example.com",
    "phone": "08123456789",
    "created_at": "2024-01-15T10:30:45Z",
    "updated_at": "2024-01-15T10:30:45Z"
  }
}
```

#### Users - Get All
```http
GET /api/v1/users?limit=10&offset=0
```
Response (200 OK):
```json
{
  "message": "Users retrieved successfully",
  "data": [
    {
      "id": "550e8400-...",
      "name": "John Doe",
      "email": "john@example.com",
      "phone": "08123456789",
      "created_at": "2024-01-15T10:30:45Z",
      "updated_at": "2024-01-15T10:30:45Z"
    }
  ],
  "pagination": {
    "limit": 10,
    "offset": 0,
    "total": 1
  }
}
```

#### Users - Get by ID
```http
GET /api/v1/users/550e8400-e29b-41d4-a716-446655440000
```
Response (200 OK): User object

#### Users - Update
```http
PUT /api/v1/users/550e8400-e29b-41d4-a716-446655440000
Content-Type: application/json

{
  "name": "Jane Doe",
  "email": "jane@example.com",
  "phone": "08987654321"
}
```
Response (200 OK): Updated user object

#### Users - Delete
```http
DELETE /api/v1/users/550e8400-e29b-41d4-a716-446655440000
```
Response (204 No Content)

### Error Response

Semua error akan return dalam format:
```json
{
  "error": {
    "code": 400,
    "message": "Invalid request: email is required"
  }
}
```

Common HTTP Status Codes:
- `200 OK` - Success
- `201 Created` - Resource created
- `204 No Content` - Delete success
- `400 Bad Request` - Invalid input
- `404 Not Found` - Resource not found
- `500 Internal Server Error` - Server error

---

## 🧪 API Testing Guide

### Prerequisites
- curl (built-in macOS/Linux)
- atau Postman: https://www.postman.com/downloads/

### cURL Examples

#### 1. Health Check
```bash
curl -X GET http://localhost:8080/health
```

#### 2. Create User
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "phone": "08123456789"
  }'
```

#### 3. Get All Users
```bash
curl -X GET "http://localhost:8080/api/v1/users?limit=10&offset=0"
```

#### 4. Get User by ID
```bash
curl -X GET http://localhost:8080/api/v1/users/550e8400-e29b-41d4-a716-446655440000
```

#### 5. Update User
```bash
curl -X PUT http://localhost:8080/api/v1/users/550e8400-e29b-41d4-a716-446655440000 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Jane Doe",
    "email": "jane@example.com",
    "phone": "08987654321"
  }'
```

#### 6. Delete User
```bash
curl -X DELETE http://localhost:8080/api/v1/users/550e8400-e29b-41d4-a716-446655440000
```

### Batch Testing Script

Simpan sebagai `test-api.sh`:

```bash
#!/bin/bash

BASE_URL="http://localhost:8080/api/v1"

echo "=== Testing User API ==="

# Create User
echo -e "\n1. Creating user..."
RESPONSE=$(curl -s -X POST $BASE_URL/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test User",
    "email": "test@example.com",
    "phone": "081234567890"
  }')

echo "$RESPONSE" | jq .
USER_ID=$(echo "$RESPONSE" | jq -r '.data.id')
echo "Created user ID: $USER_ID"

# Get All Users
echo -e "\n2. Getting all users..."
curl -s -X GET "$BASE_URL/users?limit=10&offset=0" | jq .

# Get User by ID
echo -e "\n3. Getting user by ID..."
curl -s -X GET "$BASE_URL/users/$USER_ID" | jq .

# Update User
echo -e "\n4. Updating user..."
curl -s -X PUT "$BASE_URL/users/$USER_ID" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Updated User",
    "phone": "089876543210"
  }' | jq .

# Delete User
echo -e "\n5. Deleting user..."
curl -s -X DELETE "$BASE_URL/users/$USER_ID" | jq .

echo -e "\n=== Testing Complete ==="
```

Run:
```bash
chmod +x test-api.sh
./test-api.sh
```

### Postman Setup

1. Create Collection: "Boilerplate Go"
2. Add Environment variable: `base_url = http://localhost:8080`
3. Import requests dari struktur API di atas
4. Use "Tests" untuk auto-capture user_id:

```javascript
if (pm.response.code === 201) {
  var jsonData = pm.response.json();
  pm.environment.set("user_id", jsonData.data.id);
}
```

### Load Testing

```bash
# Install Apache Bench
# macOS: brew install httpd
# Linux: sudo apt-get install apache2-utils

# Test dengan 1000 requests, 100 concurrent
ab -n 1000 -c 100 http://localhost:8080/health
```

---

## 🗄️ Database Setup

### Create Tables (Migration Up)

File: `migrations/000001_create_users_table.up.sql`

```sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_created_at ON users(created_at);
```

### Drop Tables (Migration Down)

File: `migrations/000001_create_users_table.down.sql`

```sql
DROP INDEX IF EXISTS idx_users_created_at;
DROP INDEX IF EXISTS idx_users_email;
DROP TABLE IF EXISTS users;
```

### Running Migrations

```bash
# Run all pending migrations
make migrate-up

# Rollback last migration
make migrate-down

# Go to specific version
migrate -path ./migrations -database "$DATABASE_URL" goto 1

# Check current version
migrate -path ./migrations -database "$DATABASE_URL" version
```

---

## 🐳 Docker Setup

### Docker Compose (Recommended)

```bash
# Start services
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down

# Remove volumes (cleanup database)
docker-compose down -v
```

### Services dalam docker-compose.yml

1. **PostgreSQL**
   - Image: postgres:16-alpine
   - Port: 5432
   - Database: boilerplate_go
   - User: postgres
   - Password: postgres

2. **Application** (optional dalam docker-compose)
   - Build dari Dockerfile
   - Port: 8080

### Build Docker Image

```bash
# Build dengan tag
docker build -t boilerplate-go:1.0.0 .

# Run container
docker run -p 8080:8080 boilerplate-go:1.0.0

# Run dengan environment variables
docker run -p 8080:8080 \
  -e APP_ENV=production \
  -e DB_HOST=host.docker.internal \
  boilerplate-go:1.0.0
```

### Multi-stage Build (Production)

Dockerfile sudah menggunakan multi-stage build untuk optimize image size:

```dockerfile
# Stage 1: Build
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main ./cmd/api

# Stage 2: Runtime
FROM alpine:latest
COPY --from=builder /app/main .
CMD ["./main"]
```

Hasil: Image size ~15MB (vs 800MB+ jika tanpa multi-stage)

---

## 🔨 Makefile Commands

### Setup & Installation
```bash
make setup              # Install dependencies
make clean              # Remove build artifacts
```

### Development
```bash
make dev                # Run with hot reload (Air)
make run                # Run once
make build              # Build binary
```

### Testing
```bash
make test               # Run all tests
make test-coverage      # Generate coverage report
make lint               # Lint code
make fmt                # Format code
```

### Database
```bash
make migrate-up         # Run migrations
make migrate-down       # Rollback migrations
```

### Docker
```bash
make docker-build       # Build Docker image
make docker-run         # Run in Docker
```

### Help
```bash
make help               # Show all commands
```

---

## 🆕 Adding New Endpoints

### Example: Create Product Endpoint

#### 1. Add Model (internal/model/product.go)

```go
package model

import "time"

type Product struct {
    ID        string    `json:"id"`
    Name      string    `json:"name"`
    Price     float64   `json:"price"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type CreateProductRequest struct {
    Name  string  `json:"name" binding:"required"`
    Price float64 `json:"price" binding:"required,gt=0"`
}
```

#### 2. Add Repository (internal/repository/product_repository.go)

```go
package repository

import (
    "boilerplate-go/internal/model"
    "gorm.io/gorm"
)

type ProductRepository interface {
    Create(product *model.Product) error
    GetByID(id string) (*model.Product, error)
    GetAll(limit, offset int) ([]model.Product, int64, error)
}

type productRepository struct {
    db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
    return &productRepository{db: db}
}

func (r *productRepository) Create(product *model.Product) error {
    return r.db.Create(product).Error
}

func (r *productRepository) GetByID(id string) (*model.Product, error) {
    var product model.Product
    err := r.db.First(&product, "id = ?", id).Error
    return &product, err
}

func (r *productRepository) GetAll(limit, offset int) ([]model.Product, int64, error) {
    var products []model.Product
    var total int64
    err := r.db.Model(&model.Product{}).Count(&total).
        Limit(limit).
        Offset(offset).
        Find(&products).Error
    return products, total, err
}
```

#### 3. Add Service (internal/service/product_service.go)

```go
package service

import (
    "fmt"
    "boilerplate-go/internal/model"
    "boilerplate-go/internal/repository"
    "github.com/google/uuid"
)

type ProductService interface {
    CreateProduct(req *model.CreateProductRequest) (*model.Product, error)
    GetProductByID(id string) (*model.Product, error)
}

type productService struct {
    repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
    return &productService{repo: repo}
}

func (s *productService) CreateProduct(req *model.CreateProductRequest) (*model.Product, error) {
    product := &model.Product{
        ID:    uuid.New().String(),
        Name:  req.Name,
        Price: req.Price,
    }
    if err := s.repo.Create(product); err != nil {
        return nil, fmt.Errorf("failed to create product: %w", err)
    }
    return product, nil
}

func (s *productService) GetProductByID(id string) (*model.Product, error) {
    product, err := s.repo.GetByID(id)
    if err != nil {
        return nil, fmt.Errorf("failed to get product: %w", err)
    }
    return product, nil
}
```

#### 4. Add Handler (internal/handler/product_handler.go)

```go
package handler

import (
    "net/http"
    "boilerplate-go/internal/model"
    "boilerplate-go/internal/service"
    "boilerplate-go/pkg/logger"
    "github.com/gin-gonic/gin"
)

type ProductHandler struct {
    service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
    return &ProductHandler{service: service}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
    var req model.CreateProductRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    product, err := h.service.CreateProduct(&req)
    if err != nil {
        logger.Error("Failed to create product", zap.Error(err))
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Product created successfully",
        "data":    product,
    })
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
    id := c.Param("id")
    product, err := h.service.GetProductByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Product retrieved successfully",
        "data":    product,
    })
}
```

#### 5. Register Routes (cmd/api/main.go)

```go
func setupRouter(cfg *config.Config, userHandler *handler.UserHandler, 
    productHandler *handler.ProductHandler) *gin.Engine {
    router := gin.Default()
    
    // Middleware
    router.Use(middleware.CORSMiddleware(cfg.CORSAllowOrigins))
    router.Use(middleware.LoggerMiddleware())
    router.Use(middleware.ErrorHandlerMiddleware())
    
    // Health
    router.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "OK"})
    })
    
    // User routes
    api := router.Group("/api/v1")
    api.POST("/users", userHandler.CreateUser)
    api.GET("/users", userHandler.GetAllUsers)
    
    // Product routes
    api.POST("/products", productHandler.CreateProduct)
    api.GET("/products/:id", productHandler.GetProductByID)
    
    return router
}
```

#### 6. Create Migration

File: `migrations/000002_create_products_table.up.sql`

```sql
CREATE TABLE products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_products_name ON products(name);
```

---

## 🏛️ Architecture Decisions

### ADR-001: Clean Layered Architecture

**Decision**: Use 4-layer clean architecture (Handler → Service → Repository → Model)

**Rationale**:
- Separation of concerns yang jelas
- Easy testability dengan interface-based design
- Scalable dan maintainable
- Follows SOLID principles

---

### ADR-002: GORM sebagai ORM

**Decision**: Use GORM instead of raw SQL queries

**Rationale**:
- Rapid development
- Type-safe queries
- Built-in migration support
- Good performance dengan proper indexing

**Alternatives**: sqlc, sqlx, raw SQL

---

### ADR-003: UUID sebagai Primary Key

**Decision**: Use UUID v4 strings as primary keys

**Rationale**:
- Globally unique
- Privacy (tidak reveal record count)
- Portability across databases

**Trade-off**: Larger storage (36 bytes vs 8 bytes), slower indexing

---

### ADR-004: Interface-based Repository Pattern

**Decision**: Define repository interfaces

**Rationale**:
- Easy to mock untuk testing
- Can swap implementations
- Dependency inversion

---

### ADR-005: Environment-based Configuration

**Decision**: Load config dari environment variables

**Rationale**:
- 12 Factor App compliance
- Security (secrets in env, not in code)
- Flexibility per environment

---

### ADR-006: Centralized Error Handling

**Decision**: Global error handler middleware + structured logging

**Rationale**:
- Consistent error responses
- Easy to debug dan monitor
- Single point for error formatting

---

## 🚀 Production Deployment

### Pre-Deployment Checklist

#### Code Quality
- [ ] All tests passed: `make test`
- [ ] Code coverage >80%: `make test-coverage`
- [ ] No linting errors: `make lint`
- [ ] Code formatted: `make fmt`

#### Configuration
- [ ] `.env.production` created dengan values yang correct
- [ ] Database credentials updated
- [ ] `APP_ENV=production` set
- [ ] `DB_SSL_MODE=require`
- [ ] CORS_ALLOW_ORIGINS diupdate

#### Database
- [ ] Database created
- [ ] Migrations run: `migrate -path ./migrations -database "<DB_URL>" up`
- [ ] Indexes created
- [ ] Backup strategy configured

#### Security
- [ ] HTTPS/TLS certificate configured
- [ ] Input validation in place
- [ ] Rate limiting implemented
- [ ] CORS properly configured

#### Monitoring
- [ ] Logging to aggregation tool (ELK, Datadog, etc)
- [ ] Health endpoint accessible
- [ ] Metrics collection (Prometheus, etc)
- [ ] Alert rules configured

### Deployment Steps

#### Step 1: Build Binary
```bash
make build
```

#### Step 2: Run Migrations
```bash
migrate -path ./migrations -database "$DATABASE_URL" up
```

#### Step 3: Start Application
```bash
# Using systemd
sudo systemctl start boilerplate-go

# Or using supervisor
supervisorctl start boilerplate-go

# Or in Docker
docker-compose -f docker-compose.prod.yml up -d
```

#### Step 4: Verify
```bash
curl https://api.yourdomain.com/health
```

### Scaling Considerations

1. **Database Connection Pooling**
   ```go
   // For high traffic
   sqlDB.SetMaxOpenConns(200)
   ```

2. **Load Balancing**
   - Deploy multiple instances behind Nginx/HAProxy
   - Health check endpoint: `/health`

3. **Caching**
   - Add Redis untuk session/cache
   - Cache frequently accessed data

4. **Monitoring**
   - Setup error tracking (Sentry, etc)
   - Monitor database connections
   - Monitor memory usage

---

## 👥 Contributing Guidelines

### How to Contribute

#### 1. Fork Repository
```bash
# Click "Fork" on GitHub
```

#### 2. Clone Your Fork
```bash
git clone https://github.com/YOUR_USERNAME/boilerplate-go.git
cd boilerplate-go
```

#### 3. Create Feature Branch
```bash
git checkout -b feature/amazing-feature
# atau
git checkout -b fix/bug-fix
```

#### 4. Make Changes
- Write code following project conventions
- Add tests untuk new features
- Update documentation if needed

#### 5. Commit Changes
```bash
git add .
git commit -m "feat: Add amazing feature"
```

**Commit message format**:
- `feat:` untuk fitur baru
- `fix:` untuk bug fixes
- `docs:` untuk dokumentasi
- `refactor:` untuk refactoring
- `test:` untuk test additions

#### 6. Push & Create Pull Request
```bash
git push origin feature/amazing-feature
```

### Code Conventions

- Follow Go conventions (gofmt)
- Use meaningful variable names
- Add comments untuk complex logic
- Maximum line length: 100 characters
- Write tests untuk coverage >80%

### Testing Before Submit
```bash
make fmt                # Format code
make lint               # Check linting
make test               # Run tests
make test-coverage      # Check coverage
```

---

## 💡 Best Practices

### 1. Error Handling
```go
// ✅ Good: Wrap dengan context
if err := repo.Create(user); err != nil {
    return fmt.Errorf("failed to create user: %w", err)
}

// ❌ Bad: Ignore error
repo.Create(user)
```

### 2. Logging
```go
// ✅ Good: Structured logging
logger.Info("User created", zap.String("email", user.Email))

// ❌ Bad: String formatting
log.Printf("User created: " + user.Email)
```

### 3. Validation
```go
// ✅ Good: Validate early
if err := c.ShouldBindJSON(&req); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
}

// ❌ Bad: Skip validation
data := c.GetString("email") // Tidak safe
```

### 4. Connection Management
```go
// ✅ Good: Use connection pool
sqlDB.SetMaxOpenConns(100)
sqlDB.SetMaxIdleConns(10)

// ❌ Bad: Create new connection per request
db := sql.Open(...)  // Jangan di handler
```

### 5. API Response Format
```go
// ✅ Good: Consistent format
c.JSON(http.StatusOK, gin.H{
    "message": "User created",
    "data": user,
})

// ❌ Bad: Inconsistent format
c.JSON(http.StatusOK, user)
```

### 6. Database Queries
```go
// ✅ Good: Use parametrized queries via GORM
user, err := repo.GetByEmail(email)

// ❌ Bad: String interpolation
db.Where("email = '" + email + "'").First(&user)  // SQL Injection!
```

### 7. Transaction Handling
```go
// ✅ Good: Use transactions untuk consistency
tx := db.BeginTx(ctx, nil)
defer tx.Rollback()
// ... operations ...
tx.Commit()

// ❌ Bad: No transaction
db.Create(...).Update(...).Delete(...)
```

### 8. Graceful Shutdown
```go
// ✅ Good: Handle shutdown signal
sigChan := make(chan os.Signal, 1)
signal.Notify(sigChan, syscall.SIGTERM)
<-sigChan
server.Shutdown(ctx)

// ❌ Bad: Immediate exit
os.Exit(0)
```

---

## 🆘 Troubleshooting

### Build Errors

#### Error: "missing go.sum entry"
```
Solution: Run go mod tidy
go mod tidy
```

#### Error: "cannot find module"
```
Solution: Download dependencies
go mod download
go mod tidy
```

### Database Errors

#### Error: "connection refused"
```
Solution: 
1. Check PostgreSQL running: docker-compose ps
2. Verify credentials di .env
3. Check DB_HOST adalah localhost (tidak 0.0.0.0)
```

#### Error: "database does not exist"
```
Solution:
createdb boilerplate_go
make migrate-up
```

### Runtime Errors

#### Error: "address already in use"
```
Solution: Port 8080 sudah terpakai
# Find process using port 8080
lsof -i :8080  # macOS/Linux
netstat -ano | findstr :8080  # Windows

# Kill process atau gunakan port lain
APP_PORT=8081 make run
```

#### Error: "panic: json: cannot unmarshal"
```
Solution: Request JSON format salah
✓ Pastikan request body adalah valid JSON
✓ Pastikan Content-Type: application/json
✓ Pastikan field names sesuai struct tags
```

### Docker Errors

#### Error: "port 5432 already in use"
```
Solution:
docker-compose down
docker-compose up -d
```

#### Error: "cannot create container: mkdir ... permission denied"
```
Solution: Run dengan sudo atau fix Docker permissions
sudo docker-compose up -d
```

### General Debugging

#### Enable Debug Logging
```bash
APP_ENV=development make dev
```

#### Check Database Connection
```bash
psql -h localhost -U postgres -d boilerplate_go
```

#### View Application Logs
```bash
docker-compose logs -f app
# atau
journalctl -u boilerplate-go -f
```

#### Test API Connectivity
```bash
curl -v http://localhost:8080/health
```

---

## 📚 Additional Resources

- [Go Documentation](https://golang.org/doc/)
- [Gin Framework](https://gin-gonic.com/)
- [GORM Documentation](https://gorm.io/)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
- [Zap Logger](https://github.com/uber-go/zap)
- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

---

## 📄 License

MIT License - feel free to use this boilerplate for your projects!

---

## 💬 Support

Jika ada pertanyaan atau issues:
1. Check dokumentasi section yang relevan
2. Lihat troubleshooting
3. Buat GitHub issue dengan detail problem

---

**Happy Coding! 🚀**
