package libraries

import (
	"github.com/homenitor/back/core/entities"
	"github.com/homenitor/back/core/values"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (m *RepositoryMock) ListProbes() ([]*entities.ProbeListingView, error) {
	args := m.Called()
	err := args.Error(1)
	if err != nil {
		return nil, err
	}

	return args.Get(0).([]*entities.ProbeListingView), nil
}

func (m *RepositoryMock) GetProbe(probeID int) (*entities.Probe, error) {
	args := m.Called(probeID)
	err := args.Error(1)
	if err != nil {
		return nil, err
	}

	return args.Get(0).(*entities.Probe), nil
}

func (m *RepositoryMock) SaveProbe(probe *entities.Probe) error {
	args := m.Called(probe)
	return args.Error(0)
}

func (m *RepositoryMock) SaveSample(probeID int, sample *entities.Sample) error {
	args := m.Called(probeID, sample)
	return args.Error(0)
}

func (m *RepositoryMock) GetLastSample(probeID int, category values.SampleCategory) (*entities.Sample, error) {
	args := m.Called(probeID, category)

	err := args.Error(1)
	if err != nil {
		return nil, err
	}

	return args.Get(0).(*entities.Sample), nil
}
