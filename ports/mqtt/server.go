package mqtt

import (
	"strconv"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/homenitor/back/app/libraries"
	"github.com/homenitor/back/app/samples"
)

type Server struct {
	client  mqtt.Client
	logging libraries.Logging
	service *samples.Service
}

func NewServer(
	mqttClient mqtt.Client,
	service *samples.Service,
	logging libraries.Logging,
) (*Server, error) {

	mqttServer := &Server{
		service: service,
		logging: logging,
	}

	mqttServer.client = mqttClient

	return mqttServer, nil
}

func getRoomFromMessage(msg mqtt.Message) string {
	topic := msg.Topic()
	return strings.Split(topic, "/")[0]
}

func parseFloatPayload(msg mqtt.Message) (float64, error) {
	payload := string(msg.Payload())
	return strconv.ParseFloat(string(payload), 64)
}
