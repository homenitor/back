package samples

import (
	"time"

	"github.com/homenitor/back/core/app/common"
	"github.com/homenitor/back/core/app/libraries"
	"github.com/homenitor/back/core/entities"
	"github.com/homenitor/back/core/values"
)

type Service struct {
	repository libraries.Repository
	logging    libraries.Logging
}

func NewService(
	repository libraries.Repository,
	logging libraries.Logging,
) (*Service, error) {
	if repository == nil {
		return nil, common.ErrNilRepository
	}

	if logging == nil {
		return nil, common.ErrNilLogging
	}

	return &Service{
		repository: repository,
		logging:    logging,
	}, nil
}

func (s *Service) saveSample(room string, category values.SampleCategory, date time.Time, value float64) error {
	sample, err := entities.NewSample(room, category, date, value)
	if err != nil {
		return err
	}

	s.logging.Debugf("Save \"%s\" sample for room \"%s\"", category, room)

	return s.repository.SaveSample(sample)
}

func (s *Service) getLastSample(room string, category values.SampleCategory) (*entities.Sample, error) {
	t, err := s.repository.GetLastSample(room, category)
	if err != nil {
		s.logging.Errorf("Error \"%s\" occured while getting last \"%s\" sample for room \"%s\"", err.Error(), category, room)
		return nil, err
	}

	s.logging.Debugf("Fetched last \"%s\" sample for room \"%s\"", category, room)
	return t, nil
}
