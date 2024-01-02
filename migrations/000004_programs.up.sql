-- programs_table.sql
-- Create programs table
CREATE TYPE program_type AS ENUM ('weight_loss', 'stress_work');

CREATE TABLE programs (
                          id SERIAL PRIMARY KEY,
                          type program_type NOT NULL
);
