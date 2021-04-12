package services

import (
	"time"

	"github.com/homenitor/back/core/app/common"
	"github.com/homenitor/back/core/entities"
)

func (s *Service) ListProbes() ([]*ListProbesReturn, error) {
	s.logging.Debugf("List all probes")
	probes, err := s.repository.ListProbes()
	if err != nil {
		return nil, err
	}

	probeReturns := make([]*ListProbesReturn, 0)
	for _, p := range probes {
		newProbeListReturn := &ListProbesReturn{
			ID:   p.ID,
			Name: p.Name,
		}
		probeReturns = append(probeReturns, newProbeListReturn)
	}

	return probeReturns, nil
}

func (s *Service) DiscoverProbe(probeID int) error {
	_, err := s.repository.GetProbe(probeID)
	if err != nil {
		if err == common.ErrProbeNotFound {
			return s.createProbe(probeID)
		}
		return err
	}

	return nil
}

func (s *Service) createProbe(prodeID int) error {
	probe := entities.NewProbeWithID(prodeID)

	return s.repository.SaveProbe(probe)
}

func (s *Service) StartProbesDiscovery() {
	go func() {
		for {
			s.sendDiscoveryMessage()
			time.Sleep(s.discoveryPeriod)
		}
	}()
}

func (s *Service) sendDiscoveryMessage() {
	s.probesLibrary.SendDiscoveryMessage()
}
