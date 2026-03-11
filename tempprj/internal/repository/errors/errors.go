package errors

import "errors"

var (
	ErrUserAlreadyExists = errors.New("User with this email already exists")
	ErrUserNotFound      = errors.New("User with this email not fount")
)
