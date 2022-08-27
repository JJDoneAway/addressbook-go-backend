package main

import (
	"log"

	"github.com/JJDoneAway/addressbook/controllers"
	"github.com/JJDoneAway/addressbook/middleware"
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

	middleware.RegisterSwagger(router)

	middleware.RegisterPrometheus(router)

	controllers.AddUserRouts(router)

	log.Fatal(router.Run(":8080"))

}
