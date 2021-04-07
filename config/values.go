package config

import (
	"time"
)

type values struct {
	discoveryPeriod  time.Duration
	logLevel         string
	mqttHost         string
	mqttPort         int
	qualityOfService int
}

var instance *values

func init() {
	instance = &values{
		discoveryPeriod: getTimeDuration("DISCOVERY_PERIOD", "10s"),
		mqttHost:        getString("MQTT_HOST", "127.0.0.1"),
		mqttPort:        getInt("MQTT_PORT", 1883),
		logLevel:        getString("LOG_LEVEL", "debug"),
	}
}

func DiscoveryPeriod() time.Duration {
	return instance.discoveryPeriod
}

func LogLevel() string {
	return instance.logLevel
}

func MQTTHost() string {
	return instance.mqttHost
}

func MQTTPort() int {
	return instance.mqttPort
}

func MQTTQualityOfService() int {
	return instance.qualityOfService
}
