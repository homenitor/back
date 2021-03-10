package app

import (
	"time"

	"github.com/homenitor/back/entities"
)

type Service struct {
	repository Repository
}

func NewService(
	repository Repository,
) (*Service, error) {
	if repository == nil {
		return nil, ErrNilRepository
	}

	return &Service{
		repository: repository,
	}, nil
}

func (s *Service) SaveTemperature(room string, date time.Time, value float64) error {
	t, err := entities.NewTemperature(room, date, value)
	if err != nil {
		return err
	}

	return s.repository.SaveTemperature(t)
}

func (s *Service) GetLastTemperature(room string) (*entities.Temperature, error) {
	t, err := s.repository.GetLastTemperature(room)
	if err != nil {
		return nil, err
	}

	return t, nil
}
