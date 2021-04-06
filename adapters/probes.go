package adapters

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/homenitor/back/core/app/common"
	"github.com/homenitor/back/core/app/libraries"
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
) (*MQTTProbes, error) {
	if mqttClient == nil {
		return nil, ErrNilMqttClient
	}

	if logging == nil {
		return nil, common.ErrNilLogging
	}

	return &MQTTProbes{
		mqttClient: mqttClient,
		logging:    logging,
	}, nil
}

func (p *MQTTProbes) SendDiscoveryMessage() {
	p.logging.Debug("Send discovery message")

	token := p.mqttClient.Publish(discoveryTopic, 0, true, "")
	token.Wait()
}
