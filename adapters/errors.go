package adapters

import "errors"

var (
	ErrNilMqttClient         = errors.New("mqtt client is nil")
	ErrNilMqttServer         = errors.New("mqtt server is nil")
	ErrUnknownLogLevel       = errors.New("unknown log level")
	ErrUnknownSampleCategory = errors.New("unknown sample category")
)
