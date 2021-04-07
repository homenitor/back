package mqtt

import (
	"strconv"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/homenitor/back/core/app/libraries"
	"github.com/homenitor/back/core/app/samples"
)

type MQTTServer struct {
	client           mqtt.Client
	logging          libraries.Logging
	service          *samples.Service
	qualityOfService int
}

func NewMQTTServer(
	mqttClient mqtt.Client,
	service *samples.Service,
	logging libraries.Logging,
	qualityOfService int,
) *MQTTServer {
	mqttServer := &MQTTServer{
		service:          service,
		logging:          logging,
		qualityOfService: qualityOfService,
	}

	mqttServer.client = mqttClient

	return mqttServer
}

func (s *MQTTServer) subscribe(topic string, handler mqtt.MessageHandler) {
	token := s.client.Subscribe(topic, byte(s.qualityOfService), s.HumidityHandler)
	token.Wait()

	s.logging.Debugf("Subscribed to \"%s\"", topic)
}

func getRoomFromMessage(msg mqtt.Message) string {
	topic := msg.Topic()
	return strings.Split(topic, "/")[0]
}

func parseFloatPayload(msg mqtt.Message) (float64, error) {
	payload := string(msg.Payload())
	return strconv.ParseFloat(string(payload), 64)
}
