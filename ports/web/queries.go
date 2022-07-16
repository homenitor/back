package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/homenitor/back/core/values"
	"github.com/homenitor/back/ports"
)

func (s *WebServer) ListProbes(c *gin.Context) {
	probes, err := s.service.ListProbes()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
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

func (s *WebServer) GetLatestSample(c *gin.Context) {
	probeID := c.GetString("probeID")

	category := c.Param("category")
	if category == "" || category != string(values.TEMPERATURE_SAMPLE_CATEGORY) && category != string(values.HUMIDITY_SAMPLE_CATEGORY) {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{Message: ports.ErrUnknownSampleCategory.Error()})
		return
	}

	sample, err := s.service.GetLatestSample(probeID, values.SampleCategory(category))
	hasError := s.handleError(c, err)
	if hasError {
		return
	}

	response := &GetLastSampleResponse{
		Value:     sample.Value(),
		Timestamp: sample.Timestamp().Unix(),
	}

	c.JSON(http.StatusOK, response)
}
