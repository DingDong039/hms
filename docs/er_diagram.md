# Hospital Middleware System ER Diagram

## Database Schema

```
+------------------+       +------------------+       +------------------+
|     Hospitals    |       |      Staff       |       |     Patients     |
+------------------+       +------------------+       +------------------+
| id (PK)          |       | id (PK)          |       | id (PK)          |
| name             |<----->| hospital_id (FK) |       | national_id      |
| code             |       | username         |       | passport_id      |
| created_at       |       | password         |       | first_name_th    |
| updated_at       |       | created_at       |       | middle_name_th   |
+------------------+       | updated_at       |       | last_name_th     |
                           +------------------+       | first_name_en    |
                                                     | middle_name_en    |
                                                     | last_name_en      |
                                                     | date_of_birth     |
                                                     | patient_hn        |
                                                     | phone_number      |
                                                     | email             |
                                                     | gender            |
                                                     | hospital_id (FK)  |
                                                     | created_at        |
                                                     | updated_at        |
                                                     +------------------+
```

## Relationships

1. **Hospital to Staff**: One-to-Many
   - One hospital can have multiple staff members
   - Each staff member belongs to exactly one hospital
   - Foreign key: `staff.hospital_id` references `hospitals.id`

2. **Hospital to Patients**: One-to-Many
   - One hospital can have multiple patients
   - Each patient belongs to exactly one hospital
   - Foreign key: `patients.hospital_id` references `hospitals.id`

## Indexes

### Hospitals Table
- Primary Key: `id`
- Unique Index: `code`

### Staff Table
- Primary Key: `id`
- Unique Index: `(username, hospital_id)`
- Index: `hospital_id`

### Patients Table
- Primary Key: `id`
- Unique Index: `(national_id, hospital_id)` when `national_id` is not null
- Unique Index: `(passport_id, hospital_id)` when `passport_id` is not null
- Index: `hospital_id`
- Index: `national_id`
- Index: `passport_id`
- Index: `patient_hn`

## Data Types

### Hospitals Table
- `id`: SERIAL (auto-incrementing integer)
- `name`: VARCHAR(255)
- `code`: VARCHAR(50)
- `created_at`: TIMESTAMP WITH TIME ZONE
- `updated_at`: TIMESTAMP WITH TIME ZONE

### Staff Table
- `id`: SERIAL (auto-incrementing integer)
- `hospital_id`: INTEGER (foreign key)
- `username`: VARCHAR(100)
- `password`: VARCHAR(255) (bcrypt hashed)
- `created_at`: TIMESTAMP WITH TIME ZONE
- `updated_at`: TIMESTAMP WITH TIME ZONE

### Patients Table
- `id`: SERIAL (auto-incrementing integer)
- `national_id`: VARCHAR(13)
- `passport_id`: VARCHAR(50)
- `first_name_th`: VARCHAR(100)
- `middle_name_th`: VARCHAR(100)
- `last_name_th`: VARCHAR(100)
- `first_name_en`: VARCHAR(100)
- `middle_name_en`: VARCHAR(100)
- `last_name_en`: VARCHAR(100)
- `date_of_birth`: TIMESTAMP WITH TIME ZONE
- `patient_hn`: VARCHAR(50)
- `phone_number`: VARCHAR(20)
- `email`: VARCHAR(100)
- `gender`: CHAR(1)
- `hospital_id`: INTEGER (foreign key)
- `created_at`: TIMESTAMP WITH TIME ZONE
- `updated_at`: TIMESTAMP WITH TIME ZONE

## Constraints

### Hospitals Table
- `id` is the primary key and auto-increments
- `code` must be unique

### Staff Table
- `id` is the primary key and auto-increments
- The combination of `username` and `hospital_id` must be unique
- `hospital_id` must reference a valid hospital

### Patients Table
- `id` is the primary key and auto-increments
- At least one of `national_id` or `passport_id` must be provided
- If `national_id` is provided, the combination of `national_id` and `hospital_id` must be unique
- If `passport_id` is provided, the combination of `passport_id` and `hospital_id` must be unique
- `hospital_id` must reference a valid hospital
- `gender` must be either 'M' or 'F'
