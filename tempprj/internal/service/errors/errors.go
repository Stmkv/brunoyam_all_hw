package errors

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

const (
	IncorrectFieldValuesError = "incorrect field values: %w"
)
