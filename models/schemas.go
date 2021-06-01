package models

var usersTable = `
CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    created_at TEXT NOT NULL,
    updated_at TEXT,
    email TEXT NOT NULL UNIQUE,
    username TEXT NOT NULL,
    password TEXT NOT NULL
);
`

const tokensTable = `
CREATE TABLE IF NOT EXISTS tokens (
    id BIGSERIAL NOT NULL,
    created_at TEXT NOT NULL,
    updated_at TEXT,
    token TEXT UNIQUE NOT NULL,
    user_id BIGINT NOT NULL,
    CONSTRAINT tokens_pk PRIMARY KEY (id),
    CONSTRAINT users_fk FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
`
