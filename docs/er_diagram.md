# Hospital Middleware System ER Diagram

## Overview

This document describes the database schema for the Hospital Middleware System (HMS). The system uses PostgreSQL and consists of three main tables: Hospitals, Staff, and Patients, with relationships between them.

**Related Documentation:**
- [README](../README.md) - Main project documentation
- [API Specification](./api_spec.md) - API endpoints and usage
- [Project Structure](./project_structure.md) - Code organization

## Database Schema Visualization

```
+------------------+       +------------------+
|      Staff       |       |     Patients     |
+------------------+       +------------------+
| id (PK)          |       | id (PK)          |
| username         |       | national_id      |
| password         |       | passport_id      |
| created_at       |       | first_name_th    |
| updated_at       |       | middle_name_th   |
+------------------+       | last_name_th     |
                           | first_name_en    |
                           | middle_name_en   |
                           | last_name_en     |
                           | date_of_birth    |
                           | patient_hn       |
                           | phone_number     |
                           | email            |
                           | gender           |
                           | created_at       |
                           | updated_at       |
                           +------------------+
```

## Entity Relationships

The system no longer uses the Hospitals table. Staff and Patients are now independent entities without a direct relationship to a Hospitals table.

## Indexes

### Hospitals Table
- Primary Key: `id`
- Unique Index: `code`

### Staff Table
- Primary Key: `id`
- Unique Index: `(username)`

### Patients Table
- Primary Key: `id`
- Unique Index: `(national_id)` when `national_id` is not null
- Unique Index: `(passport_id)` when `passport_id` is not null
- Index: `patient_hn`

## Table Specifications

### Staff Table

| Column | Type | Description | Constraints |
|--------|------|-------------|-------------|
| `id` | SERIAL | Primary key | NOT NULL, AUTO INCREMENT |
| `username` | VARCHAR(100) | Staff username | NOT NULL, UNIQUE |
| `password` | VARCHAR(255) | Bcrypt hashed password | NOT NULL |
| `created_at` | TIMESTAMP WITH TIME ZONE | Creation timestamp | NOT NULL, DEFAULT NOW() |
| `updated_at` | TIMESTAMP WITH TIME ZONE | Update timestamp | NOT NULL, DEFAULT NOW() |

### Patients Table

| Column | Type | Description | Constraints |
|--------|------|-------------|-------------|
| `id` | SERIAL | Primary key | NOT NULL, AUTO INCREMENT |
| `national_id` | VARCHAR(13) | Thai national ID | NULL |
| `passport_id` | VARCHAR(50) | Passport ID for foreigners | NULL |
| `first_name_th` | VARCHAR(100) | First name in Thai | NOT NULL |
| `middle_name_th` | VARCHAR(100) | Middle name in Thai | NULL |
| `last_name_th` | VARCHAR(100) | Last name in Thai | NOT NULL |
| `first_name_en` | VARCHAR(100) | First name in English | NOT NULL |
| `middle_name_en` | VARCHAR(100) | Middle name in English | NULL |
| `last_name_en` | VARCHAR(100) | Last name in English | NOT NULL |
| `date_of_birth` | TIMESTAMP WITH TIME ZONE | Date of birth | NOT NULL |
| `patient_hn` | VARCHAR(50) | Hospital number | NOT NULL |
| `phone_number` | VARCHAR(20) | Phone number | NULL |
| `email` | VARCHAR(100) | Email address | NULL |
| `gender` | CHAR(1) | Gender (M/F) | NOT NULL |
| `created_at` | TIMESTAMP WITH TIME ZONE | Creation timestamp | NOT NULL, DEFAULT NOW() |
| `updated_at` | TIMESTAMP WITH TIME ZONE | Update timestamp | NOT NULL, DEFAULT NOW() |

## Business Rules and Constraints

### Staff Table
- `id` is the primary key and auto-increments
- `username` must be unique
- Password must be stored as a bcrypt hash, never in plain text

### Patients Table
- `id` is the primary key and auto-increments
- At least one of `national_id` or `passport_id` must be provided (CHECK constraint)
- `gender` must be either 'M' or 'F' (CHECK constraint)
- `national_id` must be unique when not null
- `passport_id` must be unique when not null

## Migration Scripts

Migration scripts for creating these tables can be found in the `/migrations` directory:

- `001_create_staff_table.sql`
- `002_create_patients_table.sql`

## Database Diagram

For a visual representation of this schema, see the ER diagram image at `docs/er_diagram.png`
