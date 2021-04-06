package mqtt

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	humidityTopicTemplate = "%s/humidity"
)

func (s *MQTTServer) SubscribeToRoomHumidity(room string) {
	topic := fmt.Sprintf(humidityTopicTemplate, room)

	token := s.client.Subscribe(topic, 1, s.HumidityHandler)
	token.Wait()

	s.logging.Debugf("Subscribed to \"%s\" humidity", room)
}

func (s *MQTTServer) HumidityHandler(client mqtt.Client, msg mqtt.Message) {
	room := getRoomFromMessage(msg)

	humidityValue, err := parseFloatPayload(msg)
	if err != nil {
		return
	}

	s.logging.Debugf("Received humidity sample \"%f\" for room \"%s\"", humidityValue, room)
	s.service.SaveHumidity(room, time.Now(), humidityValue)
}
