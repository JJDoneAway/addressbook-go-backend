package middleware

import (
	"net/http"

	"github.com/JJDoneAway/addressbook/docs"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterSwagger(router *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/"

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html#/")
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
