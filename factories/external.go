package factories

import (
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/homenitor/back/external"
)

var (
	mqttClient mqtt.Client
)

func GetMQTTClient() mqtt.Client {
	if mqttClient == nil {
		mqttHost := os.Getenv("MQTT_HOST")
		if mqttHost == "" {
			mqttHost = "127.0.0.1"
		}

		mqttClient = external.NewMQTTClient(mqttHost, 1883, GetLoggingLibrary())
	}

	return mqttClient
}
