package libraries

import (
	"github.com/homenitor/back/core/entities"
	"github.com/homenitor/back/core/values"
)

type Repository interface {
	SaveSample(sample *entities.Sample) error
	GetLastSample(room string, category values.SampleCategory) (*entities.Sample, error)
}
