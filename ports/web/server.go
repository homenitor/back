package web

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/homenitor/back/adapters"
	"github.com/homenitor/back/core/app/samples"
)

var (
	lastTemperaturePath = "/temperatures/:room"
	lastHumidityPath    = "/humidities/:room"
)

type WebServer struct {
	service *samples.Service
}

func NewWebServer(service *samples.Service) *WebServer {
	if service == nil {
		panic("service is nil")
	}

	return &WebServer{
		service: service,
	}
}

func (s *WebServer) ConfigureRoutes(r *gin.Engine) {
	r.GET(lastTemperaturePath, s.GetLastTemperature)
	r.GET(lastHumidityPath, s.GetLastHumidity)
}

func (s *WebServer) handleError(c *gin.Context, err error) bool {
	if err != nil {
		if errors.Is(err, adapters.ErrRoomNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return true
		}
	}

	return false
}
