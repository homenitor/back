package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/homenitor/back/ports"
)

func parseProbeIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		probeID := c.Param("probeID")
		if probeID == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{Message: ports.ErrNilProbeID.Error()})
			return
		}

		c.Set("probeID", probeID)
	}
}
