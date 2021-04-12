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

	mqttServer.SubscribeToDiscoverProbes()

	web.Start(service, logging)
}
