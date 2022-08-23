package main

import (
	"log"

	"github.com/JJDoneAway/addressbook/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	controllers.AddUserRouts(router)

	log.Fatal(router.Run(":8080"))

}
