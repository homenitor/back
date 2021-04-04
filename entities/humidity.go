package entities

import "time"

type Humidity struct {
	room       string
	sampleDate time.Time
	value      float64
}

func NewHumidity(
	room string,
	sampleDate time.Time,
	value float64,
) (*Humidity, error) {
	if room == "" {
		return nil, ErrEmptyRoom
	}

	return &Humidity{
		room:       room,
		sampleDate: sampleDate,
		value:      value,
	}, nil
}

func (h *Humidity) Room() string {
	return h.room
}

func (h *Humidity) Value() float64 {
	return h.value
}
