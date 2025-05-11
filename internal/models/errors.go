package models

import (
	"errors"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrUnknownStatus      = errors.New("todos: error parsing status string")
	ErrUnknownType        = errors.New("todos: error parsing type string")
	ErrMaxTodos           = errors.New("todos: maximum number of todos reached")
	ErrDuplicateEmail     = errors.New("users: duplicate email found")
	ErrInvalidCredentials = errors.New("users: invalid credentials")
	ErrUnknownAuth        = errors.New("users: unexpected auth error")
)
