package mqtt

import (
	"fmt"
	"testing"

	"github.com/homenitor/back/clients"
	"github.com/homenitor/back/core/app/common"
	"github.com/homenitor/back/core/app/libraries"
	"github.com/homenitor/back/core/app/services"
	"github.com/homenitor/back/core/values"
	"github.com/stretchr/testify/mock"
)

func TestDiscoverProbesHandlerInvalidPayload(t *testing.T) {
	payload := []byte("invalid")
	category := values.HUMIDITY_SAMPLE_CATEGORY
	topic := fmt.Sprintf(sampleTopicTemplate, probeID, category)
	msgMock := &clients.MessageMock{}
	msgMock.On("Topic").Return(topic)
	msgMock.On("Payload").Return(payload)

	tokenMock := clients.NewTokenMock()
	tokenMock.On("Wait").Return(true)
	mqttClientMock := clients.NewMQTTClientMock()
	serviceMock := services.NewServiceMock()
	serviceMock.On("DiscoverProbe", probeID).Return(common.ErrUnknown)
	loggingLibraryMock := &libraries.LoggingMock{}
	qualityOfService := 2

	server := NewMQTTServer(mqttClientMock, serviceMock, loggingLibraryMock, qualityOfService)

	server.DiscoverProbesHandler(mqttClientMock, msgMock)

	serviceMock.AssertNotCalled(t, "DiscoverProbe")
	mqttClientMock.AssertNotCalled(t, "Subscribe")
}

func TestDiscoverProbesHandlerDiscoverProbesError(t *testing.T) {
	payload := []byte("1")
	category := values.HUMIDITY_SAMPLE_CATEGORY
	topic := fmt.Sprintf(sampleTopicTemplate, probeID, category)
	msgMock := &clients.MessageMock{}
	msgMock.On("Topic").Return(topic)
	msgMock.On("Payload").Return(payload)

	tokenMock := clients.NewTokenMock()
	tokenMock.On("Wait").Return(true)
	mqttClientMock := clients.NewMQTTClientMock()
	serviceMock := services.NewServiceMock()
	serviceMock.On("DiscoverProbe", probeID).Return(common.ErrUnknown)
	loggingLibraryMock := &libraries.LoggingMock{}
	qualityOfService := 2

	server := NewMQTTServer(mqttClientMock, serviceMock, loggingLibraryMock, qualityOfService)

	server.DiscoverProbesHandler(mqttClientMock, msgMock)

	serviceMock.AssertCalled(t, "DiscoverProbe", probeID)
	mqttClientMock.AssertNotCalled(t, "Subscribe")
}

func TestDiscoverProbesHandlerOK(t *testing.T) {
	payload := []byte("1")
	category := values.HUMIDITY_SAMPLE_CATEGORY
	topic := fmt.Sprintf(sampleTopicTemplate, probeID, category)
	msgMock := &clients.MessageMock{}
	msgMock.On("Topic").Return(topic)
	msgMock.On("Payload").Return(payload)

	tokenMock := clients.NewTokenMock()
	tokenMock.On("Wait").Return(true)
	mqttClientMock := clients.NewMQTTClientMock()
	mqttClientMock.On("Subscribe", mock.Anything, mock.Anything, mock.Anything).Return(tokenMock)
	serviceMock := services.NewServiceMock()
	serviceMock.On("DiscoverProbe", probeID).Return(nil)
	loggingLibraryMock := &libraries.LoggingMock{}
	qualityOfService := 2

	server := NewMQTTServer(mqttClientMock, serviceMock, loggingLibraryMock, qualityOfService)

	server.DiscoverProbesHandler(mqttClientMock, msgMock)

	serviceMock.AssertCalled(t, "DiscoverProbe", probeID)
	mqttClientMock.AssertNotCalled(t, "Subscribe", mock.Anything)
}
