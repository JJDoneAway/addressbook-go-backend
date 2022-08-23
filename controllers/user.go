package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/JJDoneAway/addressbook/models"
	"github.com/gin-gonic/gin"
)

func AddUserRouts(router *gin.Engine) {
	router.GET("users", doGetAll)
	router.POST("users", doPOST)
	router.DELETE("users", doDeleteAll)
	router.GET("users/:id", doGet)
	router.PUT("users/:id", doPut)
	router.DELETE("users/:id", doDelete)
}

func doGetAll(c *gin.Context) {
	c.JSON(http.StatusOK, (&models.User{}).GetAllUsers())
}

func doPOST(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := user.InsertUser(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func doDeleteAll(c *gin.Context) {
	(&models.User{}).DeleteAllUsers()
	c.JSON(http.StatusOK, (&models.User{}).GetAllUsers())
}

func doGet(c *gin.Context) {
	ID := getID(c)
	user, err := (&models.User{ID: ID}).GetUserByID()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("For your ID '%s' we got '%v'", c.Param("id"), err.Error())})
		return
	}
	c.JSON(http.StatusOK, user)
}

func doPut(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ID := getID(c)
	if ID != user.ID {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("The id out of the user entity '%d' must be equal to the id of the url path '%d', but wasn't", user.ID, ID)})
		return
	}

	err := user.UpdateUser()
	if err == models.ErrUnknownID {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("The user with the id '%d' is unknown. Maybe you mend a POST call", user.ID)})
		return
	}

	c.JSON(http.StatusOK, user)
}

func doDelete(c *gin.Context) {
	ID := getID(c)
	err := (&models.User{ID: ID}).DeleteUserByID()
	if err == models.ErrUnknownID {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("The user with the id '%d' is unknown.", ID)})
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("The user with the id '%d' is deleted.", ID))

}

////////////////
// Helpers    //
////////////////

// Helper function as I don't know how to do this in gin
func getID(c *gin.Context) uint64 {
	ret, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return 0
	}
	return ret
}
