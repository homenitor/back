package clients

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/stretchr/testify/mock"
)

type MessageMock struct {
	mock.Mock
}

func NewMessageMock() mqtt.Message {
	return &MessageMock{}
}

func (m *MessageMock) Ack() {}

func (m *MessageMock) Duplicate() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MessageMock) Qos() byte {
	args := m.Called()
	return byte(args.Int(0))
}

func (m *MessageMock) MessageID() uint16 {
	args := m.Called()
	return uint16(args.Int(0))
}

func (m *MessageMock) Payload() []byte {
	args := m.Called()
	return args.Get(0).([]byte)
}

func (m *MessageMock) Retained() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MessageMock) Topic() string {
	args := m.Called()
	return args.String(0)
}
