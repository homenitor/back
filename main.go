package main

import (
	"github.com/homenitor/back/factories"
	"github.com/homenitor/back/ports/web"
)

func main() {
	repository := factories.GetMongoDBRepository()
	defer func() {
		if err := repository.Disconnect(); err != nil {
			panic(err)
		}
	}()

	logging := factories.GetLoggingLibrary()
	service := factories.GetService(repository)
	mqttServer := factories.GetMQTTServer()

	service.StartProbesDiscovery()

	mqttServer.SubscribeToDiscoverProbes()

	web.Start(service, logging)
}
