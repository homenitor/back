package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *WebServer) GetLastTemperature(c *gin.Context) {
	room := c.Param("room")
	if room == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	temperature, err := s.service.GetLastTemperature(room)
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
	room := c.Param("room")
	if room == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	humidity, err := s.service.GetLastHumidity(room)
	hasError := s.handleError(c, err)
	if hasError {
		return
	}

	response := &GetLastHumidityResponse{
		Value: humidity.Value(),
	}

	c.JSON(http.StatusOK, response)
}
