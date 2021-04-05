package adapters

import (
	"sync"

	"github.com/homenitor/back/app/libraries"
	"github.com/homenitor/back/entities"
)

type InMemoryRepository struct {
	lock *sync.RWMutex

	humidities   map[string][]*entities.Humidity
	temperatures map[string][]*entities.Temperature
}

func NewInMemoryRepository() libraries.Repository {
	return &InMemoryRepository{
		humidities:   make(map[string][]*entities.Humidity, 0),
		temperatures: make(map[string][]*entities.Temperature, 0),
		lock:         &sync.RWMutex{},
	}
}

func (r *InMemoryRepository) SaveTemperature(t *entities.Temperature) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	room := t.Room()
	_, ok := r.temperatures[room]
	if ok {
		r.temperatures[room] = append(r.temperatures[room], t)
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

func (r *InMemoryRepository) SaveHumidity(h *entities.Humidity) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	room := h.Room()
	_, ok := r.temperatures[room]
	if ok {
		r.humidities[room] = append(r.humidities[room], h)
	} else {
		r.humidities[room] = []*entities.Humidity{h}
	}

	return nil
}

func (r *InMemoryRepository) GetLastHumidity(room string) (*entities.Humidity, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	humiditiesInRoom, ok := r.humidities[room]
	if !ok {
		return nil, ErrRoomNotFound
	}

	lastHumidityIndex := len(humiditiesInRoom) - 1
	return humiditiesInRoom[lastHumidityIndex], nil
}
