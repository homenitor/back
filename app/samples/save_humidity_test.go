package samples

import (
	"testing"

	"github.com/homenitor/back/app"
	"github.com/homenitor/back/app/libraries"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSaveHumidityRepositorySaveHumidityError(t *testing.T) {
	repositoryMock := &libraries.RepositoryMock{}
	repositoryMock.On("SaveHumidity", mock.Anything).Return(app.ErrUnknown)
	loggingMock := &libraries.LoggingMock{}

	service, err := NewService(repositoryMock, loggingMock)
	assert.NoError(t, err)

	err = service.SaveHumidity(room, date, value)
	assert.Equal(t, app.ErrUnknown, err)
}

func TestSaveHumidityOK(t *testing.T) {
	repositoryMock := &libraries.RepositoryMock{}
	repositoryMock.On("SaveHumidity", mock.Anything).Return(nil)
	loggingMock := &libraries.LoggingMock{}

	service, err := NewService(repositoryMock, loggingMock)
	assert.NoError(t, err)

	err = service.SaveHumidity(room, date, value)
	assert.NoError(t, err)
}
