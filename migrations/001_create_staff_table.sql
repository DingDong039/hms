-- Create staff table
CREATE TABLE IF NOT EXISTS staff (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(255) NOT NULL,
    hospital_id INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(username, hospital_id)
);

-- Create index for faster lookups
CREATE INDEX IF NOT EXISTS idx_staff_username ON staff(username);
CREATE INDEX IF NOT EXISTS idx_staff_hospital_id ON staff(hospital_id);

-- Create hospitals table
CREATE TABLE IF NOT EXISTS hospitals (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Insert some sample hospitals
INSERT INTO hospitals (name) VALUES ('Hospital A') ON CONFLICT DO NOTHING;
INSERT INTO hospitals (name) VALUES ('Hospital B') ON CONFLICT DO NOTHING;

-- Down migration
-- DROP TABLE IF EXISTS staff;
-- DROP TABLE IF EXISTS hospitals;
