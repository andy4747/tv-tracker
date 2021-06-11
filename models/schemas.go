package models

var usersTable = `
CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    created_at TEXT NOT NULL,
    updated_at TEXT,
    email VARCHAR(320) NOT NULL UNIQUE,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);
`

const tokensTable = `
CREATE TABLE IF NOT EXISTS tokens (
    id BIGSERIAL NOT NULL,
    created_at TEXT NOT NULL,
    updated_at TEXT,
    token VARCHAR(32) UNIQUE NOT NULL,
    user_id BIGINT UNIQUE NOT NULL,
    CONSTRAINT tokens_pk PRIMARY KEY (id),
    CONSTRAINT users_fk FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
`

const moviesTable = `
CREATE TABLE IF NOT EXISTS movies (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGSERIAL NOT NULL,
    created_at TEXT NOT NULL,
    updated_at TEXT,
    name VARCHAR(255) NOT NULL,
    status VARCHAR(1) NOT NULL,
    current_length BIGINT DEFAULT 0,
    year DATE DEFAULT '1970-01-01',
    total_length BIGINT DEFAULT 0,
    language VARCHAR(255) DEFAULT 'english',
    CONSTRAINT users_movie_fk FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
`
