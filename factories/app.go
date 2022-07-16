package factories

import (
	"github.com/homenitor/back/config"
	"github.com/homenitor/back/core/app/libraries"
	"github.com/homenitor/back/core/app/services"
)

var (
	service services.Service
)

func GetService(repository libraries.Repository) services.Service {
	if service == nil {
		var err error

		discoveryPeriod := config.DiscoveryPeriod()

		service, err = services.NewService(
			repository,
			GetLoggingLibrary(),
			GetMQTTProbesLibrary(),
			discoveryPeriod,
		)

		if err != nil {
			panic(err)
		}
	}

	return service
}
