package main

import (
	"fmt"
	"net/http"

	"github.com/JJDoneAway/addressbook/controllers"
)

var port = "8080"

func main() {
	fmt.Println("Register front controllers...")
	controllers.RegisterControllers()

	fmt.Println("Start http server on port", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Print(err)
	}

}
