package adapters

import (
	"sync"

	"github.com/homenitor/back/core/app/libraries"
	"github.com/homenitor/back/core/entities"
	"github.com/homenitor/back/core/values"
)

type InMemoryRepository struct {
	lock *sync.RWMutex

	rooms map[string]*Room
}

type Room struct {
	samples map[values.SampleCategory][]*entities.Sample
}

func NewRoom() *Room {
	return &Room{
		samples: make(map[values.SampleCategory][]*entities.Sample, 0),
	}
}

func NewInMemoryRepository() libraries.Repository {
	return &InMemoryRepository{
		rooms: make(map[string]*Room, 0),
		lock:  &sync.RWMutex{},
	}
}

func (r *InMemoryRepository) SaveSample(sample *entities.Sample) error {
	roomID := sample.Room()
	category := sample.Category()

	room, ok := r.rooms[roomID]
	if ok {
		_, ok := room.samples[category]
		if ok {
			room.samples[category] = append(room.samples[category], sample)
		} else {
			room.samples[category] = []*entities.Sample{sample}
		}
	} else {
		newRoom := NewRoom()
		newRoom.samples[category] = []*entities.Sample{sample}

		r.rooms[roomID] = newRoom
	}

	return nil
}

func (r *InMemoryRepository) GetLastSample(roomID string, category values.SampleCategory) (*entities.Sample, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	room, ok := r.rooms[roomID]

	if !ok {
		return nil, ErrRoomNotFound
	}

	lastSampleIndex := len(room.samples[category]) - 1

	return room.samples[category][lastSampleIndex], nil
}
