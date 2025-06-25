![CI](https://github.com/SilverNate/book-lending-api/actions/workflows/ci.yml/badge.svg)

# 📚 Book Lending API

A simple RESTful API for borrowing and returning books, built with Go, Gin, MySQL, and Redis. Follows Clean Architecture and SOLID principles with dependency injection using Wire.

---

## ✨ Features

- User Authentication (JWT)
- Book CRUD (with soft delete & metadata)
- Borrow & Return Books
- Rate Limiting (max 5 borrows per 7 days per user)
- Unit Tests (using mockery & testify)
- Docker & Docker Compose support
- Swagger documentation
- Logging to Elasticsearch via Logrus

---

## 🚀 Getting Started

### 1. Clone & Setup
```bash
git clone https://github.com/yourname/book-lending-api.git
cd book-lending-api
```
### 2. Create .env File
```bash
PORT=8080
DB_DSN=bookuser:bookpass@tcp(mysql:3306)/bookdb?parseTime=true
JWT_SECRET=supersecret123
JWT_ISSUER=book-lending-api
```

### 3. Run with Docker
```bash
docker-compose up --build
```
- App runs on: http://localhost:8080
- Swagger docs: http://localhost:8080/swagger/index.html
- Kibana UI: http://localhost:5601
- Elasticsearch: http://localhost:9200

### 4. Migrations 
already migrations using docker-compose, but to make sure run this:
- up
```bash
migrate -path ./migrations -database "mysql://bookuser:bookpass@tcp(localhost:3306)/bookdb?parseTime=true" up
```
- rollback
```bash
migrate -path ./migrations -database "mysql://bookuser:bookpass@tcp(localhost:3306)/bookdb?parseTime=true" down
```
### 5. Unit Test
```bash
 go test -v ./internal/...
```
### 6. Mockery
- sample
```bash
 mockery --name=IBookRepository    --dir=internal/book/repository         --output=internal/book/mocks         --with-expecter
```
### 6. Tech Stack
```bash
Golang
Gin
GORM
MySQL
Redis
Wire
Logrus + Elasticsearch
Swagger UI
```