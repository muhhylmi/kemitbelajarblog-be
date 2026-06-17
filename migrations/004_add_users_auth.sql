-- 004_add_users_auth.sql
-- Add username and password_hash to authors table to use them as users

ALTER TABLE authors ADD COLUMN IF NOT EXISTS username VARCHAR(100) UNIQUE;
ALTER TABLE authors ADD COLUMN IF NOT EXISTS password_hash VARCHAR(255);

-- We need to generate a default password for existing authors.
-- bcrypt hash for 'password123' is $2a$10$wE8w.6uY.oP4P.W4O/uY.eV/g0V.bHqZ8QxZg1kC3t.9/rXwHl.V.
-- So we will assign a default username based on their name, and this password hash.

-- Give Elias Thorne the username 'elias'
UPDATE authors 
SET username = 'elias', 
    password_hash = '$2a$10$wE8w.6uY.oP4P.W4O/uY.eV/g0V.bHqZ8QxZg1kC3t.9/rXwHl.V.' 
WHERE id = 'a0000000-0000-0000-0000-000000000001';

-- Give Julian Thorne the username 'julian'
UPDATE authors 
SET username = 'julian', 
    password_hash = '$2a$10$wE8w.6uY.oP4P.W4O/uY.eV/g0V.bHqZ8QxZg1kC3t.9/rXwHl.V.' 
WHERE id = 'a0000000-0000-0000-0000-000000000002';
