package mqtt

import (
	"fmt"
	"testing"

	"github.com/homenitor/back/clients"
	"github.com/homenitor/back/core/values"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	probeID = 1
)

type MessageMock struct {
	mock.Mock
}

func TestParseFloatPayload(t *testing.T) {
	value := 10.0
	payload := []byte("10.0")

	msg := &clients.MessageMock{}
	msg.On("Payload").Return(payload)

	result, err := parseFloatPayload(msg)

	assert.Nil(t, err)
	assert.Equal(t, value, result)
}

func TestParseIntPayload(t *testing.T) {
	value := 10
	payload := []byte("10")

	msg := &clients.MessageMock{}
	msg.On("Payload").Return(payload)

	result, err := parseIntPayload(msg)

	assert.Nil(t, err)
	assert.Equal(t, value, result)
}

func TestGetCategoryFromTopic(t *testing.T) {
	category := values.HUMIDITY_SAMPLE_CATEGORY
	topic := fmt.Sprintf(sampleTopicTemplate, probeID, category)
	result := getCategoryFromTopic(topic)

	assert.Equal(t, category, result)
}

func TestGetProbeIDFromTopicValueNotAnInteger(t *testing.T) {
	category := values.HUMIDITY_SAMPLE_CATEGORY
	topic := fmt.Sprintf("%s/samples/%s", "wrong-id", category)
	result, err := getProbeIDFromTopic(topic)

	assert.NotNil(t, err)
	assert.Equal(t, 0, result)
}

func TestGetProbeIDFromTopicOK(t *testing.T) {
	category := values.HUMIDITY_SAMPLE_CATEGORY
	topic := fmt.Sprintf(sampleTopicTemplate, probeID, category)
	result, err := getProbeIDFromTopic(topic)

	assert.Nil(t, err)
	assert.Equal(t, probeID, result)
}
