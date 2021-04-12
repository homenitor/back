package libraries

import (
	"github.com/homenitor/back/core/entities"
	"github.com/homenitor/back/core/values"
)

type ListProbesRepositoryReturn struct {
	ID   int
	Name string
}

type Repository interface {
	SaveSample(probeID int, sample *entities.Sample) error
	GetLastSample(probeID int, category values.SampleCategory) (*entities.Sample, error)

	GetProbe(probeID int) (*entities.Probe, error)
	SaveProbe(probe *entities.Probe) error
	ListProbes() ([]*ListProbesRepositoryReturn, error)
}
