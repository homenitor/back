package mqtt

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	temperatureTopicTemplate = "%s/temperature"
)

func (s *Server) SubscribeToRoomTemperature(room string) {
	topic := fmt.Sprintf(temperatureTopicTemplate, room)

	token := s.client.Subscribe(topic, 1, s.TemperatureHandler)
	token.Wait()

	s.logging.Debugf("Subscribed to \"%s\" temperature", room)
}

func (s *Server) TemperatureHandler(client mqtt.Client, msg mqtt.Message) {
	room := getRoomFromMessage(msg)

	temperatureValue, err := parseFloatPayload(msg)
	if err != nil {
		return
	}

	s.logging.Debugf("Received temperature sample \"%f\" for room \"%s\"", temperatureValue, room)

	s.service.SaveTemperature(room, time.Now(), temperatureValue)
}
