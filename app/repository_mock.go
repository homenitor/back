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
