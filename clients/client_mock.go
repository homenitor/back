package clients

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/stretchr/testify/mock"
)

type MQTTClientMock struct {
	mock.Mock
}

func NewMQTTClientMock() *MQTTClientMock {
	return &MQTTClientMock{}
}

func (m *MQTTClientMock) Disconnect(quiesce uint) {
	m.Called(quiesce)
}

func (m *MQTTClientMock) Publish(topic string, qos byte, retained bool, payload interface{}) mqtt.Token {
	args := m.Called(topic, qos, retained)
	return args.Get(0).(mqtt.Token)
}

func (m *MQTTClientMock) Subscribe(topic string, qos byte, callback mqtt.MessageHandler) mqtt.Token {
	args := m.Called(topic, qos, callback)
	return args.Get(0).(mqtt.Token)
}

func (m *MQTTClientMock) SubscribeMultiple(filters map[string]byte, callback mqtt.MessageHandler) mqtt.Token {
	args := m.Called(filters, callback)
	return args.Get(0).(mqtt.Token)
}

func (m *MQTTClientMock) Unsubscribe(topics ...string) mqtt.Token {
	args := m.Called(topics)
	return args.Get(0).(mqtt.Token)
}

func (m *MQTTClientMock) OptionsReader() mqtt.ClientOptionsReader {
	args := m.Called()
	return args.Get(0).(mqtt.ClientOptionsReader)
}

func (m *MQTTClientMock) Connect() mqtt.Token {
	args := m.Called()
	return args.Get(0).(mqtt.Token)
}

func (m *MQTTClientMock) IsConnected() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MQTTClientMock) IsConnectionOpen() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MQTTClientMock) AddRoute(topic string, callback mqtt.MessageHandler) {
	m.Called(topic, callback)
}
