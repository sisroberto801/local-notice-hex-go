# Local Notice Hexagonal Go

**Repository:** https://github.com/sisroberto801/local-notice-hex-go.git

## Quick Start

```bash
# Create database
createdb notice_db

# Setup environment
cp .env.example .env

# Install dependencies
go mod download

# Run the application
go run main.go
```

## Technology Stack

- **Go** - Programming language
- **Gin** - HTTP web framework
- **PostgreSQL** - Database
- **JWT** - Authentication (using golang-jwt/jwt/v5)
- **GORM** - ORM for database operations
- **Swagger** - API documentation
- **Docker** - Containerization
- **GCP Cloud Run** - Cloud deployment (ready)

## Hexagonal Architecture

The project follows a hexagonal (clean architecture) with the following layers:

```
local-notice-hex-go/
├── main.go                        # Main application entry point
├── configs/                       # Configuration layer
│   └── config.go                  # Application configuration
├── internal/                      # Internal application code
│   ├── domain/                    # Domain layer (core)
│   │   ├── model/                 # Domain entities
│   │   │   └── user.go           # User entity and DTOs
│   │   └── ports/                 # Domain interfaces
│   │       └── user_repository_port.go # Repository interface
│   ├── service/                   # Application services
│   │   └── user/                  # User service implementations
│   │       └── service.go         # User service logic
│   └── infrastructure/            # Infrastructure layer
│       ├── database/              # Database implementations
│       │   └── postgres/          # PostgreSQL repository
│       │       └── user_repository_adapter.go # User repository implementation
│       └── http/                  # HTTP layer
│           ├── handler/           # HTTP handlers
│           │   └── user_handler.go # User HTTP handlers
│           ├── middleware/        # HTTP middleware
│           │   └── middleware.go  # JWT and CORS middleware
│           └── router/            # HTTP router setup
│               └── router.go      # Route configuration
├── Dockerfile                     # Docker container configuration
├── docker-compose.yml             # Docker Compose for local development
├── pkg/                          # Public packages
│   ├── database/                 # Database utilities
│   │   └── database.go           # PostgreSQL connector
│   └── migration/                # Migration utilities
│       ├── 001_create_users_table.sql # Database migration file
│       └── migrator.go           # Database migrator
└── swagger-ui/                    # API documentation
    ├── index.html                 # Swagger UI interface
    └── project.swagger.yaml      # OpenAPI specification
```

### Architecture Flow:
1. **Controller/Handler** receives HTTP requests
2. **Service** coordinates use cases and business logic
3. **Domain** contains pure business entities and interfaces
4. **Infrastructure** implements technical details (DB, HTTP handlers)

## Run Application

### Local Development

For detailed setup, see the Quick Start section above.

### Docker Development

```bash
# Copy environment file
cp .env.example .env

# Build and run with Docker Compose
docker-compose up --build

# Run in background
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down

# Stop and remove volumes
docker-compose down -v
```

### Health Check

Once running, verify the application health:

```bash
# Health endpoint
curl http://localhost:8080/health

# Ready endpoint
curl http://localhost:8080/ready
```

## Environment Configuration

The application uses a `.env` file for configuration:

1. **Copy the template**:
   ```bash
   cp .env.example .env
   ```

2. **Edit `.env`** with your configuration:
   ```env
   # Server Configuration
   PORT=8080

   # PostgreSQL Configuration (base variables)
   POSTGRES_DB=notice_db
   POSTGRES_USER=postgres
   POSTGRES_PASSWORD=postgres
   POSTGRES_HOST=postgres
   POSTGRES_PORT=5432
   POSTGRES_SSLMODE=disable

   # Generated Database URL (built from individual variables)
   DATABASE_URL=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=${POSTGRES_SSLMODE}

   # JWT Configuration
   JWT_SECRET=your-secret-key

   # Environment
   ENVIRONMENT=development
   ```

3. **Security Note**: Change `JWT_SECRET` in production to a secure, random string.

## API Documentation

- **Swagger UI**: http://localhost:8080/swagger-ui/
- **OpenAPI Spec**: http://localhost:8080/swagger-ui/project.swagger.yaml

### Key Endpoints

- `GET /health` - Health check
- `GET /ready` - Readiness check  
- `POST /api/auth/login` - User login
- `GET /api/users` - List users (protected)
- `POST /api/users` - Create user
- `GET /api/users/:id` - Get user by ID
- `PUT /api/users/:id` - Update user (protected)
- `DELETE /api/users/:id` - Delete user (protected)


## Authentication

The application uses JWT (JSON Web Tokens) for authentication:

1. **Login**: POST `/api/auth/login` with username and password
2. **Token**: Receive JWT token in response
3. **Protected Routes**: Include JWT token in Authorization header: `Bearer <token>`
4. **Token Validation**: Middleware validates JWT token for protected endpoints

**Note**: The `pkg/auth` package exists but is currently not used. The application uses `jwt.MapClaims` directly in the middleware and handlers.

## GCP Cloud Run Deployment

The application is ready for deployment on Google Cloud Platform using Cloud Run.

### Prerequisites
- GCP project with Cloud Run API enabled
- Docker installed locally
- gcloud CLI configured

### Build and Deploy

```bash
# Build Docker image
docker build -t gcr.io/[PROJECT-ID]/local-notice-hex-go .

# Push to Artifact Registry
docker push gcr.io/[PROJECT-ID]/local-notice-hex-go

# Deploy to Cloud Run
gcloud run deploy local-notice-hex-go \
  --image gcr.io/[PROJECT-ID]/local-notice-hex-go \
  --platform managed \
  --region [REGION] \
  --allow-unauthenticated \
  --set-env-vars "PORT=8080" \
  --set-env-vars "POSTGRES_HOST=[CLOUD_SQL_INSTANCE_IP]" \
  --set-env-vars "POSTGRES_DB=notice_db" \
  --set-env-vars "POSTGRES_USER=postgres" \
  --set-env-vars "POSTGRES_PASSWORD=[PASSWORD]" \
  --set-env-vars "JWT_SECRET=[JWT_SECRET]" \
  --set-env-vars "ENVIRONMENT=production"
```

### Terraform Deployment

For infrastructure-as-code deployment, use Terraform with the following resources:

- **Cloud Run Service** - Application container
- **Cloud SQL** - PostgreSQL database
- **Artifact Registry** - Docker image storage
- **Service Account** - IAM permissions
- **Secret Manager** - Sensitive data storage

### Health Checks

Cloud Run automatically uses the `/health` endpoint for health checks with:
- **Interval**: 30 seconds
- **Timeout**: 5 seconds
- **Start period**: 5 seconds
- **Retries**: 3

### Environment Variables for Production

For production deployment, set these environment variables:
- `PORT`: Cloud Run port (8080)
- `POSTGRES_HOST`: Cloud SQL instance connection name
- `POSTGRES_DB`: Database name
- `POSTGRES_USER`: Database user
- `POSTGRES_PASSWORD`: Database password (use Secret Manager)
- `JWT_SECRET`: JWT secret (use Secret Manager)
- `ENVIRONMENT`: production

### Security Considerations

- **Non-root user**: Container runs as non-root user (UID 1001)
- **HTTPS**: Cloud Run automatically provides HTTPS
- **Secrets**: Use Secret Manager for sensitive data
- **VPC**: Consider VPC connector for private Cloud SQL access