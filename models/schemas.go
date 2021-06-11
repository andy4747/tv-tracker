package models

var usersTable = `
CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    created_at VARCHAR(55) NOT NULL,
    updated_at VARCHAR(55),
    email VARCHAR(320) NOT NULL UNIQUE,
    username VARCHAR(55) NOT NULL,
    password VARCHAR(55) NOT NULL
);
`

const tokensTable = `
CREATE TABLE IF NOT EXISTS tokens (
    id BIGSERIAL NOT NULL,
    created_at VARCHAR(55) NOT NULL,
    updated_at VARCHAR(55),
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
    created_at VARCHAR(55) NOT NULL,
    updated_at VARCHAR(55),
    name VARCHAR(255) NOT NULL,
    status STATUS NOT NULL,
    current_length INTEGER,
    year DATE,
    total_length INTEGER,
    language VARCHAR(55),
    CONSTRAINT users_movie_fk FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
`

// const statusType = `
//     CREATE TYPE STATUS AS ENUM ('c','w','p');
// `
