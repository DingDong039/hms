-- Up migration: create staff table and index
CREATE TABLE IF NOT EXISTS staff (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(username)
);

-- Create index for faster lookups
CREATE INDEX IF NOT EXISTS idx_staff_username ON staff(username);
