package libraries

import (
	"github.com/homenitor/back/entities"
	"github.com/homenitor/back/values"
)

type Repository interface {
	SaveSample(sample *entities.Sample) error
	GetLastSample(room string, category values.SampleCategory) (*entities.Sample, error)
}
