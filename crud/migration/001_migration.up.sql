-- CREATE TABLE users(
--     id UUID PRIMARY KEY,
--     first_name VARCHAR(255),
--     last_name VARCHAR(255),
--     department_name VARCHAR(255),
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
--     );

CREATE TABLE IF NOT EXISTS users(
    id uuid,
    first_name varchar,
    telegram_id int,
    status varchar,
    starttime timestamp
    );
