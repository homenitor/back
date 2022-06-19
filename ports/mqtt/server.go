package mqtt

import (
	"strconv"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/homenitor/back/core/app/libraries"
	"github.com/homenitor/back/core/app/services"
)

type MQTTServer struct {
	client           mqtt.Client
	logging          libraries.Logging
	service          *services.Service
	qualityOfService int
}

func NewMQTTServer(
	mqttClient mqtt.Client,
	service *services.Service,
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
	token := s.client.Subscribe(topic, byte(s.qualityOfService), handler)
	token.Wait()

	s.logging.Debugf("Subscribed to \"%s\"", topic)
}

func getProbeIDFromMessage(msg mqtt.Message) (string, error) {
	topic := msg.Topic()

	return strings.Split(topic, "/")[0], nil
}

func parseFloatPayload(msg mqtt.Message) (float64, error) {
	payload := string(msg.Payload())
	return strconv.ParseFloat(string(payload), 64)
}

func parseIntPayload(msg mqtt.Message) (int, error) {
	payload := string(msg.Payload())
	return strconv.Atoi(payload)
}

func parseStringPayload(msg mqtt.Message) (string, error) {
	return string(msg.Payload()), nil
}
