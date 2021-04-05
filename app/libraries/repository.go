package libraries

import (
	"github.com/homenitor/back/entities"
)

type Repository interface {
	SaveTemperature(temperature *entities.Temperature) error
	GetLastTemperature(room string) (*entities.Temperature, error)

	SaveHumidity(humidity *entities.Humidity) error
	GetLastHumidity(room string) (*entities.Humidity, error)
}
