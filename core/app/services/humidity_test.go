package services

import (
	"testing"
	"time"

	"github.com/homenitor/back/core/app/common"
	"github.com/homenitor/back/core/app/libraries"
	"github.com/homenitor/back/core/entities"
	"github.com/homenitor/back/core/values"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetLastHumidityRepositoryError(t *testing.T) {
	repositoryMock := &libraries.RepositoryMock{}
	repositoryMock.On("GetLastSample", mock.Anything, values.HUMIDITY_SAMPLE_CATEGORY).Return(nil, common.ErrUnknown)

	loggingMock := &libraries.LoggingMock{}

	probesLibrary := &libraries.ProbesLibraryMock{}

	service, err := NewService(repositoryMock, loggingMock, probesLibrary, time.Second)
	assert.NoError(t, err)

	result, err := service.GetLastHumidity(probeID)

	assert.Nil(t, result)
	assert.Equal(t, common.ErrUnknown, err)
}

func TestGetLastHumidityOK(t *testing.T) {
	temperature := &entities.Sample{}
	repositoryMock := &libraries.RepositoryMock{}
	repositoryMock.On("GetLastSample", mock.Anything, values.HUMIDITY_SAMPLE_CATEGORY).Return(temperature, nil)

	loggingMock := &libraries.LoggingMock{}

	probesLibrary := &libraries.ProbesLibraryMock{}

	service, err := NewService(repositoryMock, loggingMock, probesLibrary, time.Second)
	assert.NoError(t, err)

	result, err := service.GetLastHumidity(probeID)

	assert.Equal(t, temperature, result)
	assert.Nil(t, err)
}

func TestSaveHumidityRepositorySaveHumidityError(t *testing.T) {
	repositoryMock := &libraries.RepositoryMock{}
	repositoryMock.On("SaveSample", probeID, mock.Anything).Return(common.ErrUnknown)
	loggingMock := &libraries.LoggingMock{}
	probeLibraryMock := &libraries.ProbesLibraryMock{}

	service, err := NewService(repositoryMock, loggingMock, probeLibraryMock, time.Second)
	assert.NoError(t, err)

	err = service.SaveHumidity(probeID, date, value)
	assert.Equal(t, common.ErrUnknown, err)
}

func TestSaveHumidityOK(t *testing.T) {
	repositoryMock := &libraries.RepositoryMock{}
	repositoryMock.On("SaveSample", probeID, mock.Anything).Return(nil)
	loggingMock := &libraries.LoggingMock{}
	probeLibraryMock := &libraries.ProbesLibraryMock{}

	service, err := NewService(repositoryMock, loggingMock, probeLibraryMock, time.Second)
	assert.NoError(t, err)

	err = service.SaveHumidity(probeID, date, value)
	assert.NoError(t, err)
}
