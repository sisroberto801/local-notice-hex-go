# Local Notice Hexagonal Go

**Repository:** https://github.com/sisroberto801/local-notice-hex-go.git

## Technology Stack

- **Go** - Programming language
- **Gin** - HTTP web framework
- **PostgreSQL** - Database
- **JWT** - Authentication (using golang-jwt/jwt/v5)
- **GORM** - ORM for database operations
- **Swagger** - API documentation

## Hexagonal Architecture

The project follows a hexagonal (clean architecture) with the following layers:

```
local-notice-hex-go/
├── main.go                        # Main application entry point
├── configs/                       # Configuration layer
│   └── config.go                  # Application configuration
├── internal/                      # Internal application code
│   ├── domain/                    # Domain layer (core)
│   │   └── user/                  # User domain entities and ports
│   │       ├── user.go            # User entity and DTOs
│   │       └── repository.go      # Repository interface
│   ├── service/                   # Application services
│   │   └── user/                  # User service implementations
│   │       └── service.go         # User service logic
│   └── infrastructure/            # Infrastructure layer
│       ├── database/              # Database implementations
│       │   └── postgres/          # PostgreSQL repository
│       │       └── repository.go  # User repository implementation
│       └── http/                  # HTTP layer
│           ├── handler/           # HTTP handlers
│           │   └── user_handler.go # User HTTP handlers
│           ├── middleware/        # HTTP middleware
│           │   └── middleware.go  # JWT and CORS middleware
│           └── router/            # HTTP router setup
│               └── router.go      # Route configuration
├── migrations/                     # Database migrations
├── pkg/                          # Public packages
│   ├── database/                 # Database utilities
│   │   └── database.go           # PostgreSQL connector
│   └── migration/                # Migration utilities
│       └── migrator.go           # Database migrator
└── swagger-ui/                    # API documentation
```

### Architecture Flow:
1. **Controller/Handler** receives HTTP requests
2. **Service** coordinates use cases and business logic
3. **Domain** contains pure business entities and interfaces
4. **Infrastructure** implements technical details (DB, HTTP handlers)

## Run Application

```bash
# Create database
createdb notice_db

# Copy environment file
cp .env.example .env

# Install dependencies
go mod download

# Run the application
go run main.go

# Or build and run
go build -o notice-app
./notice-app

# Check if application is running
ps aux | grep "go run\|main.go\|:8080" | grep -v grep
```

## Environment Configuration

The application uses a `.env` file for configuration:

1. **Copy the template**:
   ```bash
   cp .env.example .env
   ```

2. **Edit `.env`** with your configuration:
   ```env
   SERVER_PORT=8080
   DATABASE_URL=postgres://postgres:postgres@localhost:5432/notice_db?sslmode=disable
   JWT_SECRET=your-secret-key
   ENVIRONMENT=development
   ```

3. **Security Note**: Change `JWT_SECRET` in production to a secure, random string.

## API Documentation

Once the application is running, you can access the Swagger UI at:

- **Swagger UI**: http://localhost:8080/swagger-ui/
- **OpenAPI Info**: http://localhost:8080/swagger-ui/project.swagger.yaml

The Swagger documentation provides interactive API testing and detailed endpoint information.

### REST Endpoints

#### Authentication
- `POST /api/auth/login` - User login and JWT token generation

#### Users Management
- `POST /api/users` - Create a new user (public)
- `GET /api/users/:id` - Get user by ID (public)
- `GET /api/users` - Get all users (protected - JWT required)
- `PUT /api/users/:id` - Update user (protected - JWT required)
- `DELETE /api/users/:id` - Delete user (protected - JWT required)

#### User Models

**User Request (Create/Update)**
```json
{
  "username": "john_doe",
  "password": "securepassword123",
  "status": true
}
```

**User Response**
```json
{
  "id": 50,
  "username": "john_doe",
  "status": true,
  "createdAt": "2026-05-12T22:19:51.105964Z",
  "updatedAt": "2026-05-12T22:19:51.105984Z"
}
```

## Database

- PostgreSQL: `notice_db`
- User: `postgres`
- Password: `postgres`
- Port: `5432`

**Connection URL:** `postgres://postgres:postgres@localhost:5432/notice_db?sslmode=disable`

## Configuration

The application uses environment variables for configuration:

- `SERVER_PORT`: Server port (default: 8080)
- `DATABASE_URL`: PostgreSQL connection URL
- `JWT_SECRET`: JWT secret key for authentication
- `ENVIRONMENT`: Application environment (default: development)

## Authentication

The application uses JWT (JSON Web Tokens) for authentication:

1. **Login**: POST `/api/auth/login` with username and password
2. **Token**: Receive JWT token in response
3. **Protected Routes**: Include JWT token in Authorization header: `Bearer <token>`
4. **Token Validation**: Middleware validates JWT token for protected endpoints

**Note**: The `pkg/auth` package exists but is currently not used. The application uses `jwt.MapClaims` directly in the middleware and handlers.