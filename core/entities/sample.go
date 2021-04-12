package entities

import (
	"time"

	"github.com/homenitor/back/core/values"
)

type Sample struct {
	category   values.SampleCategory
	sampleDate time.Time
	value      float64
}

func NewSample(
	category values.SampleCategory,
	sampleDate time.Time,
	value float64,
) (*Sample, error) {
	return &Sample{
		category:   category,
		sampleDate: sampleDate,
		value:      value,
	}, nil
}

func (s *Sample) Value() float64 {
	return s.value
}

func (s *Sample) Category() values.SampleCategory {
	return s.category
}
