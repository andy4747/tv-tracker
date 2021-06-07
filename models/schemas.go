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
    user_id BIGINT UNIQUE NOT NULL,
    CONSTRAINT tokens_pk PRIMARY KEY (id),
    CONSTRAINT users_fk FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
`

const moviesTable = `
CREATE TABLE IF NOT EXISTS movies (
    id BIGSERIAL NOT NULL,
    created_at TEXT NOT NULL,
    updated_at TEXT,
    name VARCHAR(255) NOT NULL,
    year DATE NOT NULL,
    length INTEGER NOT NULL,
    language VARCHAR(55) NOT NULL 
);
`
const directorsTable = `
CREATE TABLE IF NOT EXISTS directors(
    id BIGSERIAL PRIMARY KEY,
    created_at TEXT NOT NULL,
    updated_at TEXT,
    first_name VARCHAR(55) NOT NULL,
    last_name VARCHAR(55) NOT NULL,
    nationality VARCHAR(55) NOT NULL,
    dob DATE NOT NULL,
    gender VARCHAR(6) NOT NULL
);
`

const genresTable = `
CREATE TABLE IF NOT EXISTS genres (
    id BIGSERIAL PRIMARY KEY,
    created_at TEXT NOT NULL,
    updated_at TEXT,
    name VARCHAR(25)
);
`

const movieGenreTable = `
CREATE TABLE IF NOT EXISTS movie_genres (
    movie_id BIGSERIAL NOT NULL,
    genre_id BIGSERIAL NOT NULL,
    created_at TEXT NOT NULL,
    updated_at TEXT,
    CONSTRAINT movie_fk FOREIGN KEY (movie_id) REFERENCES movies(id) ON DELETE CASCADE,
    CONSTRAINT genre_fk FOREIGN KEY (genre_id) REFERENCES genres(id) ON DELETE CASCADE,
    PRIMARY KEY(movie_id, genre_id)
);
`
