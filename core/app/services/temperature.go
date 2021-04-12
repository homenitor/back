package services

import (
	"time"

	"github.com/homenitor/back/core/entities"
	"github.com/homenitor/back/core/values"
)

func (s *Service) SaveTemperature(probeID int, date time.Time, value float64) error {
	return s.saveSample(probeID, values.TEMPERATURE_SAMPLE_CATEGORY, date, value)
}

func (s *Service) GetLastTemperature(probeID int) (*entities.Sample, error) {
	return s.getLastSample(probeID, values.TEMPERATURE_SAMPLE_CATEGORY)
}
