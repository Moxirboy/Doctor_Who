-- exercise_table.sql
-- Create exercise table
CREATE TABLE exercise (
                          id SERIAL PRIMARY KEY,
                          program_id INT,
                          name VARCHAR(255),
                          info varchar(255),
                          link VARCHAR(255),
                          CONSTRAINT fk_exercise_program FOREIGN KEY (program_id) REFERENCES programs(id)
);
