package factories

import (
	"github.com/homenitor/back/adapters"
	"github.com/homenitor/back/config"
	"github.com/homenitor/back/core/app/libraries"
)

var (
	logging            *adapters.Logging
	probesLibrary      libraries.ProbesLibrary
	inMemoryRepository libraries.Repository
)

func GetLoggingLibrary() libraries.Logging {
	if logging == nil {
		logging = adapters.NewLogging()
	}

	return logging
}

func GetMQTTProbesLibrary() libraries.ProbesLibrary {
	if probesLibrary == nil {
		var err error

		qualityOfService := config.MQTTQualityOfService()

		probesLibrary, err = adapters.NewMQTTProbes(
			GetMQTTClient(),
			GetMQTTServer(),
			GetLoggingLibrary(),
			qualityOfService,
		)

		if err != nil {
			panic(err)
		}
	}

	return probesLibrary
}

func GetInMemoryRepository() libraries.Repository {
	if inMemoryRepository == nil {
		inMemoryRepository = adapters.NewInMemoryRepository()
	}

	return inMemoryRepository
}
