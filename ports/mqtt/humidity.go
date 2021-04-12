package mqtt

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	humidityTopicTemplate = "%d/humidity"
)

func (s *MQTTServer) SubscribeToProbeHumidity(probeID int) {
	topic := fmt.Sprintf(humidityTopicTemplate, probeID)

	s.subscribe(topic, s.HumidityHandler)
}

func (s *MQTTServer) HumidityHandler(client mqtt.Client, msg mqtt.Message) {
	probeID, err := getProbeIDFromMessage(msg)
	if err != nil {
		s.logging.Error(err)
		return
	}

	humidityValue, err := parseFloatPayload(msg)
	if err != nil {
		s.logging.Error(err)
		return
	}

	s.logging.Debugf("Received humidity sample \"%f\" for probe \"%d\"", humidityValue, probeID)
	s.service.SaveHumidity(probeID, time.Now(), humidityValue)
}
