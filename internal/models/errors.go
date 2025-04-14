package models

import (
	"errors"
)

var ErrNoRecord = errors.New("models: no matching record found")
var ErrUnknownStatus = errors.New("models: error parsing status string")
