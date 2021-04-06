package adapters

import (
	"sync"

	"github.com/homenitor/back/app/libraries"
	"github.com/homenitor/back/entities"
	"github.com/homenitor/back/values"
)

type InMemoryRepository struct {
	lock *sync.RWMutex

	humidities   map[string][]*entities.Sample
	temperatures map[string][]*entities.Sample
}

func NewInMemoryRepository() libraries.Repository {
	return &InMemoryRepository{
		humidities:   make(map[string][]*entities.Sample, 0),
		temperatures: make(map[string][]*entities.Sample, 0),
		lock:         &sync.RWMutex{},
	}
}

func (r *InMemoryRepository) SaveSample(sample *entities.Sample) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	room := sample.Room()

	var samples []*entities.Sample
	var ok bool
	if sample.Category() == values.TEMPERATURE_SAMPLE_CATEGORY {
		samples, ok = r.temperatures[room]
	} else {
		samples, ok = r.humidities[room]
	}

	if ok {
		samples = append(samples, sample)
	} else {
		samples = []*entities.Sample{sample}
	}

	return nil
}

func (r *InMemoryRepository) GetLastSample(room string, category values.SampleCategory) (*entities.Sample, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	var samples []*entities.Sample
	var ok bool
	if category == values.TEMPERATURE_SAMPLE_CATEGORY {
		samples = r.temperatures[room]
	} else {
		samples = r.humidities[room]
	}

	if !ok {
		return nil, ErrRoomNotFound
	}

	lastTemperatureIndex := len(samples) - 1
	return samples[lastTemperatureIndex], nil
}
