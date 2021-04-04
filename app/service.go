package app

import (
	"time"

	"github.com/homenitor/back/entities"
)

type Service struct {
	repository Repository
	logging    LoggingLibrary
}

func NewService(
	repository Repository,
	logging LoggingLibrary,
) (*Service, error) {
	if repository == nil {
		return nil, ErrNilRepository
	}

	if logging == nil {
		return nil, ErrNilLogging
	}

	return &Service{
		repository: repository,
		logging:    logging,
	}, nil
}

func (s *Service) SaveTemperature(room string, date time.Time, value float64) error {
	t, err := entities.NewTemperature(room, date, value)
	if err != nil {
		return err
	}

	s.logging.Debugf("Save temperature sample for room \"%s\"", room)
	return s.repository.SaveTemperature(t)
}

func (s *Service) GetLastTemperature(room string) (*entities.Temperature, error) {
	t, err := s.repository.GetLastTemperature(room)
	if err != nil {
		s.logging.Errorf("Error \"%s\" occured while getting last temperature sample for room \"%s\"", err.Error(), room)
		return nil, err
	}

	s.logging.Debugf("Fetched last temperature sample for room \"%s\"", room)
	return t, nil
}

func (s *Service) SaveHumidity(room string, date time.Time, value float64) error {
	h, err := entities.NewHumidity(room, date, value)
	if err != nil {
		return err
	}

	s.logging.Debugf("Save temperature sample for room \"%s\"", room)
	return s.repository.SaveHumidity(h)
}
