package main

import (
	"log"

	"github.com/JJDoneAway/addressbook/controllers"
	"github.com/JJDoneAway/addressbook/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

	middleware.AddDummies()

	addCors(router)

	middleware.RegisterSwagger(router)

	middleware.RegisterPrometheus(router)

	controllers.AddAddressRouts(router)

	log.Fatal(router.Run(":8080"))

}

func addCors(router *gin.Engine) {
	router.Use(cors.Default())
}
