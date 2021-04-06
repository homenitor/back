package web

import (
	"github.com/gin-gonic/gin"
	"github.com/homenitor/back/app/samples"
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
