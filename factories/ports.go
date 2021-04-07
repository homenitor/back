package factories

import (
	"github.com/homenitor/back/config"
	"github.com/homenitor/back/ports/mqtt"
)

var (
	mqttServer *mqtt.MQTTServer
)

func GetMQTTServer() *mqtt.MQTTServer {
	if mqttServer == nil {
		qualityOfService := config.MQTTQualityOfService()
		mqttServer = mqtt.NewMQTTServer(
			GetMQTTClient(),
			GetSamplesService(),
			GetLoggingLibrary(),
			qualityOfService,
		)
	}

	return mqttServer
}
