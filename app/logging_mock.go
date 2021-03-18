package app

import "github.com/stretchr/testify/mock"

type LoggingLibraryMock struct {
	mock mock.Mock
}

func (l *LoggingLibraryMock) Info(args ...interface{})                 {}
func (l *LoggingLibraryMock) Infof(format string, args ...interface{}) {}

func (l *LoggingLibraryMock) Error(args ...interface{})                 {}
func (l *LoggingLibraryMock) Errorf(format string, args ...interface{}) {}

func (l *LoggingLibraryMock) Debug(args ...interface{})                 {}
func (l *LoggingLibraryMock) Debugf(format string, args ...interface{}) {}
