package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/homenitor/back/core/values"
)

const (
	discoverProbesTopic = "discover/probes"
)

func (s *MQTTServer) SubscribeToDiscoverProbes() {
	s.subscribe(discoverProbesTopic, s.DiscoverProbesHandler)
}

func (s *MQTTServer) DiscoverProbesHandler(client mqtt.Client, msg mqtt.Message) {
	probeID, err := parseIntPayload(msg)
	if err != nil {
		s.logging.Error(err)
		return
	}

	err = s.service.DiscoverProbe(probeID)
	if err != nil {
		s.logging.Error(err)
		return
	}

	s.SubscribeToProbeSample(probeID, values.HUMIDITY_SAMPLE_CATEGORY)
	s.SubscribeToProbeSample(probeID, values.TEMPERATURE_SAMPLE_CATEGORY)
}
