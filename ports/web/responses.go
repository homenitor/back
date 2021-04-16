package web

type GetLastTemperatureResponse struct {
	Value float64 `json:"value"`
}

type GetLastHumidityResponse struct {
	Value float64 `json:"value"`
}

type GetProbesResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"`
}
