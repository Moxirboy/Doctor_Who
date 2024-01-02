-- users_table.sql
-- Create users table
CREATE TYPE customrole AS ENUM ('doctor', 'user');
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       phone_number VARCHAR(15) NOT NULL,
                       password VARCHAR(128) NOT NULL,
                       role customrole NOT NULL DEFAULT 'user',
                       created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
                       updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
                       deleted_at TIMESTAMP,
                       CONSTRAINT unique_phone_number UNIQUE (phone_number)
);
