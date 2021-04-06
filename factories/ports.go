package factories

import "github.com/homenitor/back/ports/mqtt"

var (
	mqttServer *mqtt.MQTTServer
)

func GetMQTTServer() *mqtt.MQTTServer {
	if mqttServer == nil {
		newMqttServer, err := mqtt.NewMQTTServer(
			GetMQTTClient(),
			GetSamplesService(),
			GetLoggingLibrary(),
		)

		if err != nil {
			panic(err)
		}

		mqttServer = newMqttServer
	}

	return mqttServer
}
