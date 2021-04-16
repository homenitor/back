package services

import (
	"testing"
	"time"

	"github.com/homenitor/back/core/app/common"
	"github.com/homenitor/back/core/app/libraries"
	"github.com/stretchr/testify/assert"
)

func TestNewSamplesServiceNilRepository(t *testing.T) {
	result, err := NewService(nil, nil, nil, time.Second)

	assert.Nil(t, result)
	assert.Equal(t, common.ErrNilRepository, err)
}

func TestNewSamplesServiceNilLogging(t *testing.T) {
	repositoryMock := &libraries.RepositoryMock{}

	result, err := NewService(repositoryMock, nil, nil, time.Second)

	assert.Nil(t, result)
	assert.Equal(t, common.ErrNilLogging, err)
}

func TestNewSamplesServiceNilProbesLibrary(t *testing.T) {
	repositoryMock := &libraries.RepositoryMock{}
	loggingMock := &libraries.LoggingMock{}

	result, err := NewService(repositoryMock, loggingMock, nil, time.Second)

	assert.Nil(t, result)
	assert.Equal(t, common.ErrNilProbeLibrary, err)
}

func TestNewSamplesServiceOK(t *testing.T) {
	repositoryMock := &libraries.RepositoryMock{}
	loggingMock := &libraries.LoggingMock{}
	probesLibraryMock := &libraries.ProbesLibraryMock{}

	result, err := NewService(repositoryMock, loggingMock, probesLibraryMock, time.Second)

	assert.NotNil(t, result)
	assert.Nil(t, err)
}
