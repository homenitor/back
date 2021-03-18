package app

import "errors"

var (
	ErrUnknown = errors.New("unknown")

	ErrNilRepository = errors.New("repository is nil")
	ErrNilLogging    = errors.New("logging is nil")
)
