# Local Notice Hexagonal Go

**Repository:** https://github.com/sisroberto801/local-notice-hex-go.git

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
│   ├── service/                   # Application services
│   │   └── user/                  # User service implementations
│   └── infrastructure/            # Infrastructure layer
│       ├── database/              # Database implementations
│       │   └── postgres/          # PostgreSQL repository
│       └── http/                  # HTTP layer
│           ├── handler/           # HTTP handlers
│           └── router/            # HTTP router setup
├── migrations/                     # Database migrations
├── swagger-ui/                    # API documentation
└── pkg/                          # Public packages
```

### Architecture Flow:
1. **Controller** receives HTTP requests
2. **Service** coordinates use cases
3. **Domain** contains pure business logic
4. **Infrastructure** implements technical details (DB, external APIs)

## Run Application

```bash
# Create database
createdb notice_db

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

## API Documentation

Once the application is running, you can access the Swagger UI at:

- **Swagger UI**: http://localhost:8080/swagger-ui/
- **OpenAPI Info**: http://localhost:8080/swagger-ui/project.swagger.yaml

The Swagger documentation provides interactive API testing and detailed endpoint information.

### REST Endpoints

#### Authentication
- `POST /api/auth/login` - User login and JWT token generation

#### Users Management
- `GET /api/users` - Get all users
- `POST /api/users` - Create a new user
- `GET /api/users/{id}` - Get user by ID
- `PUT /api/users/{id}` - Update user
- `DELETE /api/users/{id}` - Delete user

#### User Model
```json
{
  "id": 50,
  "username": "s",
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