package samples

import (
	"testing"

	"github.com/homenitor/back/app/common"
	"github.com/homenitor/back/app/libraries"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSaveTemperatureRepositorySaveTemperatureError(t *testing.T) {
	repositoryMock := &libraries.RepositoryMock{}
	repositoryMock.On("SaveSample", mock.Anything).Return(common.ErrUnknown)
	loggingMock := &libraries.LoggingMock{}

	service, err := NewService(repositoryMock, loggingMock)
	assert.NoError(t, err)

	err = service.SaveTemperature(room, date, value)
	assert.Equal(t, common.ErrUnknown, err)
}

func TestSaveTemperatureOK(t *testing.T) {
	repositoryMock := &libraries.RepositoryMock{}
	repositoryMock.On("SaveSample", mock.Anything).Return(nil)
	loggingMock := &libraries.LoggingMock{}

	service, err := NewService(repositoryMock, loggingMock)
	assert.NoError(t, err)

	err = service.SaveTemperature(room, date, value)
	assert.NoError(t, err)
}
