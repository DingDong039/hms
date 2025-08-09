# Hospital Middleware System (HMS) Project Structure

This document outlines the structure and organization of the Hospital Middleware System project.

**Related Documentation:**
- [README](../README.md) - Main project documentation
- [API Specification](./api_spec.md) - API endpoints and usage
- [ER Diagram](./er_diagram.md) - Database schema and relationships

## Overview

The Hospital Middleware System (HMS) follows a clean architecture pattern with clear separation of concerns. The project is organized into several layers:

1. **Presentation Layer**: API handlers and middleware
2. **Business Logic Layer**: Services implementing core business logic
3. **Data Access Layer**: Repositories for database operations
4. **Domain Layer**: Models representing business entities
5. **Infrastructure**: Configuration, database connection, and utilities

## Directory Structure

```
HMS/
├── cmd/                          # Application entry points
│  └── main/                      # Main application
│      └── main.go                # Entry point
├── internal/                     # Private application code
│   ├── config/                   # Configuration management
│   │   └── config.go             # Configuration loading and structures
│   ├── handlers/                 # HTTP request handlers (controllers)
│   │   ├── auth_handler.go       # Authentication endpoints
│   │   ├── patient_handler.go    # Patient search endpoint
│   │   └── routes.go             # Route registration
│   ├── services/                 # Business logic layer
│   │   ├── auth_service.go       # Authentication logic
│   │   ├── patient_service.go    # Patient business logic
│   │   └── hospital_api_service.go # External API integration
│   ├── repositories/             # Data access layer
│   │   ├── base_repository.go    # Base repository pattern
│   │   ├── staff_repository.go   # Staff database operations
│   │   └── patient_repository.go # Patient database operations
│   ├── models/                   # Domain models
│   │   ├── staff.go              # Staff entity and DTOs
│   │   ├── patient.go            # Patient entity and DTOs
│   │   └── response.go           # API response models
│   ├── middleware/               # HTTP middleware
│   │   ├── auth_middleware.go    # JWT authentication
│   │   ├── cors_middleware.go    # CORS handling
│   │   └── logging_middleware.go # Request logging
│   ├── database/                 # Database infrastructure
│   │   ├── connection.go         # Database connection
│   │   └── migrations.go         # Database migrations
│   └── utils/                    # Utility functions
│       ├── jwt.go                # JWT token generation/validation
│       ├── password.go           # Password hashing
│       └── validator.go          # Request validation
├── pkg/                          # Public libraries
│   └── errors/                   # Custom error types
│       └── errors.go             # Error definitions
├── tests/                        # Test files
│   ├── handlers/                 # Handler tests
│   │   ├── auth_handler_test.go
│   │   └── patient_handler_test.go
│   ├── services/                 # Service tests
│   │   ├── auth_service_test.go
│   │   └── patient_service_test.go
│   └── testdata/                 # Test data
│       └── mock_responses.json   # Mock API responses
├── migrations/                   # SQL migration files
│   ├── 001_create_staff_table.sql
│   └── 002_create_patients_table.sql
├── docker/                       # Docker configuration
│   ├── Dockerfile                # Go application container
│   └── nginx.conf                # Nginx configuration
├── docs/                         # Documentation
│   ├── api_spec.md               # API specification
│   ├── er_diagram.md             # ER diagram
│   └── project_structure.md      # This file
├── .env                          # Environment variables
├── .env.example                  # Environment template
├── .gitignore                    # Git ignore file
├── docker-compose.yml            # Docker services setup
├── go.mod                        # Go modules
├── go.sum                        # Go dependencies
├── Makefile                      # Build commands
└── README.md                     # Project documentation
```

## Architecture

The HMS project follows a clean layered architecture pattern with clear separation of concerns:

```
┌─────────────────────────────────────────────────────┐
│                  HTTP Requests                      │
└───────────────────────┬─────────────────────────────┘
                        ▼
┌─────────────────────────────────────────────────────┐
│             Middleware (auth, CORS, logging)        │
└───────────────────────┬─────────────────────────────┘
                        ▼
┌─────────────────────────────────────────────────────┐
│         Handlers (Presentation Layer)               │
│                                                     │
│  - Parse and validate HTTP requests                 │
│  - Call appropriate services                        │
│  - Format and return HTTP responses                 │
└───────────────────────┬─────────────────────────────┘
                        ▼
┌─────────────────────────────────────────────────────┐
│         Services (Business Logic Layer)             │
│                                                     │
│  - Authentication and authorization                 │
│  - Patient search logic                             │
│  - Integration with external Hospital APIs          │
│  - Data validation and transformation               │
└───────────────────────┬─────────────────────────────┘
                        ▼
┌─────────────────────────────────────────────────────┐
│         Repositories (Data Access Layer)            │
│                                                     │
│  - CRUD operations for each entity                  │
│  - Query construction and execution                 │
│  - Data mapping between database and domain models  │
└───────────────────────┬─────────────────────────────┘
                        ▼
┌─────────────────────────────────────────────────────┐
│                    Database                         │
└─────────────────────────────────────────────────────┘
```

### 1. Presentation Layer (Handlers)

The handlers are responsible for:
- Parsing and validating HTTP requests
- Calling appropriate services
- Formatting and returning HTTP responses
- Route registration and configuration

**Key Files:** `auth_handler.go`, `patient_handler.go`, `routes.go`

### 2. Business Logic Layer (Services)

The services implement the core business logic:
- Authentication and authorization
- Patient search logic
- Integration with external Hospital APIs
- Data validation and transformation

**Key Files:** `auth_service.go`, `patient_service.go`, `hospital_api_service.go`

### 3. Data Access Layer (Repositories)

The repositories handle database operations:
- CRUD operations for each entity
- Query construction and execution
- Data mapping between database and domain models
- Transaction management

**Key Files:** `base_repository.go`, `staff_repository.go`, `patient_repository.go`

### 4. Domain Layer (Models)

The models represent the business entities:
- Staff and Patient entities
- Data Transfer Objects (DTOs) for requests and responses
- API response formatting

**Key Files:** `staff.go`, `patient.go`, `response.go`

### 5. Infrastructure

Infrastructure components include:
- Configuration management
- Database connection and migrations
- Middleware for authentication, CORS, and logging
- Utility functions for JWT, password hashing, etc.

**Key Files:** `config.go`, `connection.go`, `migrations.go`, `jwt.go`, `password.go`

## Dependency Flow

The dependencies flow inward, with outer layers depending on inner layers:

```
Handlers → Services → Repositories → Database
```

This ensures that:
- Business logic is isolated from HTTP concerns
- Data access is abstracted behind repository interfaces
- Each layer has a single responsibility

## Design Patterns

The HMS project uses several design patterns:

1. **Repository Pattern**: Abstracts data access behind interfaces
2. **Dependency Injection**: Services and repositories are injected into handlers
3. **Middleware Pattern**: For cross-cutting concerns like authentication
4. **DTO Pattern**: For data transfer between layers

## Error Handling

The project uses a centralized error handling approach:
- Custom error types in the `pkg/errors` package
- Error wrapping for context preservation
- Consistent error responses in API handlers

## Configuration

Configuration is managed through:
- Environment variables
- `.env` file for local development
- Docker environment variables for containerized deployment
