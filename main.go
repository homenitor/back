package main

import (
	"github.com/homenitor/back/factories"
	"github.com/homenitor/back/ports/web"
)

func main() {
	logging := factories.GetLoggingLibrary()
	samplesService := factories.GetSamplesService()
	mqttServer := factories.GetMQTTServer()
	probesService := factories.GetProbesService()

	probesService.StartProbesDiscovery()

	mqttServer.SubscribeToRoomTemperature("livingroom")
	mqttServer.SubscribeToRoomHumidity("livingroom")

	web.Start(samplesService, logging)
}
