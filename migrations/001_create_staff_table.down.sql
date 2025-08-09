-- Down migration: drop staff index and table
DROP INDEX IF EXISTS idx_staff_username;
DROP TABLE IF EXISTS staff;
