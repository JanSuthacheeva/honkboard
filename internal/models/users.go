package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID             int
	Email          string
	HashedPassword string
	Name           string
	CreatedAt      time.Time
}

type UserModel struct {
	DB *sql.DB
}
