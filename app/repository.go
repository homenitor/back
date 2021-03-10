package app

import (
	"github.com/homenitor/back/entities"
)

type Repository interface {
	SaveTemperature(temperature *entities.Temperature) error
	GetLastTemperature(room string) (*entities.Temperature, error)
}
