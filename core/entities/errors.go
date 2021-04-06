package entities

import "errors"

var (
	ErrEmptyId   = errors.New("empty id")
	ErrEmptyRoom = errors.New("empty room")
)
