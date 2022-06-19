package adapters

import (
	"sync"

	"github.com/homenitor/back/core/app/common"
	"github.com/homenitor/back/core/app/libraries"
	"github.com/homenitor/back/core/entities"
	"github.com/homenitor/back/core/values"
)

type InMemoryRepository struct {
	lock *sync.RWMutex

	probes map[string]*entities.Probe
}

func NewInMemoryRepository() libraries.Repository {
	return &InMemoryRepository{
		probes: make(map[string]*entities.Probe, 0),
		lock:   &sync.RWMutex{},
	}
}

func (r *InMemoryRepository) ListProbes() ([]*entities.ProbeListingView, error) {
	probeReturns := make([]*entities.ProbeListingView, 0)
	for _, p := range r.probes {
		probeReturn := &entities.ProbeListingView{
			ID:   p.ID(),
			Name: p.Name(),
		}

		probeReturns = append(probeReturns, probeReturn)
	}

	return probeReturns, nil
}

func (r *InMemoryRepository) GetProbe(id string) (*entities.Probe, error) {
	probe, ok := r.probes[id]
	if !ok {
		return nil, common.ErrProbeNotFound
	}

	return probe, nil
}

func (r *InMemoryRepository) SaveProbe(probe *entities.Probe) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.probes[probe.ID()] = probe

	return nil
}

func (r *InMemoryRepository) SaveSample(probeID string, sample *entities.Sample) error {
	probe, isProbeFound := r.probes[probeID]
	if !isProbeFound {
		return common.ErrProbeNotFound
	}

	if sample.Category() == values.HUMIDITY_SAMPLE_CATEGORY {
		probe.RecordHumidity(sample)
	} else {
		probe.RecordTemperature(sample)
	}

	return nil
}

func (r *InMemoryRepository) GetLastSample(probeID string, category values.SampleCategory) (*entities.Sample, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	probe, ok := r.probes[probeID]

	if !ok {
		return nil, ErrProbeNotFound
	}

	switch category {
	case values.HUMIDITY_SAMPLE_CATEGORY:
		return r.getLatestHumidity(probe)
	case values.TEMPERATURE_SAMPLE_CATEGORY:
		return r.getLatestTemperature(probe)
	}

	return nil, ErrUnknownSampleCategory
}

func (r *InMemoryRepository) getLatestTemperature(probe *entities.Probe) (*entities.Sample, error) {
	lastSample, err := probe.LatestTemperature()
	if err != nil {
		return nil, err
	}

	return lastSample, nil
}

func (r *InMemoryRepository) getLatestHumidity(probe *entities.Probe) (*entities.Sample, error) {
	lastSample, err := probe.LatestHumidity()
	if err != nil {
		return nil, err
	}

	return lastSample, nil
}
