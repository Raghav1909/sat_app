-- +goose Up
-- +goose StatementBegin
CREATE TABLE students (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) UNIQUE NOT NULL,
    address VARCHAR(1000),
    city VARCHAR(100),
    country VARCHAR(100),
    pincode VARCHAR(6),
    sat_score INT,
    passed TINYINT(1) DEFAULT 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE students;
-- +goose StatementEnd

