package domain

import "errors"

var (
	ErrInvalidData = errors.New("invalid data")
	ErrUniqueData  = errors.New("unique data")
)
