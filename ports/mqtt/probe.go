package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	discoverProbesTopic = "discover/probes"
)

func (s *MQTTServer) SubscribeToDiscoverProbes(room string) {
	s.subscribe(discoverProbesTopic, s.HumidityHandler)
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

	s.SubscribeToProbeHumidity(probeID)
	s.SubscribeToProbeTemperature(probeID)
}
