package web

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

func (s *WebServer) GetLastTemperature(c *gin.Context) {
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

	temperature, err := s.service.GetLastTemperature(probeID)
	hasError := s.handleError(c, err)
	if hasError {
		return
	}

	response := &GetLastTemperatureResponse{
		Value: temperature.Value(),
	}

	c.JSON(http.StatusOK, response)
}

func (s *WebServer) GetLastHumidity(c *gin.Context) {
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
	humidity, err := s.service.GetLastHumidity(probeID)
	hasError := s.handleError(c, err)
	if hasError {
		return
	}

	response := &GetLastHumidityResponse{
		Value: humidity.Value(),
	}

	c.JSON(http.StatusOK, response)
}
