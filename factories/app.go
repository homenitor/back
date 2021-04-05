package factories

import (
	"github.com/homenitor/back/app/probes"
	"github.com/homenitor/back/app/samples"
)

var (
	samplesService *samples.Service
	probesService  *probes.Service
)

func GetSamplesService() *samples.Service {
	if samplesService == nil {
		newSampleService, err := samples.NewService(
			GetInMemoryRepository(),
			GetLoggingLibrary(),
		)

		if err != nil {
			panic(err)
		}

		samplesService = newSampleService
	}

	return samplesService
}

func GetProbesService() *probes.Service {
	if probesService == nil {
		newProbesService, err := probes.NewService(
			GetInMemoryRepository(),
			GetLoggingLibrary(),
			GetProbesLibrary(),
		)

		if err != nil {
			panic(err)
		}

		probesService = newProbesService
	}

	return probesService
}
