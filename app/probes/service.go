package probes

import (
	"time"

	"github.com/homenitor/back/app"
	"github.com/homenitor/back/app/libraries"
)

type Service struct {
	repository    libraries.Repository
	logging       libraries.Logging
	probesLibrary libraries.Probes
}

func NewService(
	repository libraries.Repository,
	logging libraries.Logging,
	probesLibrary libraries.Probes,
) (*Service, error) {
	if repository == nil {
		return nil, app.ErrNilRepository
	}

	if logging == nil {
		return nil, app.ErrNilLogging
	}

	if probesLibrary == nil {
		return nil, app.ErrNilProbes
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
