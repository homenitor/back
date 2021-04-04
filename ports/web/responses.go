package web

type GetLastTemperatureResponse struct {
	Value float64 `json:"value"`
}

type GetLastHumidityResponse struct {
	Value float64 `json:"value"`
}
