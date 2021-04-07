package factories

import (
	"github.com/homenitor/back/config"
	"github.com/homenitor/back/core/app/probes"
	"github.com/homenitor/back/core/app/samples"
)

var (
	samplesService *samples.Service
	probesService  *probes.Service
)

func GetSamplesService() *samples.Service {
	if samplesService == nil {
		var err error

		samplesService, err = samples.NewService(
			GetInMemoryRepository(),
			GetLoggingLibrary(),
		)

		if err != nil {
			panic(err)
		}
	}

	return samplesService
}

func GetProbesService() *probes.Service {
	if probesService == nil {
		var err error

		discoveryPeriod := config.DiscoveryPeriod()

		probesService, err = probes.NewService(
			GetInMemoryRepository(),
			GetLoggingLibrary(),
			GetProbesLibrary(),
			discoveryPeriod,
		)

		if err != nil {
			panic(err)
		}
	}

	return probesService
}
