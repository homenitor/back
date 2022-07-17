package config

import (
	"time"
)

type values struct {
	discoveryPeriod      time.Duration
	logLevel             string
	mqttHost             string
	mqttPort             int
	mqttQualityOfService int
	mongodbUri           string
}

var instance *values

func init() {
	instance = &values{
		discoveryPeriod:      getTimeDuration("DISCOVERY_PERIOD", "1m"),
		logLevel:             getString("LOG_LEVEL", "info"),
		mqttHost:             getString("MQTT_HOST", "127.0.0.1"),
		mqttPort:             getInt("MQTT_PORT", 1883),
		mqttQualityOfService: getInt("MQTT_QUALITY_OF_SERVICE", 2),
		mongodbUri:           getString("MONGODB_URI", "mongodb://localhost:27017"),
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
	return instance.mqttQualityOfService
}

func MongoDBURI() string {
	return instance.mongodbUri
}
