package main

import (
	"fmt"
	"time"

	"github.com/homenitor/back/adapters"
	"github.com/homenitor/back/app"
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
}
