package samples

import (
	"testing"

	"github.com/homenitor/back/app/common"
	"github.com/homenitor/back/app/libraries"
	"github.com/stretchr/testify/assert"
)

func TestNewSamplesServiceNilRepository(t *testing.T) {
	result, err := NewService(nil, nil)

	assert.Nil(t, result)
	assert.Equal(t, common.ErrNilRepository, err)
}

func TestNewSamplesServiceNilLogging(t *testing.T) {
	repositoryMock := &libraries.RepositoryMock{}

	result, err := NewService(repositoryMock, nil)

	assert.Nil(t, result)
	assert.Equal(t, common.ErrNilLogging, err)
}

func TestNewSamplesServiceOK(t *testing.T) {
	repositoryMock := &libraries.RepositoryMock{}
	loggingMock := &libraries.LoggingMock{}

	result, err := NewService(repositoryMock, loggingMock)

	assert.NotNil(t, result)
	assert.Nil(t, err)
}
