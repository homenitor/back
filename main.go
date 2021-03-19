package main

import (
	"github.com/homenitor/back/adapters"
	"github.com/homenitor/back/app"
	"github.com/homenitor/back/ports"
	"github.com/homenitor/back/ports/web"
)

func main() {
	repository := adapters.NewInMemoryRepository()
	logging := adapters.NewLogging()

	service, err := app.NewService(repository, logging)
	if err != nil {
		panic(err)
	}

	mqttServer, err := ports.NewMQTTServer("127.0.0.1", 1883, service, logging)
	if err != nil {
		panic(err)
	}

	mqttServer.SubscribeToRoomTemperature("livingroom")
	web.Start(service, logging)
}
