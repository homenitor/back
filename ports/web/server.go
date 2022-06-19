package web

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/homenitor/back/adapters"
	"github.com/homenitor/back/core/app/services"
)

var (
	lastTemperaturePath = "/probes/:probeID/getLatestTemperature"
	lastHumidityPath    = "/probes/:probeID/getLatestHumidity"
	listProbesPath      = "/probes"
)

type WebServer struct {
	service *services.Service
}

func NewWebServer(service *services.Service) *WebServer {
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
	r.GET(listProbesPath, s.ListProbes)
}

func (s *WebServer) handleError(c *gin.Context, err error) bool {
	if err != nil {
		if errors.Is(err, adapters.ErrProbeNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return true
		}
	}

	return false
}
