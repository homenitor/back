package mqtt

import (
	"strconv"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/homenitor/back/core/app/libraries"
	"github.com/homenitor/back/core/app/services"
	"github.com/homenitor/back/core/values"
)

type MQTTServer struct {
	client           mqtt.Client
	logging          libraries.Logging
	service          services.Service
	qualityOfService int
}

func NewMQTTServer(
	mqttClient mqtt.Client,
	service services.Service,
	logging libraries.Logging,
	qualityOfService int,
) *MQTTServer {
	return &MQTTServer{
		client:           mqttClient,
		service:          service,
		logging:          logging,
		qualityOfService: qualityOfService,
	}
}

func (s *MQTTServer) subscribe(topic string, handler mqtt.MessageHandler) {
	token := s.client.Subscribe(topic, byte(s.qualityOfService), handler)
	token.Wait()

	s.logging.Debugf("Subscribed to \"%s\"", topic)
}

func getProbeIDFromTopic(topic string) (int, error) {
	probeIDIndex := 0
	probeIDString := strings.Split(topic, "/")[probeIDIndex]

	probeID, err := strconv.Atoi(probeIDString)
	if err != nil {
		return 0, err
	}

	return probeID, nil
}

func getCategoryFromTopic(topic string) values.SampleCategory {
	splitTopic := strings.Split(topic, "/")
	categoryIndexInTopic := len(splitTopic) - 1
	categoryString := splitTopic[categoryIndexInTopic]

	return values.SampleCategory(categoryString)
}

func parseFloatPayload(msg mqtt.Message) (float64, error) {
	payload := string(msg.Payload())
	return strconv.ParseFloat(string(payload), 64)
}

func parseIntPayload(msg mqtt.Message) (int, error) {
	payload := string(msg.Payload())
	return strconv.Atoi(payload)
}
