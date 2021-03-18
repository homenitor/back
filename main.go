package main

import (
	"time"

	"github.com/homenitor/back/adapters"
	"github.com/homenitor/back/app"
	"github.com/homenitor/back/ports"
)

func main() {
	repository := adapters.NewInMemoryRepository()
	logging := adapters.NewLogging()

	service, err := app.NewService(repository, logging)
	if err != nil {
		panic(err)
	}

	err = service.SaveTemperature("test", time.Now(), 11.492)
	if err != nil {
		panic(err)
	}

	_, err = service.GetLastTemperature("test")
	if err != nil {
		panic(err)
	}

	mqttServer, err := ports.NewMQTTServer("127.0.0.1", 1883, service, logging)
	if err != nil {
		panic(err)
	}

	mqttServer.SubscribeToRoomTemperature("livingroom")
	mqttServer.Start()
}
