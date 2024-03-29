package services

import (
	"time"

	"github.com/homenitor/back/core/app/common"
	"github.com/homenitor/back/core/app/libraries"
	"github.com/homenitor/back/core/entities"
	"github.com/homenitor/back/core/values"
)

type service struct {
	repository    libraries.Repository
	logging       libraries.Logging
	probesLibrary libraries.ProbesLibrary

	discoveryPeriod time.Duration
}

type Service interface {
	GetSamplesByCategory(category values.SampleCategory, sample_range string) ([]*entities.GetSamplesView, error)
	GetLatestSample(probeID string, category values.SampleCategory) (*entities.Sample, error)
	SaveSample(probeID string, category values.SampleCategory, date time.Time, value float64) error
	ListProbes() ([]*entities.ProbeListingView, error)
	DiscoverProbe(probeID string) error
	StartProbesDiscovery()
}

func NewService(
	repository libraries.Repository,
	logging libraries.Logging,
	probesLibrary libraries.ProbesLibrary,
	discoveryPeriod time.Duration,
) (Service, error) {
	if repository == nil {
		return nil, common.ErrNilRepository
	}

	if logging == nil {
		return nil, common.ErrNilLogging
	}

	if probesLibrary == nil {
		return nil, common.ErrNilProbeLibrary
	}

	return &service{
		repository:      repository,
		logging:         logging,
		probesLibrary:   probesLibrary,
		discoveryPeriod: discoveryPeriod,
	}, nil
}
