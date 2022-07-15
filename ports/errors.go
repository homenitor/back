package ports

import "errors"

var (
	ErrNilProbeID = errors.New("probe id is nil")
  ErrUnknownSampleCategory = errors.New("unknown sample category")
)
