package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSaveHumidityRepositorySaveHumidityError(t *testing.T) {
	repositoryMock := &RepositoryMock{}
	repositoryMock.On("SaveHumidity", mock.Anything).Return(ErrUnknown)
	loggingMock := &LoggingLibraryMock{}

	service, err := NewService(repositoryMock, loggingMock)
	assert.NoError(t, err)

	err = service.SaveHumidity(room, date, value)
	assert.Equal(t, ErrUnknown, err)
}

func TestSaveHumidityOK(t *testing.T) {
	repositoryMock := &RepositoryMock{}
	repositoryMock.On("SaveHumidity", mock.Anything).Return(nil)
	loggingMock := &LoggingLibraryMock{}

	service, err := NewService(repositoryMock, loggingMock)
	assert.NoError(t, err)

	err = service.SaveHumidity(room, date, value)
	assert.NoError(t, err)
}
