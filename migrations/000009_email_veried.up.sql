
-- Create verify_emails table
CREATE TABLE verify_emails (
                               id SERIAL PRIMARY KEY,
                               user_id INTEGER ,
                               secret_code VARCHAR NOT NULL,
                               is_used BOOLEAN NOT NULL DEFAULT false,
                               created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
                               expire_at TIMESTAMPTZ NOT NULL DEFAULT (now() + INTERVAL '15 minutes'),
                               CONSTRAINT fk_verify_emails_user FOREIGN KEY (user_id) REFERENCES users(id)

);