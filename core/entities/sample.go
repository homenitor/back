package entities

import (
	"time"

	"github.com/homenitor/back/core/values"
)

type Sample struct {
	room       string
	category   values.SampleCategory
	sampleDate time.Time
	value      float64
}

func NewSample(
	room string,
	category values.SampleCategory,
	sampleDate time.Time,
	value float64,
) (*Sample, error) {
	if room == "" {
		return nil, ErrEmptyRoom
	}

	return &Sample{
		room:       room,
		category:   category,
		sampleDate: sampleDate,
		value:      value,
	}, nil
}

func (s *Sample) Room() string {
	return s.room
}

func (s *Sample) Value() float64 {
	return s.value
}

func (s *Sample) Category() values.SampleCategory {
	return s.category
}
