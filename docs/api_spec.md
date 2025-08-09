# Hospital Middleware System API Specification

## Overview

This document outlines the API endpoints provided by the Hospital Middleware System (HMS). The API allows authenticated staff members to search for patient information across connected hospital systems.

**Related Documentation:**
- [README](../README.md) - Main project documentation
- [ER Diagram](./er_diagram.md) - Database schema and relationships
- [Project Structure](./project_structure.md) - Code organization

## Base URL

**Development Environment:**
```
http://localhost:8080/api/v1
```

**Production Environment:**
```
https://api.hms.example.com/api/v1
```

## Authentication

The API uses JWT (JSON Web Token) for authentication. 

### Authentication Flow

1. Staff member logs in using username and password via `/auth/staff/login`
2. System validates credentials and returns a JWT token
3. Client includes this token in subsequent requests
4. System validates token for each protected endpoint

### Token Usage

To access protected endpoints, include the token in the Authorization header:

```
Authorization: Bearer <token>
```

### Token Expiration

Tokens expire after 24 hours. The expiration timestamp is included in the login response.

## API Response Format

All API responses follow a standard format:

### Success Response

```json
{
  "success": true,
  "data": { /* response data */ }
}
```

### Error Response

```json
{
  "success": false,
  "error": {
    "code": 400,
    "message": "Error message"
  }
}
```

## Endpoints

### Health Check

**GET /health**

Check if the API is running.

**Request**

```bash
curl -X GET http://localhost:8080/api/v1/health
```

**Response**

```json
{
  "success": true,
  "data": {
    "status": "ok",
    "version": "1.0.0",
    "timestamp": "2025-08-09T13:34:04Z"
  }
}
```

### Authentication

#### Create Staff

**POST /auth/staff/create**

Create a new staff member. Requires staff privileges.

**Request Headers**

```
Content-Type: application/json
Authorization: Bearer <staff_token>
```

**Request Body**

```json
{
  "username": "staffuser",
  "password": "password123",
}
```

**Request Example**

```bash
curl -X POST http://localhost:8080/api/v1/auth/staff/create \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <staff_token>" \
  -d '{"username":"staffuser","password":"password123"}'
```

**Response**

```json
{
  "success": true,
  "data": {
    "id": 1,
    "username": "staffuser",
    "created_at": "2025-08-08T12:00:00Z",
    "updated_at": "2025-08-08T12:00:00Z"
  }
}
```

**Validation Rules**

- `username`: Required, must be unique within the hospital
- `password`: Required, minimum 8 characters

#### Staff Login

**POST /auth/staff/login**

Authenticate a staff member and receive a JWT token.

**Request Headers**

```
Content-Type: application/json
```

**Request Body**

```json
{
  "username": "staffuser",
  "password": "password123"
}
```

**Request Example**

```bash
curl -X POST http://localhost:8080/api/v1/auth/staff/login \
  -H "Content-Type: application/json" \
  -d '{"username":"staffuser","password":"password123"}'
```

**Response**

```json
{
  "success": true,
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "expires_at": "2025-08-10T13:34:04Z"
  }
}
```

**Error Responses**

- **Invalid Credentials**:
```json
{
  "success": false,
  "error": {
    "code": 401,
    "message": "Invalid username or password"
  }
}
```

### Patient Endpoints

#### Search Patient

**POST /patients/search**

Search for a patient by ID (national ID or passport ID). Requires authentication. Staff can only access patients from their own hospital.

**Request Headers**

```
Content-Type: application/json
Authorization: Bearer <token>
```

**Request Body**

```json
{
  "id_type": "national_id",
  "id": "1234567890123"
}
```

The `id_type` field can be either `national_id` or `passport_id`.

**Request Example**

```bash
curl -X POST http://localhost:8080/api/v1/patients/search \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -d '{"id_type":"national_id","id":"1234567890123"}'
```

**Success Response**

```json
{
  "success": true,
  "data": {
    "first_name_th": "สมชาย",
    "middle_name_th": "",
    "last_name_th": "ใจดี",
    "first_name_en": "Somchai",
    "middle_name_en": "",
    "last_name_en": "Jaidee",
    "date_of_birth": "1990-01-01T00:00:00Z",
    "patient_hn": "HN12345",
    "national_id": "1234567890123",
    "passport_id": "",
    "phone_number": "0812345678",
    "email": "somchai@example.com",
    "gender": "M"
  }
}
```

**Error Responses**

- **Patient Not Found**:
```json
{
  "success": false,
  "error": {
    "code": 404,
    "message": "Patient not found"
  }
}
```

- **Unauthorized Access**:
```json
{
  "success": false,
  "error": {
    "code": 403,
    "message": "invalid or expired token"
  }
}
```
```

## Error Handling

### Error Response Format

All error responses follow a standard format:

```json
{
  "success": false,
  "error": {
    "code": 400,
    "message": "Error message",
    "details": {} // Optional additional error details
  }
}
```

### Common Error Codes

| Code | Status | Description | Example |
|------|--------|-------------|----------|
| 400 | Bad Request | Request validation failed | Invalid input format, missing required fields |
| 401 | Unauthorized | Authentication failed | Invalid or expired token |
| 403 | Forbidden | Permission denied | Staff attempting to access data from another hospital |
| 404 | Not Found | Resource not found | Patient not found, endpoint not found |
| 422 | Unprocessable Entity | Semantic errors | Data validation errors |
| 429 | Too Many Requests | Rate limit exceeded | Too many requests in a given time |
| 500 | Internal Server Error | Server-side error | Database connection failure |

### Validation Error Example

```json
{
  "success": false,
  "error": {
    "code": 400,
    "message": "Validation failed",
    "details": {
      "id": "ID must be 13 digits for national_id"
    }
  }
}
```

## External API Integration

The HMS integrates with external hospital APIs to retrieve patient information. This section describes these integrations.

### Hospital A API

**GET /patient/search/{id}**

- **Base URL**: `https://hospital-a.api.co.th/api/v2`
- **Authentication**: API Key in header `X-API-Key: <key>`
- **Parameters**: 
  - `id`: Can be either national_id or passport_id
  - `id_type`: Specify either `national` or `passport`
- **Request Example**:
  ```bash
  curl -X GET "https://hospital-a.api.co.th/api/v2/patient/search/1234567890123?id_type=national" \
    -H "X-API-Key: your-api-key"
  ```
- **Response Example**:
  ```json
  {
    "status": "success",
    "patient": {
      "first_name_th": "สมชาย",
      "last_name_th": "ใจดี",
      "first_name_en": "Somchai",
      "last_name_en": "Jaidee",
      "hn": "HN12345",
      "national_id": "1234567890123",
      "dob": "1990-01-01",
      "gender": "male",
      "contact": {
        "phone": "0812345678",
        "email": "somchai@example.com"
      }
    }
  }
  ```

### Hospital B API

**POST /api/patients/find**

- **Base URL**: `https://hospital-b-api.healthcare.org/api`
- **Authentication**: OAuth 2.0
- **Request Body**:
  ```json
  {
    "identifier": {
      "type": "national_id",
      "value": "1234567890123"
    }
  }
  ```
- **Response**: Patient information in FHIR-compatible format

## Data Models

### Staff

```json
{
  "id": 1,
  "username": "doctor.smith",
  "created_at": "2023-01-01T00:00:00Z",
  "updated_at": "2023-01-01T00:00:00Z"
}
```

### Patient

```json
{
  "id": 1,
  "national_id": "1234567890123",
  "passport_id": null,
  "first_name_th": "สมชาย",
  "middle_name_th": null,
  "last_name_th": "ใจดี",
  "first_name_en": "Somchai",
  "middle_name_en": null,
  "last_name_en": "Jaidee",
  "date_of_birth": "1990-01-01T00:00:00Z",
  "patient_hn": "HN12345",
  "phone_number": "0812345678",
  "email": "somchai@example.com",
  "gender": "M",
  "created_at": "2023-01-01T00:00:00Z",
  "updated_at": "2023-01-01T00:00:00Z"
}
```

## Rate Limiting

To ensure system stability, the API implements rate limiting:

- **Unauthenticated requests**: 60 requests per hour per IP address
- **Authenticated requests**: 1000 requests per hour per staff member

When the rate limit is exceeded, the API returns a 429 Too Many Requests response.

## Versioning

The API uses URL versioning (e.g., `/api/v1/`). When breaking changes are introduced, a new version will be released.

## Security Considerations

- All API requests must use HTTPS
- Passwords are never returned in responses
- JWT tokens have a 24-hour expiration
- Staff can only access patient data from their own hospital

## Changelog

### v1.0.0 (2025-08-01)
- Initial release with authentication and patient search

### v1.1.0 (Planned)
- Add patient registration endpoint
- Add hospital management endpoints
