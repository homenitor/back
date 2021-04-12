package main

import (
	"github.com/homenitor/back/factories"
	"github.com/homenitor/back/ports/web"
)

func main() {
	logging := factories.GetLoggingLibrary()
	service := factories.GetService()
	mqttServer := factories.GetMQTTServer()

	service.StartProbesDiscovery()

	mqttServer.SubscribeToProbeTemperature(1)
	mqttServer.SubscribeToProbeHumidity(1)

	web.Start(service, logging)
}
