package libraries

import "github.com/stretchr/testify/mock"

type LoggingMock struct {
	mock.Mock
}

func (m *LoggingMock) Info(args ...interface{}) {
	m.Called()
}

func (m *LoggingMock) Infof(format string, args ...interface{}) {
	m.Called()
}

func (m *LoggingMock) Error(args ...interface{}) {
	m.Called()
}

func (m *LoggingMock) Errorf(format string, args ...interface{}) {
	m.Called()
}

func (m *LoggingMock) Debug(args ...interface{}) {
	m.Called()
}

func (m *LoggingMock) Debugf(format string, args ...interface{}) {
	m.Called()
}
