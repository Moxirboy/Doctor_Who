-- user_info_table.sql
-- Create user_info table
CREATE TABLE user_info (
                           id SERIAL PRIMARY KEY,
                           user_id INT NOT NULL,
                           name VARCHAR(255),
                           weigh VARCHAR(255),
                           height VARCHAR(255),
                           age INT,
                           waist VARCHAR(255),
                           CONSTRAINT fk_user_info_user FOREIGN KEY (user_id) REFERENCES users(id)
);
