# Hospital Middleware System API Specification

## Base URL

```
http://localhost/api/v1
```

## Authentication

The API uses JWT (JSON Web Token) for authentication. To access protected endpoints, include the token in the Authorization header:

```
Authorization: Bearer <token>
```

## Endpoints

### Health Check

**GET /health**

Check if the API is running.

**Response**

```json
{
  "success": true,
  "data": {
    "status": "ok"
  }
}
```

### Authentication

#### Create Staff

**POST /auth/staff/create**

Create a new staff member.

**Request Body**

```json
{
  "username": "staffuser",
  "password": "password123",
  "hospital_id": 1
}
```

**Response**

```json
{
  "success": true,
  "data": {
    "id": 1,
    "username": "staffuser",
    "hospital_id": 1,
    "created_at": "2025-08-08T12:00:00Z",
    "updated_at": "2025-08-08T12:00:00Z"
  }
}
```

#### Staff Login

**POST /auth/staff/login**

Authenticate a staff member and receive a JWT token.

**Request Body**

```json
{
  "username": "staffuser",
  "password": "password123",
  "hospital_id": 1
}
```

**Response**

```json
{
  "success": true,
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "expires_at": 1628432000
  }
}
```

### Patient

#### Search Patient

**POST /patients/search**

Search for a patient by ID (national ID or passport ID). Requires authentication.

**Request Headers**

```
Authorization: Bearer <token>
```

**Request Body**

```json
{
  "id": "1234567890123"
}
```

**Response**

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

## Error Responses

The API returns standardized error responses:

```json
{
  "success": false,
  "error": {
    "code": 400,
    "message": "validation failed"
  }
}
```

### Common Error Codes

- `400` - Bad Request (validation error)
- `401` - Unauthorized (invalid or missing token)
- `403` - Forbidden (insufficient permissions)
- `404` - Not Found (resource not found)
- `500` - Internal Server Error

## External API Integration

### Hospital A API

**GET /patient/search/{id}**

- **Base URL**: https://hospital-a.api.co.th
- **Parameters**: `id` can be either national_id or passport_id
- **Response**: Patient information in JSON format

## Data Models

### Staff

```json
{
  "id": 1,
  "username": "staffuser",
  "hospital_id": 1,
  "created_at": "2025-08-08T12:00:00Z",
  "updated_at": "2025-08-08T12:00:00Z"
}
```

### Patient

```json
{
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
```
