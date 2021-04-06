package libraries

import "github.com/stretchr/testify/mock"

type LoggingMock struct {
	mock mock.Mock
}

func (l *LoggingMock) Info(args ...interface{})                 {}
func (l *LoggingMock) Infof(format string, args ...interface{}) {}

func (l *LoggingMock) Error(args ...interface{})                 {}
func (l *LoggingMock) Errorf(format string, args ...interface{}) {}

func (l *LoggingMock) Debug(args ...interface{})                 {}
func (l *LoggingMock) Debugf(format string, args ...interface{}) {}
