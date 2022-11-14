package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddStatus(router *gin.Engine) {
	router.GET("status/up", doUp)
	router.GET("status/health", doHealth)
}

// @Summary      Tell Up status
// @Description  Will just tell you if the app is upp and running
// @ID           upStatus
// @Tags         status
// @Produce      json
// @Success      200
// @Router       /status/up [get]
func doUp(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}

// @Summary      Tell the health status
// @Description  Will just tell you if the app is healthy
// @ID           healthStatus
// @Tags         status
// @Produce      json
// @Success      200
// @Router       /status/health [get]
func doHealth(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}
