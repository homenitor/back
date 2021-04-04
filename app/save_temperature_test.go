package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSaveTemperatureRepositorySaveTemperatureError(t *testing.T) {
	repositoryMock := &RepositoryMock{}
	repositoryMock.On("SaveTemperature", mock.Anything).Return(ErrUnknown)
	loggingMock := &LoggingLibraryMock{}

	service, err := NewService(repositoryMock, loggingMock)
	assert.NoError(t, err)

	err = service.SaveTemperature(room, date, value)
	assert.Equal(t, ErrUnknown, err)
}

func TestSaveTemperatureOK(t *testing.T) {
	repositoryMock := &RepositoryMock{}
	repositoryMock.On("SaveTemperature", mock.Anything).Return(nil)
	loggingMock := &LoggingLibraryMock{}

	service, err := NewService(repositoryMock, loggingMock)
	assert.NoError(t, err)

	err = service.SaveTemperature(room, date, value)
	assert.NoError(t, err)
}
