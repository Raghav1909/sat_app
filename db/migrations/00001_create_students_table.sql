-- +goose Up
-- +goose StatementBegin
CREATE TABLE students (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) UNIQUE NOT NULL,
    address VARCHAR(1000) NOT NULL,
    city VARCHAR(100) NOT NULL,
    country VARCHAR(100) NOT NULL,
    pincode VARCHAR(6) NOT NULL,
    sat_score SMALLINT NOT NULL,
    passed BOOLEAN NOT NULL DEFAULT false
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE students;
-- +goose StatementEnd

