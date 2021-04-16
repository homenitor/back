package mqtt

import (
	"fmt"
	"testing"

	"github.com/homenitor/back/clients"
	"github.com/homenitor/back/core/app/libraries"
	"github.com/homenitor/back/core/app/services"
	"github.com/homenitor/back/core/values"
)

func TestSampleHandlerOK(t *testing.T) {
	value := 10.0
	payload := []byte("10.0")
	category := values.HUMIDITY_SAMPLE_CATEGORY
	topic := fmt.Sprintf(sampleTopicTemplate, probeID, category)
	msgMock := &clients.MessageMock{}
	msgMock.On("Topic").Return(topic)
	msgMock.On("Payload").Return(payload)

	mqttClientMock := clients.NewMQTTClientMock()
	serviceMock := services.NewServiceMock()
	serviceMock.On("SaveSample", probeID, category, value).Return(nil)
	loggingLibraryMock := &libraries.LoggingMock{}
	qualityOfService := 2

	server := NewMQTTServer(mqttClientMock, serviceMock, loggingLibraryMock, qualityOfService)

	server.SampleHandler(mqttClientMock, msgMock)

	serviceMock.AssertCalled(t, "SaveSample", probeID, category, value)
}
