package main

import (
	"os"

	"github.com/homenitor/back/adapters"
	"github.com/homenitor/back/app"
	"github.com/homenitor/back/ports/mqtt"
	"github.com/homenitor/back/ports/web"
)

func main() {
	repository := adapters.NewInMemoryRepository()
	logging := adapters.NewLogging()

	service, err := app.NewService(repository, logging)
	if err != nil {
		panic(err)
	}

	mqttHost := os.Getenv("MQTT_HOST")
	if mqttHost == "" {
		mqttHost = "127.0.0.1"
	}

	mqttServer, err := mqtt.NewServer(mqttHost, 1883, service, logging)
	if err != nil {
		panic(err)
	}

	mqttServer.SubscribeToRoomTemperature("livingroom")
	mqttServer.SubscribeToRoomHumidity("livingroom")

	web.Start(service, logging)
}
