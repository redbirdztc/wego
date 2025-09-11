# Wego

Web service boilerplate in Go using Fiber, GORM, and PostgreSQL.

## Prerequisites

### Required Software
- Go 1.24+
- PostgreSQL 12+
- GNU Make
- Docker (optional)

### Required Tools
Install the following Go tools:
```bash
# Swagger documentation generator
go install github.com/swaggo/swag/cmd/swag@latest
```

## Environment Variables
```env
POSTGRES_DSN=postgres://user:pass@localhost:5432/nodame?sslmode=disable
HTTP_PORT=8080
```

## Quick Start

1. Copy environment variables:
```bash
cp .env.example .env
```

2. Build and run:
```bash
make build
./bin/main
```

Or run directly:
```bash
go run cmd/api/main.go
```

3. Access the API at `http://localhost:8080`