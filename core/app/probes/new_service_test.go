package probes

import (
	"testing"
	"time"

	"github.com/homenitor/back/core/app/common"
	"github.com/homenitor/back/core/app/libraries"
	"github.com/stretchr/testify/assert"
)

var duration = time.Duration(200)

func TestNewProbesServiceNilRepository(t *testing.T) {
	duration := time.Duration(200)

	result, err := NewService(nil, nil, nil, duration)

	assert.Nil(t, result)
	assert.Equal(t, common.ErrNilRepository, err)
}

func TestNewProbesServiceNilLogging(t *testing.T) {
	repositoryMock := &libraries.RepositoryMock{}

	result, err := NewService(repositoryMock, nil, nil, duration)

	assert.Nil(t, result)
	assert.Equal(t, common.ErrNilLogging, err)
}

func TestNewProbesServiceNilProbesLibrary(t *testing.T) {
	repositoryMock := &libraries.RepositoryMock{}
	loggingMock := &libraries.LoggingMock{}

	result, err := NewService(repositoryMock, loggingMock, nil, duration)

	assert.Nil(t, result)
	assert.Equal(t, common.ErrNilProbes, err)
}

func TestNewProbesServiceOK(t *testing.T) {
	repositoryMock := &libraries.RepositoryMock{}
	loggingMock := &libraries.LoggingMock{}
	probesLibraryMock := &libraries.ProbesLibraryMock{}

	result, err := NewService(repositoryMock, loggingMock, probesLibraryMock, duration)

	assert.NotNil(t, result)
	assert.Nil(t, err)
}
