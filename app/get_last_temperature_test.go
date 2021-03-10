package app

import (
	"testing"

	"github.com/homenitor/back/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetLastTemperatureRepositoryError(t *testing.T) {
	repositoryMock := &RepositoryMock{}
	repositoryMock.On("GetLastTemperature", mock.Anything).Return(nil, ErrUnknown)

	service, err := NewService(repositoryMock)
	assert.NoError(t, err)

	result, err := service.GetLastTemperature(room)

	assert.Nil(t, result)
	assert.Equal(t, ErrUnknown, err)
}

func TestGetLastTemperatureOK(t *testing.T) {
	temperature := &entities.Temperature{}
	repositoryMock := &RepositoryMock{}
	repositoryMock.On("GetLastTemperature", mock.Anything).Return(temperature, nil)

	service, err := NewService(repositoryMock)
	assert.NoError(t, err)

	result, err := service.GetLastTemperature(room)

	assert.Equal(t, temperature, result)
	assert.Nil(t, err)
}
