package entities

import (
	"time"

	"github.com/homenitor/back/core/values"
)

type Sample struct {
	category  values.SampleCategory
	timestamp time.Time
	value     float64
}

func NewSample(
	category values.SampleCategory,
	timestamp time.Time,
	value float64,
) (*Sample, error) {
	return &Sample{
		category:  category,
		timestamp: timestamp,
		value:     value,
	}, nil
}

func (s *Sample) Value() float64 {
	return s.value
}

func (s *Sample) Category() values.SampleCategory {
	return s.category
}

func (s *Sample) Timestamp() time.Time {
	return s.timestamp
}
