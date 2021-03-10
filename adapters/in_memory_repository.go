package adapters

import (
	"sync"

	"github.com/homenitor/back/entities"
)

type InMemoryRepository struct {
	lock *sync.RWMutex

	temperatures []*entities.Temperature
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		temperatures: make([]*entities.Temperature, 0),
		lock:         &sync.RWMutex{},
	}
}

func (r *InMemoryRepository) SaveTemperature(t *entities.Temperature) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.temperatures = append(r.temperatures, t)

	return nil
}
