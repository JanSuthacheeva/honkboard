-- +goose Up
CREATE TABLE IF NOT EXISTS sessions (
	token CHAR(43) PRIMARY KEY,
	data BLOB NOT NULL,
	expiry TIMESTAMP(6) NOT NULL
);

CREATE INDEX sessions_expiry_idx ON sessions (expiry);

-- +goose Down

DROP TABLE IF EXISTS sessions;
