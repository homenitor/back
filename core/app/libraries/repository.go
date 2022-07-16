package libraries

import (
	"time"

	"github.com/homenitor/back/core/entities"
	"github.com/homenitor/back/core/values"
)

type GetSamplesQuery struct {
	From *time.Time
	To   *time.Time
}

type Repository interface {
	GetSamples(category values.SampleCategory, query GetSamplesQuery) ([]*entities.GetSamplesView, error)
	SaveSample(probeID string, sample *entities.Sample) error
	GetLatestSample(probeID string, category values.SampleCategory) (*entities.Sample, error)

	GetProbe(probeID string) (*entities.Probe, error)
	SaveProbe(probe *entities.Probe) error
	ListProbes() ([]*entities.ProbeListingView, error)
	Disconnect() error
}
