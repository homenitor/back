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
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	response := &GetLastTemperatureResponse{
		Value: temperature.Value(),
	}

	c.JSON(http.StatusOK, response)
}
