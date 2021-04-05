package adapters

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/homenitor/back/app/libraries"
)

const (
	discoveryTopic = "discovery"
)

type MQTTProbes struct {
	mqttClient mqtt.Client
	logging    libraries.Logging
}

func NewMQTTProbes(
	mqttClient mqtt.Client,
	logging libraries.Logging,
) *MQTTProbes {
	if mqttClient == nil {
		panic("MQTTClient is nil")
	}

	if logging == nil {
		panic("Logging is nil")
	}

	return &MQTTProbes{
		mqttClient: mqttClient,
		logging:    logging,
	}
}

func (p *MQTTProbes) SendDiscoveryMessage() {
	p.logging.Debug("Send discovery message")

	token := p.mqttClient.Publish(discoveryTopic, 0, true, "")
	token.Wait()
}
