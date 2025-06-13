# PaymentSystem - Crypto E-commerce Platform

A modern e-commerce platform built with Go that enables cryptocurrency payments through NowPayments integration. This MVP focuses on core e-commerce functionality with crypto-only payment processing.

## 🚀 What This Project Does

**PaymentSystem** is a complete e-commerce solution that allows:

- **Customers** to browse products and pay with cryptocurrency
- **Merchants** to manage inventory and process crypto payments
- **Automated** order fulfillment and payment verification

### Key Features

- 🛍️ **Product Catalog** - Complete product management with inventory tracking
- 🛒 **Shopping Cart** - Add/remove items with real-time price calculation
- 💰 **Crypto Payments** - Bitcoin, Ethereum, and other cryptocurrencies via NowPayments
- 📦 **Order Management** - Full order lifecycle from creation to fulfillment
- 🚚 **Shipping Integration** - Address management and shipping calculations
- 💸 **Discount System** - Coupon codes and promotional pricing
- 🧾 **Tax Calculation** - Automated tax computation
- 🔄 **Refund Processing** - Handle payment refunds and disputes

## 🏗️ Architecture

Built using **Clean Architecture** and **Domain-Driven Design** principles:

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   REST API      │    │   Admin Panel   │    │   Webhooks      │
└─────────────────┘    └─────────────────┘    └─────────────────┘
┌─────────────────────────────────────────────────────────────────┐
│                    Application Layer                            │
│              Use Cases • Services • Orchestration              │
└─────────────────────────────────────────────────────────────────┘
┌─────────────────────────────────────────────────────────────────┐
│                      Domain Layer                               │
│           Business Logic • Entities • Value Objects            │
└─────────────────────────────────────────────────────────────────┘
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   PostgreSQL    │    │   NowPayments   │    │   HTTP Client   │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

## 🛠️ Tech Stack

- **Backend**: Go 1.21+, Gin/Echo HTTP framework
- **Database**: PostgreSQL 15+ with migrations
- **Payments**: NowPayments API integration
- **Testing**: Testify, comprehensive test coverage
- **Architecture**: Clean Architecture, DDD patterns
- **Deployment**: Docker, containerized deployment

## 🚦 Getting Started

### Prerequisites

```bash
# Required software
- Go 1.21 or higher
- PostgreSQL 15+
- NowPayments API account
- Git
```

### Installation

```bash
# Clone repository
git clone https://github.com/your-username/paymentSystem.git
cd paymentSystem

# Install dependencies
go mod download

# Set up environment variables
cp .env.example .env
# Edit .env with your configuration

# Run database migrations
make migrate-up

# Start the server
go run cmd/api/main.go
```

### Environment Configuration

```env
# Database
DB_HOST=localhost
DB_PORT=5432
DB_NAME=paymentsystem
DB_USER=postgres
DB_PASSWORD=your_password

# NowPayments
NOWPAYMENTS_API_KEY=your_api_key
NOWPAYMENTS_IPN_SECRET=your_secret
NOWPAYMENTS_SANDBOX=true

# Server
SERVER_PORT=8080
SERVER_ENV=development
```

## 📚 API Usage

### Create Order

```bash
POST /api/v1/orders
{
  "customer_id": "uuid",
  "items": [
    {
      "product_id": "uuid",
      "quantity": 2
    }
  ]
}
```

### Process Payment

```bash
POST /api/v1/orders/{id}/checkout
{
  "shipping_address_id": "uuid",
  "preferred_crypto": "BTC"
}
```

### Check Payment Status

```bash
GET /api/v1/payments/{id}
```

## 🗃️ Database Schema

Core entities and relationships:

- **Products** - Catalog items with inventory
- **Orders** - Customer purchases and status
- **Payments** - Cryptocurrency transactions
- **Customers** - User accounts and addresses
- **Categories** - Product organization

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific domain tests
go test ./internal/domain/order/...
go test ./internal/domain/product/...
```

## 📖 Documentation

- [Complete System Design](docs/PaymentSystem_MVP_Documentation.md) - Comprehensive technical specification
- API Documentation - REST API endpoints and examples
- Database Schema - Database design and relationships
- NowPayments Integration - Payment processing guide

## 🔄 Development Workflow

```bash
# Run in development mode
make dev

# Build for production
make build

# Run database migrations
make migrate-up

# Run tests
make test

# Generate mocks
make mocks
```

## 🚀 Deployment

### Docker

```bash
# Build and run with Docker
docker-compose up --build

# Production deployment
docker-compose -f docker-compose.prod.yml up -d
```

### Manual Deployment

```bash
# Build binary
go build -o bin/api cmd/api/main.go

# Run with production config
./bin/api -config=production.yaml
```
