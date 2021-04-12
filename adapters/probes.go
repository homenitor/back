package adapters

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/homenitor/back/core/app/common"
	"github.com/homenitor/back/core/app/libraries"
	mqttPorts "github.com/homenitor/back/ports/mqtt"
)

const (
	discoveryTopic = "discovery"
)

type MQTTProbes struct {
	mqttClient       mqtt.Client
	mqttServer       *mqttPorts.MQTTServer
	logging          libraries.Logging
	qualityOfService int
}

func NewMQTTProbes(
	mqttClient mqtt.Client,
	mqttServer *mqttPorts.MQTTServer,
	logging libraries.Logging,
	qualityOfService int,
) (*MQTTProbes, error) {
	if mqttClient == nil {
		return nil, ErrNilMqttClient
	}

	if mqttServer == nil {
		return nil, ErrNilMqttServer
	}

	if logging == nil {
		return nil, common.ErrNilLogging
	}

	return &MQTTProbes{
		mqttClient:       mqttClient,
		mqttServer:       mqttServer,
		logging:          logging,
		qualityOfService: qualityOfService,
	}, nil
}

func (p *MQTTProbes) SendDiscoveryMessage() {
	p.logging.Debug("Send discovery message")

	token := p.mqttClient.Publish(discoveryTopic, byte(p.qualityOfService), true, "")
	token.Wait()
}

func (p *MQTTProbes) SubscribeToProbeHumidity(probeID int) {
	p.mqttServer.SubscribeToRoomHumidity(probeID)
}

func (p *MQTTProbes) SubscribeToProbeTemperature(probeID int) {
	p.mqttServer.SubscribeToRoomTemperature(probeID)
}
