package clients

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/homenitor/back/core/app/libraries"
)

const (
	clientID = "homenitor-back"
)

func NewMQTTClient(
	host string,
	port int,
	logging libraries.Logging,
) mqtt.Client {
	opts := mqtt.NewClientOptions()

	brokerUrl := getBrokerUrl(host, port)
	opts.AddBroker(brokerUrl)

	opts.SetClientID(clientID)
	opts.OnConnect = connectionHandler(logging)
	opts.OnConnectionLost = connectionLostHandler(logging)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return client
}

func connectionHandler(logging libraries.Logging) mqtt.OnConnectHandler {
	return func(client mqtt.Client) {
		logging.Info("Connected to MQTT broker")
	}
}

func connectionLostHandler(logging libraries.Logging) mqtt.ConnectionLostHandler {
	return func(client mqtt.Client, err error) {
		logging.Errorf("Connection lost: %v", err)
	}
}

func getBrokerUrl(host string, port int) string {
	return fmt.Sprintf("tcp://%s:%d", host, port)
}
