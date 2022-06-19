package mqtt

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/homenitor/back/core/values"
)

const (
	sampleTopicTemplate = "%s/samples/%s"
)

func (s *MQTTServer) SubscribeToProbeSample(probeID string, category values.SampleCategory) {
	topic := fmt.Sprintf(sampleTopicTemplate, probeID, category)

	s.subscribe(topic, s.SampleHandler)
}

func (s *MQTTServer) SampleHandler(client mqtt.Client, msg mqtt.Message) {
	topic := msg.Topic()
	probeID, err := getProbeIDFromTopic(topic)
	if err != nil {
		s.logging.Error(err)
		return
	}

	category := getCategoryFromTopic(topic)

	sampleValue, err := parseFloatPayload(msg)
	if err != nil {
		s.logging.Error(err)
		return
	}

	s.logging.Debugf("Received \"%s\" sample \"%f\" for probe \"%s\"", string(category), sampleValue, probeID)
	s.service.SaveSample(probeID, category, time.Now(), sampleValue)
}
