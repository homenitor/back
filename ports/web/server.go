package web

import (
	"github.com/gin-gonic/gin"
	"github.com/homenitor/back/app"
)

var (
	lastTemperaturePath = "/temperatures/:room"
)

type WebServer struct {
	service *app.Service
}

func NewWebServer(service *app.Service) *WebServer {
	if service == nil {
		panic("service is nil")
	}

	return &WebServer{
		service: service,
	}
}

func (s *WebServer) ConfigureRoutes(r *gin.Engine) {
	r.GET(lastTemperaturePath, s.GetLastTemperature)
}
