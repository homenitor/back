package web

type GetLastSampleResponse struct {
	Value     float64 `json:"value"`
	Timestamp int64   `json:"timestamp"`
}

/*
  {
    "timestamp": 1589788800.0,
    "probe_id_1": 25.4,
    "probe_id_2": 26.7
  },
  {
    "timestamp": 1589788820.0,
    "probe_id_1": 25.6,
    "probe_id_2": 26.5
  }
*/
type GetSamplesOfCategoryValue map[string]float64

type GetSamplesOfCategoryResponse struct {
	Values []GetSamplesOfCategoryValue `json:"values"`
}

type GetProbesResponse struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}

type ErrorResponse struct {
	Message string `json:"error"`
}
