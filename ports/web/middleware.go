package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func parseProbeIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		probeID := c.Param("probeID")
		if probeID == "" {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		c.Set("probeID", probeID)
	}
}
