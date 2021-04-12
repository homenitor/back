package services

import (
	"time"

	"github.com/homenitor/back/core/app/common"
	"github.com/homenitor/back/core/app/libraries"
)

type Service struct {
	repository    libraries.Repository
	logging       libraries.Logging
	probesLibrary libraries.ProbesLibrary

	discoveryPeriod time.Duration
}

func NewService(
	repository libraries.Repository,
	logging libraries.Logging,
	probesLibrary libraries.ProbesLibrary,
	discoveryPeriod time.Duration,
) (*Service, error) {
	if repository == nil {
		return nil, common.ErrNilRepository
	}

	if logging == nil {
		return nil, common.ErrNilLogging
	}

	if probesLibrary == nil {
		return nil, common.ErrNilProbeLibrary
	}

	return &Service{
		repository:      repository,
		logging:         logging,
		probesLibrary:   probesLibrary,
		discoveryPeriod: discoveryPeriod,
	}, nil
}
