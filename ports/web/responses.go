package web

type GetLastSampleResponse struct {
	Value     float64 `json:"value"`
	Timestamp int64   `json:"timestamp"`
}

type GetProbesResponse struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}

type ErrorResponse struct {
	Message string `json:"error"`
}
