package services

import (
	"time"

	"github.com/homenitor/back/core/entities"
	"github.com/homenitor/back/core/values"
)

func (s *Service) SaveHumidity(probeID string, date time.Time, value float64) error {
	return s.saveSample(probeID, values.HUMIDITY_SAMPLE_CATEGORY, date, value)
}

func (s *Service) GetLastHumidity(probeID string) (*entities.Sample, error) {
	return s.getLastSample(probeID, values.HUMIDITY_SAMPLE_CATEGORY)
}
