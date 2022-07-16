package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/homenitor/back/core/values"
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

func (s *WebServer) GetSamplesOfCategory(c *gin.Context) {
	sample_range := c.Param("range")
	if sample_range == "" {
		sample_range = "1h"
	}

	category := c.GetString("category")
	samples, err := s.service.GetSamplesByCategory(values.SampleCategory(category), sample_range)
	hasError := s.handleError(c, err)
	if hasError {
		return
	}

	values := make([]GetSamplesOfCategoryValue, 0)
	for _, sample := range samples {
		value := sample.Values
		value["measured_at"] = float64(sample.MeasuredAt.Unix())
		values = append(values, value)
	}

	response := &GetSamplesOfCategoryResponse{
		Values: values,
	}

	c.JSON(http.StatusOK, response)
}

func (s *WebServer) GetLatestSample(c *gin.Context) {
	probeID := c.GetString("probeID")
	category := c.GetString("category")

	sample, err := s.service.GetLatestSample(probeID, values.SampleCategory(category))
	hasError := s.handleError(c, err)
	if hasError {
		return
	}

	response := &GetLastSampleResponse{
		Value:     sample.Value(),
		Timestamp: sample.MeasuredAt().Unix(),
	}

	c.JSON(http.StatusOK, response)
}
