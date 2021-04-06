package probes

import (
	"time"

	"github.com/homenitor/back/core/app/common"
	"github.com/homenitor/back/core/app/libraries"
)

type Service struct {
	repository    libraries.Repository
	logging       libraries.Logging
	probesLibrary libraries.ProbesLibrary
}

func NewService(
	repository libraries.Repository,
	logging libraries.Logging,
	probesLibrary libraries.ProbesLibrary,
) (*Service, error) {
	if repository == nil {
		return nil, common.ErrNilRepository
	}

	if logging == nil {
		return nil, common.ErrNilLogging
	}

	if probesLibrary == nil {
		return nil, common.ErrNilProbes
	}

	return &Service{
		repository:    repository,
		logging:       logging,
		probesLibrary: probesLibrary,
	}, nil
}

func (s *Service) StartProbesDiscovery() {
	go func() {
		for {
			s.sendDiscoveryMessage()
			time.Sleep(10 * time.Second)
		}
	}()
}

func (s *Service) sendDiscoveryMessage() {
	s.probesLibrary.SendDiscoveryMessage()
}
