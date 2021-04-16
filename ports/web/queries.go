package web

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/homenitor/back/core/values"
)

func (s *WebServer) ListProbes(c *gin.Context) {
	probes, err := s.service.ListProbes()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	probesResponse := make([]*GetProbesResponse, 0)
	for _, probe := range probes {
		response := &GetProbesResponse{
			ID:   probe.ID,
			Name: probe.Name,
		}
		probesResponse = append(probesResponse, response)
	}

	c.JSON(http.StatusOK, probesResponse)
}

func (s *WebServer) GetLastSample(c *gin.Context) {
	probeIDString := c.Param("probeID")
	if probeIDString == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	probeID, err := strconv.Atoi(probeIDString)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	category := c.Param("category")
	if category == "" || category != string(values.TEMPERATURE_SAMPLE_CATEGORY) && category != string(values.HUMIDITY_SAMPLE_CATEGORY) {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	sample, err := s.service.GetLastSample(probeID, values.SampleCategory(category))
	hasError := s.handleError(c, err)
	if hasError {
		return
	}

	response := &GetLastSampleResponse{
		Value: sample.Value(),
	}

	c.JSON(http.StatusOK, response)
}
