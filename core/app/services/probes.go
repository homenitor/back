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

func (s *service) DiscoverProbe(probeID string) error {
	probe, err := s.repository.GetProbe(probeID)
	if probe != nil {
		return nil
	}

	isErrorExpected := err == common.ErrProbeNotFound
	if !isErrorExpected {
		return err
	}

	newProbe := entities.NewProbeWithID(probeID)
	return s.repository.SaveProbe(newProbe)
}

func (s *service) StartProbesDiscovery() {
	go func() {
		for {
			s.probesLibrary.SendDiscoveryMessage()
			time.Sleep(s.discoveryPeriod)
		}
	}()
}
