package factories

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/homenitor/back/clients"
	"github.com/homenitor/back/config"
)

var (
	mqttClient mqtt.Client
)

func GetMQTTClient() mqtt.Client {
	if mqttClient == nil {
		mqttHost := config.MQTTHost()
		mqttPort := config.MQTTPort()
		mqttClient = clients.NewMQTTClient(mqttHost, mqttPort, GetLoggingLibrary())
	}

	return mqttClient
}
