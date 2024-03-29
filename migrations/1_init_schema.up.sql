CREATE TABLE IF NOT EXISTS todos
(
    id         SERIAL PRIMARY KEY,
    created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    task       TEXT NOT NULL,
    done       BOOLEAN DEFAULT FALSE
);

