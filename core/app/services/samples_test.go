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

func TestGetLastSampleRepositoryError(t *testing.T) {
	repositoryMock := &libraries.RepositoryMock{}
	repositoryMock.On("GetLatestSample", mock.Anything, values.HUMIDITY_SAMPLE_CATEGORY).Return(nil, common.ErrUnknown)

	loggingMock := &libraries.LoggingMock{}

	probesLibrary := &libraries.ProbesLibraryMock{}

	service, err := NewService(repositoryMock, loggingMock, probesLibrary, time.Second)
	assert.NoError(t, err)

	result, err := service.GetLatestSample(probeID, values.HUMIDITY_SAMPLE_CATEGORY)

	assert.Nil(t, result)
	assert.Equal(t, common.ErrUnknown, err)
}

func TestGetLastSampleOK(t *testing.T) {
	temperature := &entities.Sample{}
	repositoryMock := &libraries.RepositoryMock{}
	repositoryMock.On("GetLatestSample", mock.Anything, values.HUMIDITY_SAMPLE_CATEGORY).Return(temperature, nil)

	loggingMock := &libraries.LoggingMock{}

	probesLibrary := &libraries.ProbesLibraryMock{}

	service, err := NewService(repositoryMock, loggingMock, probesLibrary, time.Second)
	assert.NoError(t, err)

	result, err := service.GetLatestSample(probeID, values.HUMIDITY_SAMPLE_CATEGORY)

	assert.Equal(t, temperature, result)
	assert.Nil(t, err)
}

func TestSaveSampleRepositorySaveSampleError(t *testing.T) {
	repositoryMock := &libraries.RepositoryMock{}
	repositoryMock.On("SaveSample", probeID, mock.Anything).Return(common.ErrUnknown)
	loggingMock := &libraries.LoggingMock{}
	probeLibraryMock := &libraries.ProbesLibraryMock{}

	service, err := NewService(repositoryMock, loggingMock, probeLibraryMock, time.Second)
	assert.NoError(t, err)

	err = service.SaveSample(probeID, values.HUMIDITY_SAMPLE_CATEGORY, date, value)
	assert.Equal(t, common.ErrUnknown, err)
}

func TestSaveSampleOK(t *testing.T) {
	repositoryMock := &libraries.RepositoryMock{}
	repositoryMock.On("SaveSample", probeID, mock.Anything).Return(nil)
	loggingMock := &libraries.LoggingMock{}
	probeLibraryMock := &libraries.ProbesLibraryMock{}

	service, err := NewService(repositoryMock, loggingMock, probeLibraryMock, time.Second)
	assert.NoError(t, err)

	err = service.SaveSample(probeID, values.HUMIDITY_SAMPLE_CATEGORY, date, value)
	assert.NoError(t, err)
}

func TestGetSampleByCategoryInvalidRange(t *testing.T) {
	repositoryMock := &libraries.RepositoryMock{}
	loggingMock := &libraries.LoggingMock{}
	probeLibraryMock := &libraries.ProbesLibraryMock{}

	service, err := NewService(repositoryMock, loggingMock, probeLibraryMock, time.Second)
	assert.NoError(t, err)

	result, err := service.GetSamplesByCategory(values.HUMIDITY_SAMPLE_CATEGORY, "invalid-range")
	assert.Nil(t, result)

	assert.Equal(t, common.ErrInvalidRange, err)
}

func TestGetSamplesByCategoryOK(t *testing.T) {
	category := values.HUMIDITY_SAMPLE_CATEGORY
	repositoryMock := &libraries.RepositoryMock{}
	loggingMock := &libraries.LoggingMock{}
	probeLibraryMock := &libraries.ProbesLibraryMock{}
	samples := make([]*entities.GetSamplesView, 0)
	repositoryMock.On("GetSamples", category, mock.Anything).Return(samples, nil)

	service, err := NewService(repositoryMock, loggingMock, probeLibraryMock, time.Second)
	assert.NoError(t, err)

	result, err := service.GetSamplesByCategory(category, "1m")
	assert.Equal(t, samples, result)
	assert.NoError(t, err)
}
