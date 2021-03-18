package entities

import "time"

type Temperature struct {
	room       string
	sampleDate time.Time
	value      float64
}

func NewTemperature(
	room string,
	sampleDate time.Time,
	value float64,
) (*Temperature, error) {
	if room == "" {
		return nil, ErrEmptyRoom
	}

	return &Temperature{
		room:       room,
		sampleDate: sampleDate,
		value:      value,
	}, nil
}

func (t *Temperature) Room() string {
	return t.room
}

func (t *Temperature) Value() float64 {
	return t.value
}
