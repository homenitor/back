package mqtt

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	temperatureTopicTemplate = "%s/temperature"
)

func (s *MQTTServer) SubscribeToProbeTemperature(probeID string) {
	topic := fmt.Sprintf(temperatureTopicTemplate, probeID)

	s.subscribe(topic, s.HumidityHandler)
}

func (s *MQTTServer) TemperatureHandler(client mqtt.Client, msg mqtt.Message) {
	probeID, err := getProbeIDFromMessage(msg)
	if err != nil {
		s.logging.Error(err)
		return
	}

	temperatureValue, err := parseFloatPayload(msg)
	if err != nil {
		s.logging.Error(err)
		return
	}

	s.logging.Debugf("Received temperature sample \"%f\" from probe \"%s\"", temperatureValue, probeID)
	s.service.SaveTemperature(probeID, time.Now(), temperatureValue)
}
