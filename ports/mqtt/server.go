package mqtt

import (
	"fmt"
	"strconv"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/homenitor/back/app"
)

const (
	clientID = "homenitor-back"
)

type Server struct {
	client  mqtt.Client
	logging app.LoggingLibrary
	service *app.Service
}

func NewServer(
	host string,
	port int,
	service *app.Service,
	logging app.LoggingLibrary,
) (*Server, error) {
	brokerUrl := getBrokerUrl(host, port)

	mqttServer := &Server{
		service: service,
		logging: logging,
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerUrl)
	opts.SetClientID(clientID)
	opts.OnConnect = mqttServer.connectionHandler
	opts.OnConnectionLost = mqttServer.connectionLostHandler

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	mqttServer.client = client

	return mqttServer, nil
}

func (s *Server) connectionHandler(client mqtt.Client) {
	s.logging.Info("Connected to MQTT broker")
}

func (s *Server) connectionLostHandler(client mqtt.Client, err error) {
	s.logging.Infof("Connection lost: %v", err)
}

func getBrokerUrl(host string, port int) string {
	return fmt.Sprintf("tcp://%s:%d", host, port)
}

func getRoomFromMessage(msg mqtt.Message) string {
	topic := msg.Topic()
	return strings.Split(topic, "/")[0]
}

func parseFloatPayload(msg mqtt.Message) (float64, error) {
	payload := string(msg.Payload())
	return strconv.ParseFloat(string(payload), 64)
}
