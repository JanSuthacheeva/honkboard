-- +goose Up
CREATE TABLE IF NOT EXISTS todos (
	id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
	title VARCHAR(255) NOT NULL,
	status VARCHAR(40) NOT NULL,
	type VARCHAR(40) NOT NULL,
	created DATETIME NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS todos;
