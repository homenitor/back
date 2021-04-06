package samples

import (
	"testing"

	"github.com/homenitor/back/app/common"
	"github.com/homenitor/back/app/libraries"
	"github.com/homenitor/back/entities"
	"github.com/homenitor/back/values"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetLastHumidityRepositoryError(t *testing.T) {
	repositoryMock := &libraries.RepositoryMock{}
	repositoryMock.On("GetLastSample", mock.Anything, values.HUMIDITY_SAMPLE_CATEGORY).Return(nil, common.ErrUnknown)

	loggingMock := &libraries.LoggingMock{}

	service, err := NewService(repositoryMock, loggingMock)
	assert.NoError(t, err)

	result, err := service.GetLastHumidity(room)

	assert.Nil(t, result)
	assert.Equal(t, common.ErrUnknown, err)
}

func TestGetLastHumidityOK(t *testing.T) {
	temperature := &entities.Sample{}
	repositoryMock := &libraries.RepositoryMock{}
	repositoryMock.On("GetLastSample", mock.Anything, values.HUMIDITY_SAMPLE_CATEGORY).Return(temperature, nil)

	loggingMock := &libraries.LoggingMock{}

	service, err := NewService(repositoryMock, loggingMock)
	assert.NoError(t, err)

	result, err := service.GetLastHumidity(room)

	assert.Equal(t, temperature, result)
	assert.Nil(t, err)
}
