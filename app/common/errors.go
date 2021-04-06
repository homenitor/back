package common

import "errors"

var (
	ErrUnknown = errors.New("unknown")

	ErrNilRepository = errors.New("repository is nil")
	ErrNilLogging    = errors.New("logging is nil")
	ErrNilProbes     = errors.New("probes library is nil")
)
