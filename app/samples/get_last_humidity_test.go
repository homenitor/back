package samples

import (
	"testing"

	"github.com/homenitor/back/app"
	"github.com/homenitor/back/app/libraries"
	"github.com/homenitor/back/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetLastHumidityRepositoryError(t *testing.T) {
	repositoryMock := &libraries.RepositoryMock{}
	repositoryMock.On("GetLastHumidity", mock.Anything).Return(nil, app.ErrUnknown)

	loggingMock := &libraries.LoggingMock{}

	service, err := NewService(repositoryMock, loggingMock)
	assert.NoError(t, err)

	result, err := service.GetLastHumidity(room)

	assert.Nil(t, result)
	assert.Equal(t, app.ErrUnknown, err)
}

func TestGetLastHumidityOK(t *testing.T) {
	temperature := &entities.Humidity{}
	repositoryMock := &libraries.RepositoryMock{}
	repositoryMock.On("GetLastHumidity", mock.Anything).Return(temperature, nil)

	loggingMock := &libraries.LoggingMock{}

	service, err := NewService(repositoryMock, loggingMock)
	assert.NoError(t, err)

	result, err := service.GetLastHumidity(room)

	assert.Equal(t, temperature, result)
	assert.Nil(t, err)
}
