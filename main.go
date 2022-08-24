package main

import (
	"log"
	"net/http"

	"github.com/JJDoneAway/addressbook/controllers"
	docs "github.com/JJDoneAway/addressbook/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           GO Example Addressbook
// @version         1.0
// @description     This is a simple GO application with some basic REST CRUD operations.

// @Accept json
// @Produce json

// @contact.name   Johannes Höhne
// @contact.email  Johannes.Hoehne1@mail.schwarz

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/"

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html#/")
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	controllers.AddUserRouts(router)

	log.Fatal(router.Run(":8080"))

}
