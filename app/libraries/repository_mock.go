package libraries

import (
	"github.com/homenitor/back/entities"
	"github.com/homenitor/back/values"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (m *RepositoryMock) SaveSample(sample *entities.Sample) error {
	args := m.Called(sample)
	return args.Error(0)
}

func (m *RepositoryMock) GetLastSample(room string, category values.SampleCategory) (*entities.Sample, error) {
	args := m.Called(room, category)

	err := args.Error(1)
	if err != nil {
		return nil, err
	}

	return args.Get(0).(*entities.Sample), nil
}
