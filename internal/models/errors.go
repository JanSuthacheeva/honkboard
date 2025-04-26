package models

import (
	"errors"
)

var ErrNoRecord = errors.New("models: no matching record found")
var ErrUnknownStatus = errors.New("todos: error parsing status string")
var ErrUnknownType = errors.New("todos: error parsing type string")
var ErrDuplicateEmail = errors.New("users: duplicate email found")
var ErrInvalidCredentials = errors.New("users: invalid credentials")
