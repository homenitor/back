package factories

import (
	"github.com/homenitor/back/adapters"
	"github.com/homenitor/back/app/libraries"
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

func GetProbesLibrary() libraries.ProbesLibrary {
	if probesLibrary == nil {
		probesLibrary = adapters.NewMQTTProbes(GetMQTTClient(), GetLoggingLibrary())
	}

	return probesLibrary
}

func GetInMemoryRepository() libraries.Repository {
	if inMemoryRepository == nil {
		inMemoryRepository = adapters.NewInMemoryRepository()
	}

	return inMemoryRepository
}
