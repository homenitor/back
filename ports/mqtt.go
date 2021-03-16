package ports

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/homenitor/back/app"
)

const (
	clientID = "homenitor-back"

	temperatureTopicTemplate = "%s/temperature"
)

type MQTTServer struct {
	client  mqtt.Client
	service *app.Service
}

func NewMQTTServer(host string, port int, service *app.Service) (*MQTTServer, error) {
	brokerUrl := getBrokerUrl(host, port)

	mqttServer := &MQTTServer{
		service: service,
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

func (s *MQTTServer) Start() {
	forever := make(chan bool)
	<-forever
}

func (s *MQTTServer) SubscribeToRoomTemperature(room string) {
	topic := fmt.Sprintf(temperatureTopicTemplate, room)

	token := s.client.Subscribe(topic, 1, s.TemperatureHandler)
	token.Wait()

	fmt.Printf("Subscribed to topic temperature of room %s\n", room)
}

func (s *MQTTServer) TemperatureHandler(client mqtt.Client, msg mqtt.Message) {
	room := strings.Split(msg.Topic(), "/")[0]

	temperaturePayload := string(msg.Payload())
	temperatureValue, err := strconv.ParseFloat(string(temperaturePayload), 64)

	if err != nil {
		return
	}

	s.service.SaveTemperature(room, time.Now(), temperatureValue)
}

func (s *MQTTServer) connectionHandler(client mqtt.Client) {
	fmt.Println("Connected to MQTT broker")
}

func (s *MQTTServer) connectionLostHandler(client mqtt.Client, err error) {
	fmt.Printf("Connection lost: %v", err)
}

func getBrokerUrl(host string, port int) string {
	return fmt.Sprintf("tcp://%s:%d", host, port)
}
