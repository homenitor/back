package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/homenitor/back/core/values"
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

func parseCategoryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		category := c.Param("category")
		if category == "" || category != string(values.TEMPERATURE_SAMPLE_CATEGORY) && category != string(values.HUMIDITY_SAMPLE_CATEGORY) {
			c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{Message: ports.ErrUnknownSampleCategory.Error()})
			return
		}

		c.Set("category", category)
	}
}
