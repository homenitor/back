package web

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/homenitor/back/adapters"
)

func (s *WebServer) GetLastTemperature(c *gin.Context) {
	room := c.Param("room")
	if room == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	temperature, err := s.service.GetLastTemperature(room)
	if err != nil {
		if errors.Is(err, adapters.ErrRoomNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatus(http.StatusInternalServerError)
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
	if err != nil {
		if errors.Is(err, adapters.ErrRoomNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	response := &GetLastHumidityResponse{
		Value: humidity.Value(),
	}

	c.JSON(http.StatusOK, response)
}
