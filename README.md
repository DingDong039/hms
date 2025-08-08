# Hospital Middleware System (HMS)

A middleware system for hospital information management that provides APIs to search and display patient information from Hospital Information Systems (HIS).

## Features

- Patient search API with hospital data restriction
- Staff authentication with JWT
- Integration with external Hospital APIs
- Database caching of patient information
- Docker containerization for easy deployment

## Tech Stack

- **Backend**: Go with Gin Framework
- **Database**: PostgreSQL
- **Web Server**: Nginx
- **Containerization**: Docker
- **Authentication**: JWT

## Project Structure

```
HMS/
├── cmd/
│  └── main/
│      └── main.go                 # Entry point
├── internal/
│   ├── config/
│   │   └── config.go              # Configuration management
│   ├── handlers/
│   │   ├── auth_handler.go        # Staff create/login endpoints
│   │   ├── patient_handler.go     # Patient search endpoint
│   │   └── routes.go              # Route registration
│   ├── services/
│   │   ├── auth_service.go        # Authentication logic
│   │   ├── patient_service.go     # Patient business logic
│   │   └── hospital_api_service.go # Hospital A API client
│   ├── repositories/
│   │   ├── base_repository.go     # Base repository pattern
│   │   ├── staff_repository.go    # Staff database operations
│   │   └── patient_repository.go  # Patient database operations
│   ├── models/
│   │   ├── staff.go              # Staff model
│   │   ├── patient.go            # Patient model
│   │   └── response.go           # API response models
│   ├── middleware/
│   │   ├── auth_middleware.go    # JWT validation
│   │   ├── cors_middleware.go    # CORS handling
│   │   └── logging_middleware.go # Request logging
│   ├── database/
│   │   ├── connection.go         # Database connection
│   │   └── migrations.go         # Database migrations
│   └── utils/
│       ├── jwt.go               # JWT utilities
│       ├── password.go          # Password hashing
│       └── validator.go         # Input validation
├── pkg/
│   └── errors/
│       └── errors.go            # Custom error types
├── tests/
│   ├── handlers/
│   │   ├── auth_handler_test.go
│   │   └── patient_handler_test.go
│   ├── services/
│   │   ├── auth_service_test.go
│   │   └── patient_service_test.go
│   └── testdata/
│       └── mock_responses.json   # Mock Hospital A API responses
├── migrations/
│   ├── 001_create_staff_table.sql
│   └── 002_create_patients_table.sql
├── docker/
│   ├── Dockerfile
│   └── nginx.conf               # Nginx config
├── docs/
│   ├── api_spec.md             # API Documentation
│   ├── er_diagram.png          # ER Diagram
│   └── project_structure.md    # Project Structure Doc
├── .env                        # Environment variables
├── .env.example               # Environment template
├── .gitignore
├── docker-compose.yml         # Docker services setup
├── go.mod                     # Go modules
├── go.sum                     # Go dependencies
└── README.md                  # Project documentation
```

## API Endpoints

### Authentication

- `POST /api/v1/auth/staff/create`: Create a new staff member
- `POST /api/v1/auth/staff/login`: Login and get JWT token

### Patient

- `POST /api/v1/patients/search`: Search for a patient by ID (requires authentication)

## Getting Started

### Prerequisites

- Docker and Docker Compose
- Go 1.21+ (for local development)
- PostgreSQL (for local development)

### Running with Docker

1. Clone the repository
2. Configure environment variables in `.env` file
3. Run the application:

```bash
docker-compose up -d
```

### Running Locally

1. Clone the repository
2. Configure environment variables in `.env` file
3. Install dependencies:

```bash
go mod download
```

4. Run the application:

```bash
go run cmd/main/main.go
```

## Database Schema

### Staff Table

- `id`: Primary key
- `username`: Staff username
- `password`: Hashed password
- `hospital_id`: Hospital ID
- `created_at`: Creation timestamp
- `updated_at`: Update timestamp

### Patients Table

- `id`: Primary key
- `national_id`: Thai national ID
- `passport_id`: Passport ID for foreigners
- `first_name_th`: First name in Thai
- `middle_name_th`: Middle name in Thai
- `last_name_th`: Last name in Thai
- `first_name_en`: First name in English
- `middle_name_en`: Middle name in English
- `last_name_en`: Last name in English
- `date_of_birth`: Date of birth
- `patient_hn`: Hospital number
- `phone_number`: Phone number
- `email`: Email address
- `gender`: Gender (M/F)
- `hospital_id`: Hospital ID
- `created_at`: Creation timestamp
- `updated_at`: Update timestamp

## License

