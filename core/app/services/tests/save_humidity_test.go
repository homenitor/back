package services

import (
	"testing"
	"time"

	"github.com/homenitor/back/core/app/common"
	"github.com/homenitor/back/core/app/libraries"
	"github.com/homenitor/back/core/app/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSaveHumidityRepositorySaveHumidityError(t *testing.T) {
	repositoryMock := &libraries.RepositoryMock{}
	repositoryMock.On("SaveSample", probeID, mock.Anything).Return(common.ErrUnknown)
	loggingMock := &libraries.LoggingMock{}
	probeLibraryMock := &libraries.ProbesLibraryMock{}

	service, err := services.NewService(repositoryMock, loggingMock, probeLibraryMock, time.Second)
	assert.NoError(t, err)

	err = service.SaveHumidity(probeID, date, value)
	assert.Equal(t, common.ErrUnknown, err)
}

func TestSaveHumidityOK(t *testing.T) {
	repositoryMock := &libraries.RepositoryMock{}
	repositoryMock.On("SaveSample", probeID, mock.Anything).Return(nil)
	loggingMock := &libraries.LoggingMock{}
	probeLibraryMock := &libraries.ProbesLibraryMock{}

	service, err := services.NewService(repositoryMock, loggingMock, probeLibraryMock, time.Second)
	assert.NoError(t, err)

	err = service.SaveHumidity(probeID, date, value)
	assert.NoError(t, err)
}
