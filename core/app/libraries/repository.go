package libraries

import (
	"github.com/homenitor/back/core/entities"
	"github.com/homenitor/back/core/values"
)

type Repository interface {
	SaveSample(probeID string, sample *entities.Sample) error
	GetLatestSample(probeID string, category values.SampleCategory) (*entities.Sample, error)

	GetProbe(probeID string) (*entities.Probe, error)
	SaveProbe(probe *entities.Probe) error
	ListProbes() ([]*entities.ProbeListingView, error)
	Disconnect() error
}
