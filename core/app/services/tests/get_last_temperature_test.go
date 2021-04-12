package services

import (
	"testing"
	"time"

	"github.com/homenitor/back/core/app/common"
	"github.com/homenitor/back/core/app/libraries"
	"github.com/homenitor/back/core/app/services"
	"github.com/homenitor/back/core/entities"
	"github.com/homenitor/back/core/values"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetLastTemperatureRepositoryError(t *testing.T) {
	repositoryMock := &libraries.RepositoryMock{}
	repositoryMock.On("GetLastSample", mock.Anything, values.TEMPERATURE_SAMPLE_CATEGORY).Return(nil, common.ErrUnknown)

	loggingMock := &libraries.LoggingMock{}

	probeLibraryMock := &libraries.ProbesLibraryMock{}

	service, err := services.NewService(repositoryMock, loggingMock, probeLibraryMock, time.Second)
	assert.NoError(t, err)

	result, err := service.GetLastTemperature(probeID)

	assert.Nil(t, result)
	assert.Equal(t, common.ErrUnknown, err)
}

func TestGetLastTemperatureOK(t *testing.T) {
	temperature := &entities.Sample{}
	repositoryMock := &libraries.RepositoryMock{}
	repositoryMock.On("GetLastSample", mock.Anything, values.TEMPERATURE_SAMPLE_CATEGORY).Return(temperature, nil)

	loggingMock := &libraries.LoggingMock{}

	probeLibraryMock := &libraries.ProbesLibraryMock{}

	service, err := services.NewService(repositoryMock, loggingMock, probeLibraryMock, time.Second)
	assert.NoError(t, err)

	result, err := service.GetLastTemperature(probeID)

	assert.Equal(t, temperature, result)
	assert.Nil(t, err)
}
