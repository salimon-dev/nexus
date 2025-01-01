CREATE TABLE users (
    id UUID PRIMARY KEY,
    username VARCHAR(32) UNIQUE NOT NULL,
    email VARCHAR(32) UNIQUE NOT NULL,
    password VARCHAR(32) NOT NULL,
    credit NUMERIC,
    usage NUMERIC,
    status NUMERIC,
    role VARCHAR(16) NOT NULL,
    registered_at DATE,
    updated_at DATE
);
