-- users_table.sql
-- Create users table
CREATE TYPE customrole AS ENUM ('doctor', 'user');
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       email VARCHAR(30) NOT NULL,
                       is_email_verified boolean default false,
                       role customrole NOT NULL DEFAULT 'user',
                       created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
                       updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
                       deleted_at TIMESTAMP,
                       CONSTRAINT unique_email UNIQUE (email)
);
