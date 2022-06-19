package web

type GetLastSampleResponse struct {
	Value float64 `json:"value"`
}

type GetProbesResponse struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}
