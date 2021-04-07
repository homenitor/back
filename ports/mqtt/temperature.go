package mqtt

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	temperatureTopicTemplate = "%s/temperature"
)

func (s *MQTTServer) SubscribeToRoomTemperature(room string) {
	topic := fmt.Sprintf(temperatureTopicTemplate, room)

	s.subscribe(topic, s.HumidityHandler)
}

func (s *MQTTServer) TemperatureHandler(client mqtt.Client, msg mqtt.Message) {
	room := getRoomFromMessage(msg)

	temperatureValue, err := parseFloatPayload(msg)
	if err != nil {
		return
	}

	s.logging.Debugf("Received temperature sample \"%f\" for room \"%s\"", temperatureValue, room)
	s.service.SaveTemperature(room, time.Now(), temperatureValue)
}
