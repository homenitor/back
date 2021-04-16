package services

import (
	"testing"
	"time"

	"github.com/homenitor/back/core/app/common"
	"github.com/homenitor/back/core/app/libraries"
	"github.com/homenitor/back/core/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListProbesRepositoryError(t *testing.T) {
	loggingMock := &libraries.LoggingMock{}
	probesLibraryMock := &libraries.ProbesLibraryMock{}
	repositoryMock := &libraries.RepositoryMock{}

	repositoryMock.On("ListProbes").Return(nil, common.ErrUnknown)

	service, err := NewService(repositoryMock, loggingMock, probesLibraryMock, time.Second)
	assert.Nil(t, err)

	probes, err := service.ListProbes()

	assert.Nil(t, probes)
	assert.Equal(t, common.ErrUnknown, err)
}

func TestListProbesOK(t *testing.T) {
	loggingMock := &libraries.LoggingMock{}
	probesLibraryMock := &libraries.ProbesLibraryMock{}
	repositoryMock := &libraries.RepositoryMock{}

	expectedProbes := []*entities.ProbeListingView{
		{
			ID: probeID,
		},
	}

	repositoryMock.On("ListProbes").Return(expectedProbes, nil)

	service, err := NewService(repositoryMock, loggingMock, probesLibraryMock, time.Second)
	assert.Nil(t, err)

	probes, err := service.ListProbes()

	assert.Nil(t, err)
	assert.Equal(t, expectedProbes, probes)
}

func TestDiscoverProbeGetProbeError(t *testing.T) {
	repositoryMock := &libraries.RepositoryMock{}
	loggingMock := &libraries.LoggingMock{}
	probesLibrary := &libraries.ProbesLibraryMock{}
	discoveryPeriod := time.Second

	repositoryMock.On("GetProbe", probeID).Return(nil, common.ErrUnknown)
	repositoryMock.On("SaveProbe", mock.Anything).Return(nil)

	probeService, err := NewService(repositoryMock, loggingMock, probesLibrary, discoveryPeriod)
	assert.Nil(t, err)

	err = probeService.DiscoverProbe(probeID)
	assert.Equal(t, common.ErrUnknown, err)
}

func TestDiscoverProbeProbeAlreadySavedOK(t *testing.T) {
	repositoryMock := &libraries.RepositoryMock{}
	loggingMock := &libraries.LoggingMock{}
	probesLibrary := &libraries.ProbesLibraryMock{}
	discoveryPeriod := time.Second

	probe := entities.NewProbeWithID(probeID)

	repositoryMock.On("GetProbe", probeID).Return(probe, nil)

	probeService, err := NewService(repositoryMock, loggingMock, probesLibrary, discoveryPeriod)
	assert.Nil(t, err)

	err = probeService.DiscoverProbe(probeID)
	assert.Nil(t, err)
}

func TestDiscoverProbeSaveProbeError(t *testing.T) {
	repositoryMock := &libraries.RepositoryMock{}
	loggingMock := &libraries.LoggingMock{}
	probesLibrary := &libraries.ProbesLibraryMock{}
	discoveryPeriod := time.Second

	repositoryMock.On("GetProbe", probeID).Return(nil, common.ErrProbeNotFound)
	repositoryMock.On("SaveProbe", mock.Anything).Return(common.ErrUnknown)

	probeService, err := NewService(repositoryMock, loggingMock, probesLibrary, discoveryPeriod)
	assert.Nil(t, err)

	err = probeService.DiscoverProbe(probeID)
	assert.Equal(t, common.ErrUnknown, err)
}

func TestDiscoverProbeOK(t *testing.T) {
	repositoryMock := &libraries.RepositoryMock{}
	loggingMock := &libraries.LoggingMock{}
	probesLibrary := &libraries.ProbesLibraryMock{}
	discoveryPeriod := time.Second

	probe := entities.NewProbeWithID(probeID)

	repositoryMock.On("GetProbe", probeID).Return(probe, common.ErrProbeNotFound)
	repositoryMock.On("SaveProbe", probe).Return(nil)

	probeService, err := NewService(repositoryMock, loggingMock, probesLibrary, discoveryPeriod)
	assert.Nil(t, err)

	err = probeService.DiscoverProbe(probeID)
	assert.Nil(t, err)
}
