-- +goose Up
CREATE TABLE IF NOT EXISTS password_reset_tokens (
	id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
	email VARCHAR(255) NOT NULL,
	token VARCHAR(64) NOT NULL UNIQUE,
	expires DATETIME NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS password_reset_tokens;
