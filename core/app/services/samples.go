package services

import (
	"time"

	"github.com/homenitor/back/core/app/common"
	"github.com/homenitor/back/core/app/libraries"
	"github.com/homenitor/back/core/entities"
	"github.com/homenitor/back/core/values"
)

func (s *service) SaveSample(probeID string, category values.SampleCategory, date time.Time, value float64) error {
	sample := entities.NewSample(category, date, value)

	s.logging.Debugf("save sample: probe=\"%s\", category=\"%s\"", probeID, category)
	return s.repository.SaveSample(probeID, sample)
}

func (s *service) GetSamplesByCategory(category values.SampleCategory, sample_range string) ([]*entities.GetSamplesView, error) {
	duration_range, err := time.ParseDuration(sample_range)
	if err != nil {
		return nil, common.ErrInvalidRange
	}

	from := time.Now().Add(-duration_range)
	to := time.Now()
	query := libraries.GetSamplesQuery{From: &from, To: &to}

	return s.repository.GetSamples(category, query)
}

func (s *service) GetLatestSample(probeID string, category values.SampleCategory) (*entities.Sample, error) {
	t, err := s.repository.GetLatestSample(probeID, category)
	if err != nil {
		s.logging.Errorf("Error \"%s\" occured while getting last \"%s\" sample of probe \"%s\"", err.Error(), category, probeID)
		return nil, err
	}

	s.logging.Debugf("Fetched last \"%s\" sample of probe \"%s\"", category, probeID)
	return t, nil
}
