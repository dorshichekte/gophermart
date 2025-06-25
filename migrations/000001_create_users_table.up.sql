CREATE TABLE IF NOT EXISTS users (
    id  BIGSERIAL PRIMARY KEY,
    login VARCHAR(255) UNIQUE NOT NULL CHECK (length(login) >= 4),
    password VARCHAR(255) NOT NULL CHECK (length(password) >= 6),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    CHECK (modified_at >= created_at)
);