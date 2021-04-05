package factories

import "github.com/homenitor/back/ports/mqtt"

var (
	mqttServer *mqtt.Server
)

func GetMQTTServer() *mqtt.Server {
	if mqttServer == nil {
		newMqttServer, err := mqtt.NewServer(
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
