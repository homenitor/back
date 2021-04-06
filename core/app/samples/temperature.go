package samples

import (
	"time"

	"github.com/homenitor/back/core/entities"
	"github.com/homenitor/back/core/values"
)

func (s *Service) SaveTemperature(room string, date time.Time, value float64) error {
	return s.saveSample(room, values.TEMPERATURE_SAMPLE_CATEGORY, date, value)
}

func (s *Service) GetLastTemperature(room string) (*entities.Sample, error) {
	return s.getLastSample(room, values.TEMPERATURE_SAMPLE_CATEGORY)
}
