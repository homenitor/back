package common

import "errors"

var (
	ErrUnknown = errors.New("unknown")

	ErrNilRepository   = errors.New("repository is nil")
	ErrNilLogging      = errors.New("logging is nil")
	ErrNilProbeLibrary = errors.New("probes library is nil")

	ErrNoSampleValueInProbe = errors.New("no value found in probe")

	ErrProbeNotFound = errors.New("probe was not found")
)
