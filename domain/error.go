package domain

import "errors"

var (
	ErrInvalidData     = errors.New("invalid data")
	ErrInvalidIinOrBin = errors.New("invalid iin or bin")
	ErrUniqueData      = errors.New("unique data")
)
