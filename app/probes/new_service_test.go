package probes

import (
	"testing"

	"github.com/homenitor/back/app/common"
	"github.com/homenitor/back/app/libraries"
	"github.com/stretchr/testify/assert"
)

func TestNewProbesServiceNilRepository(t *testing.T) {
	result, err := NewService(nil, nil, nil)

	assert.Nil(t, result)
	assert.Equal(t, common.ErrNilRepository, err)
}

func TestNewProbesServiceNilLogging(t *testing.T) {
	repositoryMock := &libraries.RepositoryMock{}
	result, err := NewService(repositoryMock, nil, nil)

	assert.Nil(t, result)
	assert.Equal(t, common.ErrNilLogging, err)
}

func TestNewProbesServiceNilProbesLibrary(t *testing.T) {
	repositoryMock := &libraries.RepositoryMock{}
	loggingMock := &libraries.LoggingMock{}
	result, err := NewService(repositoryMock, loggingMock, nil)

	assert.Nil(t, result)
	assert.Equal(t, common.ErrNilProbes, err)
}

func TestNewProbesServiceOK(t *testing.T) {
	repositoryMock := &libraries.RepositoryMock{}
	loggingMock := &libraries.LoggingMock{}
	probesLibraryMock := &libraries.ProbesLibraryMock{}
	result, err := NewService(repositoryMock, loggingMock, probesLibraryMock)

	assert.NotNil(t, result)
	assert.Nil(t, err)
}
