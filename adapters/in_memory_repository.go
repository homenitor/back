package adapters

import (
	"sync"

	"github.com/homenitor/back/entities"
)

type InMemoryRepository struct {
	lock *sync.RWMutex

	temperatures map[string][]*entities.Temperature
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		temperatures: make(map[string][]*entities.Temperature, 0),
		lock:         &sync.RWMutex{},
	}
}

func (r *InMemoryRepository) SaveTemperature(t *entities.Temperature) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	room := t.Room()
	temperaturesInRoom, ok := r.temperatures[room]
	if ok {
		temperaturesInRoom = append(temperaturesInRoom, t)
	} else {
		r.temperatures[room] = []*entities.Temperature{t}
	}

	return nil
}

func (r *InMemoryRepository) GetLastTemperature(room string) (*entities.Temperature, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	temperaturesInRoom, ok := r.temperatures[room]
	if !ok {
		return nil, ErrRoomNotFound
	}

	lastTemperatureIndex := len(temperaturesInRoom) - 1
	return temperaturesInRoom[lastTemperatureIndex], nil
}
