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
	mongoDBRepository  libraries.Repository
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
			GetLoggingLibrary(),
			qualityOfService,
		)

		if err != nil {
			panic(err)
		}
	}

	return probesLibrary
}

func GetMongoDBRepository() libraries.Repository {
	if mongoDBRepository == nil {
		mongoDBRepository = adapters.NewMongoDBRepository(GetLoggingLibrary())
	}

	return mongoDBRepository
}
