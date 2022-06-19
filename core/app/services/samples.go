package services

import (
	"time"

	"github.com/homenitor/back/core/entities"
	"github.com/homenitor/back/core/values"
)

func (s *service) SaveSample(probeID string, category values.SampleCategory, date time.Time, value float64) error {
	sample, err := entities.NewSample(category, date, value)
	if err != nil {
		return err
	}

	s.logging.Debugf("Save \"%s\" sample of probe \"%s\"", category, probeID)
	return s.repository.SaveSample(probeID, sample)
}

func (s *service) GetLastSample(probeID string, category values.SampleCategory) (*entities.Sample, error) {
	t, err := s.repository.GetLastSample(probeID, category)
	if err != nil {
		s.logging.Errorf("Error \"%s\" occured while getting last \"%s\" sample of probe \"%s\"", err.Error(), category, probeID)
		return nil, err
	}

	s.logging.Debugf("Fetched last \"%s\" sample of probe \"%s\"", category, probeID)
	return t, nil
}
