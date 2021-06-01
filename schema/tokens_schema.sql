CREATE TABLE IF NOT EXISTS tokens (
    id BIGSERIAL NOT NULL,
    created_at TEXT NOT NULL,
    updated_at TEXT,
    token TEXT UNIQUE NOT NULL,
    user_id BIGINT NOT NULL,
    CONSTRAINT tokens_pk PRIMARY KEY (id),
    CONSTRAINT users_fk FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);