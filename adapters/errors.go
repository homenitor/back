package adapters

import "errors"

var (
	ErrRoomNotFound    = errors.New("room not found")
	ErrNilMqttClient   = errors.New("mqtt client is nil")
	ErrUnknownLogLevel = errors.New("unknown log level")
)
