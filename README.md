# PumplePet Server API

## Setup Requirements
1. PostgreSQL running locally
2. `.env` file with database credentials
3. Run `go run cmd/main.go` to start server

## Authentication Endpoints

### Register User
```bash
curl -X POST http://localhost:8000/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "test@example.com", 
    "password": "password123"
  }'
```

### Login User
```bash
curl -X POST http://localhost:8000/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123"
  }'
```

## Expected Responses

### Successful Register
```json
{
  "user": {
    "id": 1,
    "username": "testuser",
    "email": "test@example.com",
    "is_owner": false
  }
}
```

### Successful Login
```json
{
  "user": {
    "id": 1,
    "username": "testuser", 
    "email": "test@example.com",
    "is_owner": false
  }
}
```

### Error Response
```json
{
  "error": "Error message here"
}
```