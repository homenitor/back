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

func (m *RepositoryMock) GetProbe(probeID string) (*entities.Probe, error) {
	args := m.Called(probeID)
	err := args.Error(1)
	if err != nil {
		return nil, err
	}

	return args.Get(0).(*entities.Probe), nil
}

func (m *RepositoryMock) GetSamples(category values.SampleCategory, query GetSamplesQuery) ([]*entities.GetSamplesView, error) {
	args := m.Called(category, query)
	err := args.Error(1)
	if err != nil {
		return nil, err
	}

	return args.Get(0).([]*entities.GetSamplesView), nil
}

func (m *RepositoryMock) SaveProbe(probe *entities.Probe) error {
	args := m.Called(probe)
	return args.Error(0)
}

func (m *RepositoryMock) SaveSample(probeID string, sample *entities.Sample) error {
	args := m.Called(probeID, sample)
	return args.Error(0)
}

func (m *RepositoryMock) GetLatestSample(probeID string, category values.SampleCategory) (*entities.Sample, error) {
	args := m.Called(probeID, category)

	err := args.Error(1)
	if err != nil {
		return nil, err
	}

	return args.Get(0).(*entities.Sample), nil
}

func (m *RepositoryMock) Disconnect() error {
	args := m.Called()
	return args.Error(0)
}
