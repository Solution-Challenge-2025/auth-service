# Auth Service

A robust authentication service built with Go, providing user registration, login, and JWT-based authentication with role-based access control.

## Features

- User registration and login
- JWT-based authentication
- Role-based access control (Admin and User roles)
- PostgreSQL database integration
- Secure password hashing
- Environment-based configuration

## Prerequisites

- Go 1.23.6 or higher
- PostgreSQL 12 or higher
- Make (optional, for using Makefile commands)

## Setup

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd auth-service
   ```

2. Install dependencies:

   ```bash
   go mod download
   ```

3. Create a PostgreSQL database:

   ```sql
   CREATE DATABASE auth_db;
   ```

4. Configure environment variables:
   Create a `.env` file in the root directory with the following content:

   ```env
   # Database Configuration
   URI=postgresql://postgres:postgres@localhost:5432/auth_db?sslmode=disable

   # JWT Configuration
   JWT_SECRET_KEY=your-super-secret-key-change-this-in-production

   # Server Configuration
   PORT=8080
   ```

   Update the values according to your environment.

5. Run the service:
   ```bash
   go run main.go
   ```

The service will start on `http://localhost:8080` by default.

## API Documentation

### Base URL

```
http://localhost:8080/api/v1
```

### Endpoints

#### 1. Register User

Register a new user in the system.

- **URL**: `/auth/register`
- **Method**: `POST`
- **Content-Type**: `application/json`

**Request Body**:

```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123",
  "role": "user" // Optional, defaults to "user". Can be "admin" or "user"
}
```

**Success Response** (201 Created):

```json
{
  "message": "user registered successfully",
  "user": {
    "id": "uuid-here",
    "name": "John Doe",
    "email": "john@example.com",
    "role": "user"
  }
}
```

**Error Responses**:

- 400 Bad Request: Invalid input data
- 409 Conflict: Email already registered

#### 2. Login User

Authenticate user and receive JWT token.

- **URL**: `/auth/login`
- **Method**: `POST`
- **Content-Type**: `application/json`

**Request Body**:

```json
{
  "email": "john@example.com",
  "password": "password123"
}
```

**Success Response** (200 OK):

```json
{
  "token": "jwt-token-here"
}
```

**Error Responses**:

- 400 Bad Request: Invalid input data
- 401 Unauthorized: Invalid credentials

#### 3. Validate Token

Validate a JWT token and return its contents.

- **URL**: `/auth/validate`
- **Method**: `GET`
- **Headers**:
  ```
  Authorization: Bearer your-jwt-token-here
  ```

**Success Response** (200 OK):

```json
{
  "user_id": "uuid-here",
  "role": "user",
  "exp": 1234567890,
  "iat": 1234567890,
  "iss": "Auth Service"
}
```

**Error Responses**:

- 401 Unauthorized: Missing or invalid token
  ```json
  {
    "error": "authorization header is required"
  }
  ```
  or
  ```json
  {
    "error": "invalid authorization format"
  }
  ```
  or
  ```json
  {
    "error": "invalid token"
  }
  ```

### Protected Routes

Protected routes require a valid JWT token in the Authorization header.

**Authorization Header**:

```
Authorization: Bearer your-jwt-token-here
```

## Example Usage

### Using cURL

1. **Register a new user**:

   ```bash
   curl -X POST http://localhost:8080/api/v1/auth/register \
   -H "Content-Type: application/json" \
   -d '{
     "name": "John Doe",
     "email": "john@example.com",
     "password": "password123",
     "role": "user"
   }'
   ```

2. **Login with the registered user**:

   ```bash
   curl -X POST http://localhost:8080/api/v1/auth/login \
   -H "Content-Type: application/json" \
   -d '{
     "email": "john@example.com",
     "password": "password123"
   }'
   ```

3. **Access protected route**:

   ```bash
   curl -X GET http://localhost:8080/api/v1/admin/protected-route \
   -H "Authorization: Bearer your-jwt-token-here"
   ```

4. **Validate token**:
   ```bash
   curl -X GET http://localhost:8080/api/v1/auth/validate \
   -H "Authorization: Bearer your-jwt-token-here"
   ```

### Using Postman

1. Create a new collection for the Auth Service
2. Import the following requests:

**Register User**:

- Method: POST
- URL: `http://localhost:8080/api/v1/auth/register`
- Headers: `Content-Type: application/json`
- Body (raw JSON):
  ```json
  {
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123",
    "role": "user"
  }
  ```

**Login User**:

- Method: POST
- URL: `http://localhost:8080/api/v1/auth/login`
- Headers: `Content-Type: application/json`
- Body (raw JSON):
  ```json
  {
    "email": "john@example.com",
    "password": "password123"
  }
  ```

**Protected Route**:

- Method: GET
- URL: `http://localhost:8080/api/v1/admin/protected-route`
- Headers:
  - `Authorization: Bearer your-jwt-token-here`

**Validate Token**:

- Method: GET
- URL: `http://localhost:8080/api/v1/auth/validate`
- Headers:
  - `Authorization: Bearer your-jwt-token-here`

## Error Handling

The service returns appropriate HTTP status codes and error messages:

- **400 Bad Request**: Invalid input data

  ```json
  {
    "error": "validation error message"
  }
  ```

- **401 Unauthorized**: Invalid credentials or missing token

  ```json
  {
    "error": "invalid credentials"
  }
  ```

- **403 Forbidden**: Insufficient permissions

  ```json
  {
    "error": "insufficient permissions"
  }
  ```

- **409 Conflict**: Resource conflict
  ```json
  {
    "error": "email already registered"
  }
  ```

## Security Considerations

1. All passwords are hashed using bcrypt before storage
2. JWT tokens expire after 24 hours
3. Environment variables are used for sensitive configuration
4. Role-based access control is implemented for protected routes
5. Input validation is performed on all requests

## Development

### Project Structure

```
auth-service/
├── config/
│   └── database.go
├── internal/
│   ├── handlers/
│   │   └── auth-handler.go
│   ├── models/
│   │   └── models.go
│   ├── repositories/
│   │   └── auth-repository.go
│   └── services/
│       └── auth-service.go
├── package/
│   ├── jwt/
│   │   └── jwt.go
│   └── middleware/
│       └── auth-middleware.go
├── routes/
│   └── routes.go
├── .env
├── go.mod
├── go.sum
└── main.go
```

### Adding New Features

1. Create new models in `internal/models/`
2. Add repository methods in `internal/repositories/`
3. Implement service logic in `internal/services/`
4. Create handlers in `internal/handlers/`
5. Add routes in `routes/routes.go`

## License

[Your License Here]
