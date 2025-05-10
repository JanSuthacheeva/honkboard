-- +goose Up
DROP TABLE IF EXISTS validation_codes;

-- +goose Down
CREATE TABLE IF NOT EXISTS validation_codes(
	id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
	user_id INTEGER NOT NULL,
	code INTEGER NOT NULL,
	type VARCHAR(255) NOT NULL,
	expires DATETIME NOT NULL,
	FOREIGN KEY (user_id) REFERENCES users(id)
);
