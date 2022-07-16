package web

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/homenitor/back/core/app/common"
	"github.com/homenitor/back/core/app/services"
)

var (
	getProbeLatestSamplePath = "/probes/:probeID/getLatestSample/:category"
	listProbesPath           = "/probes"
)

type WebServer struct {
	service services.Service
}

func NewWebServer(service services.Service) *WebServer {
	if service == nil {
		panic("service is nil")
	}

	return &WebServer{
		service: service,
	}
}

func (s *WebServer) ConfigureRoutes(r *gin.Engine) {
	r.GET(getProbeLatestSamplePath, parseProbeIDMiddleware(), s.GetLatestSample)
	r.GET(listProbesPath, s.ListProbes)
}

func (s *WebServer) handleError(c *gin.Context, err error) bool {
	if err != nil {
		if errors.Is(err, common.ErrProbeNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, ErrorResponse{Message: err.Error()})
			return true
		}

		if errors.Is(err, common.ErrNoSampleValueInProbe) {
			c.AbortWithStatusJSON(http.StatusNotFound, ErrorResponse{Message: err.Error()})
			return true
		}
	}

	return false
}
