package app

import (
	"testing"

	"github.com/homenitor/back/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetLastHumidityRepositoryError(t *testing.T) {
	repositoryMock := &RepositoryMock{}
	repositoryMock.On("GetLastHumidity", mock.Anything).Return(nil, ErrUnknown)

	loggingMock := &LoggingLibraryMock{}

	service, err := NewService(repositoryMock, loggingMock)
	assert.NoError(t, err)

	result, err := service.GetLastHumidity(room)

	assert.Nil(t, result)
	assert.Equal(t, ErrUnknown, err)
}

func TestGetLastHumidityOK(t *testing.T) {
	temperature := &entities.Humidity{}
	repositoryMock := &RepositoryMock{}
	repositoryMock.On("GetLastHumidity", mock.Anything).Return(temperature, nil)

	loggingMock := &LoggingLibraryMock{}

	service, err := NewService(repositoryMock, loggingMock)
	assert.NoError(t, err)

	result, err := service.GetLastHumidity(room)

	assert.Equal(t, temperature, result)
	assert.Nil(t, err)
}
