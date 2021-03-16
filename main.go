package main

import (
	"fmt"
	"time"

	"github.com/homenitor/back/adapters"
	"github.com/homenitor/back/app"
	"github.com/homenitor/back/ports"
)

func main() {
	fmt.Println("homenitor started")

	repository := adapters.NewInMemoryRepository()

	service, err := app.NewService(repository)
	if err != nil {
		panic(err)
	}

	err = service.SaveTemperature("test", time.Now(), 1.6)
	if err != nil {
		panic(err)
	}

	fmt.Println("Temperature inserted in storage")

	t, err := service.GetLastTemperature("test")
	if err != nil {
		panic(err)
	}

	fmt.Println(t)

	mqttServer, err := ports.NewMQTTServer("127.0.0.1", 1883, service)
	if err != nil {
		panic(err)
	}

	mqttServer.SubscribeToRoomTemperature("livingroom")
	mqttServer.Start()
}
