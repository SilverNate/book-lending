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
### 7. Tech Stack
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

### 8. Add Sample Postman Curl
- register
```bash
curl --location 'http://localhost:8080/auth/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "me@mail.com",
    "password": "abc12345"
}'
```
- login
```bash
- curl --location 'http://localhost:8080/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{"email":"me@mail.com", "password":"abc12345"}'
```
- add book
```bash
- curl --location 'http://localhost:8080/books' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im1lQG1haWwuY29tIiwiZXhwIjoxNzUxMDkwNjQ4LCJpc3MiOiJib29rLWxlbmRpbmctYXBpIiwidXNlcl9pZCI6MX0.FNSMH2kIk60gdLVWMjEqGTf4rbnJntRZau6qXGPqj2s' \
--data '{
    "title": "Clean Architecture",
    "author": "Robert C. Martin",
    "isbn": "9780134494166",
    "category": "Software Engineering",
    "quantity": 5
}'
```
- get all book
```bash
curl --location 'http://localhost:8080/books?page=0&limit=10' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im1lQG1haWwuY29tIiwiZXhwIjoxNzUxMDQ4NDk5LCJpc3MiOiJib29rLWxlbmRpbmctYXBpIiwidXNlcl9pZCI6MX0.rcmlmNvuUMbqpgBVOzFpGHJD1S4UvTHRi9KDzrqOUD8'
```
- get book by id
```bash
curl --location 'http://localhost:8080/books/1' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im1lQG1haWwuY29tIiwiZXhwIjoxNzUxMDQ4NDk5LCJpc3MiOiJib29rLWxlbmRpbmctYXBpIiwidXNlcl9pZCI6MX0.rcmlmNvuUMbqpgBVOzFpGHJD1S4UvTHRi9KDzrqOUD8'
```
- update book
```bash
curl --location --request PUT 'http://localhost:8080/books/1' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im1lQG1haWwuY29tIiwiZXhwIjoxNzUxMDQ4NDk5LCJpc3MiOiJib29rLWxlbmRpbmctYXBpIiwidXNlcl9pZCI6MX0.rcmlmNvuUMbqpgBVOzFpGHJD1S4UvTHRi9KDzrqOUD8' \
--header 'Content-Type: application/json' \
--data '{
    "title": "Clean Architecture - Updated",
    "author": "Robert C. Martin",
    "isbn": "9780134494166",
    "category": "Architecture",
    "quantity": 10
}'
```
- delete book
```bash
curl --location --request DELETE 'http://localhost:8080/books/1' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im1lQG1haWwuY29tIiwiZXhwIjoxNzUxMDQ4NDk5LCJpc3MiOiJib29rLWxlbmRpbmctYXBpIiwidXNlcl9pZCI6MX0.rcmlmNvuUMbqpgBVOzFpGHJD1S4UvTHRi9KDzrqOUD8'
```
- borrowing book
```bash
curl --location 'http://localhost:8080/borrowing/borrow' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im1lQG1haWwuY29tIiwiZXhwIjoxNzUxMDU1MjU5LCJpc3MiOiJib29rLWxlbmRpbmctYXBpIiwidXNlcl9pZCI6MX0.icH0Pynqsx5kb02mfOEC6cNR6K9CEKFTNf7YC_FVtDk' \
--header 'Content-Type: application/json' \
--data '{
    "book_id": 1
}'
```
- return book
```bash
curl --location 'http://localhost:8080/borrowing/return' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im1lQG1haWwuY29tIiwiZXhwIjoxNzUxMDU1MjU5LCJpc3MiOiJib29rLWxlbmRpbmctYXBpIiwidXNlcl9pZCI6MX0.icH0Pynqsx5kb02mfOEC6cNR6K9CEKFTNf7YC_FVtDk' \
--header 'Content-Type: application/json' \
--data '{"borrowing_id": 1}'
```