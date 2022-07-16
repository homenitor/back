package entities

import (
	"time"

	"github.com/homenitor/back/core/values"
)

type Sample struct {
	category    values.SampleCategory
	measured_at time.Time
	value       float64
}

func NewSample(
	category values.SampleCategory,
	timestamp time.Time,
	value float64,
) (*Sample, error) {
	return &Sample{
		category:    category,
		measured_at: timestamp,
		value:       value,
	}, nil
}

func (s *Sample) Value() float64 {
	return s.value
}

func (s *Sample) Category() values.SampleCategory {
	return s.category
}

func (s *Sample) MeasuredAt() time.Time {
	return s.measured_at
}

func (s *Sample) FormattedMeasuredAt() string {
	return s.measured_at.Format(time.RFC3339)
}
