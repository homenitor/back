package services

import (
	"time"

	"github.com/homenitor/back/core/app/common"
	"github.com/homenitor/back/core/entities"
)

func (s *service) ListProbes() ([]*entities.ProbeListingView, error) {
	s.logging.Debugf("List all probes")
	return s.repository.ListProbes()
}

func (s *service) DiscoverProbe(probeID int) error {
	_, err := s.repository.GetProbe(probeID)
	isProbeFound := err == nil
	if isProbeFound {
		return nil
	}

	isErrorExpected := err == common.ErrProbeNotFound
	if !isErrorExpected {
		return err
	}

	probe := entities.NewProbeWithID(probeID)
	return s.repository.SaveProbe(probe)
}

func (s *service) StartProbesDiscovery() {
	go func() {
		for {
			s.probesLibrary.SendDiscoveryMessage()
			time.Sleep(s.discoveryPeriod)
		}
	}()
}
