package web

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/homenitor/back/app/libraries"
	"github.com/homenitor/back/app/samples"
)

func Start(service *samples.Service, logging libraries.Logging) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true

	r.Use(cors.New(corsConfig))
	r.Use(gin.Recovery())

	server := NewWebServer(service)
	server.ConfigureRoutes(r)

	logging.Info("Launching awesome server")
	err := r.Run(":3000")
	if err != nil {
		panic(err)
	}
}
