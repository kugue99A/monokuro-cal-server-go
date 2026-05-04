package user

import "errors"

var (
	ErrNotFound       = errors.New("user not found")
	ErrEmptyEmail     = errors.New("email is required")
	ErrEmptyName      = errors.New("name is required")
	ErrEmailDuplicate = errors.New("email already exists")
)
