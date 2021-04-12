package libraries

type ProbesLibrary interface {
	SendDiscoveryMessage()
	SubscribeToProbeHumidity(probeID int)
	SubscribeToProbeTemperature(probeID int)
}
