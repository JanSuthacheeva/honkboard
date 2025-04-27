package models

import (
	"errors"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrUnknownStatus      = errors.New("todos: error parsing status string")
	ErrUnknownType        = errors.New("todos: error parsing type string")
	ErrDuplicateEmail     = errors.New("users: duplicate email found")
	ErrInvalidCredentials = errors.New("users: invalid credentials")
)
