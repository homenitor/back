package libraries

import "github.com/stretchr/testify/mock"

type ProbesLibraryMock struct {
	mock mock.Mock
}

func (m *ProbesLibraryMock) SendDiscoveryMessage() {}

func (m *ProbesLibraryMock) SubscribeToProbeHumidity(probeID int) {}

func (m *ProbesLibraryMock) SubscribeToProbeTemperature(probeID int) {}
