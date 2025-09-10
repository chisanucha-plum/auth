# 🔐 Auth API Server

Simple authentication API server built with Go, Gin, GORM, and JWT.

## 🚀 Features

- ✅ User Registration with JWT token response
- ✅ User Login with JWT authentication
- ✅ Protected routes with JWT middleware
- ✅ PostgreSQL database with auto-migration
- ✅ Password hashing with bcrypt
- ✅ Environment configuration

## 📦 Tech Stack

- **Go** - Programming language
- **Gin** - HTTP web framework
- **GORM** - ORM library
- **PostgreSQL** - Database
- **JWT** - Authentication tokens
- **bcrypt** - Password hashing

## 🛠 Setup

### 1. Prerequisites
- Go 1.24+
- PostgreSQL database

### 2. Environment Variables
Create `.env` file:
```env
# Server
PORT=8002

# Database
DB_HOST=localhost
DB_PORT=5435
DB_USER=myuser
DB_PASSWORD=1234
DB_NAME=mydbs

# JWT
JWT_SECRET=your-secret-key-here
```

### 3. Install Dependencies
```bash
go mod tidy
```

### 4. Run Server
```bash
go run cmd/main.go
```

Server will start on `http://localhost:8002`

## 📋 API Endpoints

### Authentication

#### Register User
```http
POST /api/v1/auth/register
Content-Type: application/json

{
  "username": "testuser",
  "email": "test@example.com", 
  "password": "password123",
  "full_name": "Test User"
}
```

**Response:**
```json
{
  "success": true,
  "message": "User registered successfully",
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

#### Login User
```http
POST /api/v1/auth/login
Content-Type: application/json

{
  "username": "testuser",
  "password": "password123"
}
```

**Response:**
```json
{
  "success": true,
  "message": "Login successful", 
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

### Protected Routes

#### Get Profile
```http
GET /api/v1/protected/profile
Authorization: Bearer <access_token>
```

**Response:**
```json
{
  "success": true,
  "message": "Profile retrieved successfully",
  "data": {
    "user_id": 1,
    "issued_at": "2025-09-10T14:27:48Z",
    "expires_at": "2025-09-11T14:27:48Z"
  }
}
```

## 🗄️ Database Schema

### Users Table
```sql
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMP,
  updated_at TIMESTAMP, 
  deleted_at TIMESTAMP,
  username VARCHAR(255) UNIQUE NOT NULL,
  email VARCHAR(255) UNIQUE,
  password VARCHAR(255) NOT NULL,
  full_name VARCHAR(255),
  is_active BOOLEAN DEFAULT true
);
```

## 🧪 Testing with cURL

### 1. Register a new user
```bash
curl -X POST http://localhost:8002/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "test@example.com",
    "password": "password123",
    "full_name": "Test User"
  }'
```

### 2. Login with credentials
```bash
curl -X POST http://localhost:8002/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser", 
    "password": "password123"
  }'
```

### 3. Access protected route
```bash
curl -X GET http://localhost:8002/api/v1/protected/profile \
  -H "Authorization: Bearer YOUR_ACCESS_TOKEN"
```

## 📁 Project Structure

```
├── cmd/
│   └── main.go              # Application entry point
├── database/
│   └── db.go               # Database connection & migration
├── internal/
│   ├── configuration/
│   │   └── envs.go         # Environment configuration
│   ├── handlers/
│   │   └── auth.go         # HTTP handlers
│   ├── models/
│   │   ├── api.go          # Request/Response models
│   │   └── user.go         # User model
│   ├── routes/
│   │   └── routes.go       # Route definitions
│   └── services/
│       └── auth.go         # Business logic
├── middleware/
│   └── auth.go             # JWT middleware
├── .env                    # Environment variables
├── go.mod                  # Go modules
└── README.md               # This file
```

## 🔒 Security Features

- Password hashing with bcrypt
- JWT token-based authentication
- Environment-based configuration
- Protected routes with middleware
- SQL injection prevention with GORM


