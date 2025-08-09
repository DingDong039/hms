-- Up migration: create patients table and indexes
CREATE TABLE IF NOT EXISTS patients (
    id SERIAL PRIMARY KEY,
    national_id VARCHAR(20),
    passport_id VARCHAR(20),
    first_name_th VARCHAR(100),
    middle_name_th VARCHAR(100),
    last_name_th VARCHAR(100),
    first_name_en VARCHAR(100),
    middle_name_en VARCHAR(100),
    last_name_en VARCHAR(100),
    date_of_birth DATE,
    patient_hn VARCHAR(50) NOT NULL,
    phone_number VARCHAR(20),
    email VARCHAR(100),
    gender VARCHAR(1),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT chk_id CHECK (national_id IS NOT NULL OR passport_id IS NOT NULL),
    CONSTRAINT chk_gender CHECK (gender IN ('M', 'F'))
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_patients_national_id ON patients(national_id);
CREATE INDEX IF NOT EXISTS idx_patients_passport_id ON patients(passport_id);
CREATE INDEX IF NOT EXISTS idx_patients_patient_hn ON patients(patient_hn);
