-- +goose Up
-- +goose StatementBegin

CREATE TYPE user_gender AS ENUM ('male', 'female');
CREATE TYPE user_status AS ENUM ('active', 'inactive', 'banned', 'pending');

CREATE TABLE users (
    id UUID PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    surname VARCHAR(100) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    gender user_gender NOT NULL,
    birth_date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status user_status NOT NULL
);

CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_id ON users(id);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
DROP TYPE user_status;
DROP TYPE user_gender
-- +goose StatementEnd
