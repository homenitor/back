package clients

import (
	"time"

	"github.com/stretchr/testify/mock"
)

type TokenMock struct {
	mock.Mock
}

func NewTokenMock() *TokenMock {
	return &TokenMock{}
}

func (m *TokenMock) Wait() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *TokenMock) WaitTimeout(timeout time.Duration) bool {
	args := m.Called(timeout)
	return args.Bool(0)
}

func (m *TokenMock) Done() <-chan struct{} {
	args := m.Called()
	return args.Get(0).(<-chan struct{})
}

func (m *TokenMock) Error() error {
	args := m.Called()
	return args.Error(0)
}
