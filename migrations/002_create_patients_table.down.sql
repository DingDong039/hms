-- Down migration: drop patients indexes and table
DROP INDEX IF EXISTS idx_patients_national_id;
DROP INDEX IF EXISTS idx_patients_passport_id;
DROP INDEX IF EXISTS idx_patients_patient_hn;
DROP TABLE IF EXISTS patients;
