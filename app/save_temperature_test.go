package app

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	uuid  = "uuid"
	room  = "room"
	value = 1.5
)

var (
	date = time.Now()
)

func TestSaveTemperatureRepositorySaveTemperatureError(t *testing.T) {
	repositoryMock := &RepositoryMock{}
	repositoryMock.On("SaveTemperature", mock.Anything).Return(ErrUnknown)

	service, err := NewService(repositoryMock)
	assert.NoError(t, err)

	err = service.SaveTemperature(room, date, value)
	assert.Equal(t, err, ErrUnknown)
}

func TestSaveTemperatureOK(t *testing.T) {
	repositoryMock := &RepositoryMock{}
	repositoryMock.On("SaveTemperature", mock.Anything).Return(nil)

	service, err := NewService(repositoryMock)
	assert.NoError(t, err)

	err = service.SaveTemperature(room, date, value)
	assert.NoError(t, err)
}
