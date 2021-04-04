package app

import (
	"github.com/homenitor/back/entities"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (m *RepositoryMock) SaveTemperature(temperature *entities.Temperature) error {
	args := m.Called(temperature)
	return args.Error(0)
}

func (m *RepositoryMock) GetLastTemperature(room string) (*entities.Temperature, error) {
	args := m.Called(room)

	err := args.Error(1)
	if err != nil {
		return nil, err
	}

	return args.Get(0).(*entities.Temperature), nil
}

func (m *RepositoryMock) SaveHumidity(humidity *entities.Humidity) error {
	args := m.Called(humidity)
	return args.Error(0)
}

func (m *RepositoryMock) GetLastHumidity(room string) (*entities.Humidity, error) {
	args := m.Called(room)

	err := args.Error(1)
	if err != nil {
		return nil, err
	}

	return args.Get(0).(*entities.Humidity), nil
}
