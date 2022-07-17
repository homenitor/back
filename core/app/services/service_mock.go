package services

import (
	"time"

	"github.com/homenitor/back/core/entities"
	"github.com/homenitor/back/core/values"
	"github.com/stretchr/testify/mock"
)

type ServiceMock struct {
	mock.Mock
}

func NewServiceMock() *ServiceMock {
	return &ServiceMock{}
}

func (m *ServiceMock) GetLatestSample(probeID string, category values.SampleCategory) (*entities.Sample, error) {
	args := m.Called(probeID, category)
	err := args.Error(1)
	if err != nil {
		return nil, err
	}

	return args.Get(0).(*entities.Sample), nil
}

func (m *ServiceMock) GetSamplesByCategory(category values.SampleCategory, sample_range string) ([]*entities.GetSamplesView, error) {
	args := m.Called(category, sample_range)
	err := args.Error(1)
	if err != nil {
		return nil, err
	}

	return args.Get(0).([]*entities.GetSamplesView), nil
}

func (m *ServiceMock) SaveSample(probeID string, category values.SampleCategory, date time.Time, value float64) error {
	args := m.Called(probeID, category, value)
	return args.Error(0)
}

func (m *ServiceMock) ListProbes() ([]*entities.ProbeListingView, error) {
	args := m.Called()
	err := args.Error(1)
	if err != nil {
		return nil, err
	}

	return args.Get(0).([]*entities.ProbeListingView), nil
}

func (m *ServiceMock) DiscoverProbe(probeID string) error {
	args := m.Called(probeID)
	return args.Error(0)
}

func (m *ServiceMock) StartProbesDiscovery() {
	m.Called()
}
