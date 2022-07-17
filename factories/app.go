package factories

import (
	"github.com/homenitor/back/config"
	"github.com/homenitor/back/core/app/services"
)

var (
	service services.Service
)

func GetService() services.Service {
	if service == nil {
		var err error

		discoveryPeriod := config.DiscoveryPeriod()

		service, err = services.NewService(
			GetRepository(),
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
