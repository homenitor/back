package libraries

import "github.com/stretchr/testify/mock"

type ProbesLibraryMock struct {
	mock mock.Mock
}

func (m *ProbesLibraryMock) SendDiscoveryMessage() {}
