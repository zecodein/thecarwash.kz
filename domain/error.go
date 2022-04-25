package domain

import "errors"

var (
	ErrInvalidData = errors.New("invalid data")
	ErrInvalidIin  = errors.New("invalid iin")
	ErrUniqueData  = errors.New("unique data")
)
