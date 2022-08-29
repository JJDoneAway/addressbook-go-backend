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

// @Summary      List all users
// @Description  Provide a list of all currently known user
// @Tags         users
// @Produce      json
// @Success      200  {array}  models.User
// @Router       /users [get]
func doGetAll(c *gin.Context) {
	c.JSON(http.StatusOK, (&models.User{}).GetAllUsers())
}

// @Summary      Add a new user
// @Description  Will add a new user entity to the storage. The new created user will be returned. Don't add the Id to the user parameter
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user body models.User true "The new User without ID"
// @Success      200
// @Failure      400  {string}  string "ID must be zero, Unparsable JSON body"
// @Router       /users [post]
func doPOST(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "Pimmel": "Du blöder Pimmel"})
		return
	}

	if err := user.InsertUser(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary      Delete all users
// @Description  Will delete all users. an empty list will be returned
// @Tags         users
// @Produce      json
// @Success      200
// @Router       /users [delete]
func doDeleteAll(c *gin.Context) {
	(&models.User{}).DeleteAllUsers()
	c.JSON(http.StatusOK, (&models.User{}).GetAllUsers())
}

// @Summary      Get one user
// @Description  Get a user with the provided ID
// @Tags         users
// @Produce      json
// @Param        id path integer true "ID of the user"
// @Success      200  {object}  models.User
// @Failure      400  {string}  string "Unknown ID"
// @Router       /users/{id} [get]
func doGet(c *gin.Context) {
	ID := getID(c)
	user, err := (&models.User{ID: ID}).GetUserByID()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("For your ID '%s' we got '%v'", c.Param("id"), err.Error())})
		return
	}
	c.JSON(http.StatusOK, user)
}

// @Summary      Update an existing user
// @Description  Will update an existing user which is identified via its ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id path integer true "ID of the user"
// @Param        user body models.User true "The new User without ID"
// @Success      200  {object}  models.User
// @Failure      400  {string}  string "Unknown ID, ID miss match, Unparsable JSON body"
// @Router       /users/{id} [put]
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
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("The user with the id '%d' is unknown. Maybe you mean a POST request", user.ID)})
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary      Delete one user
// @Description  Delete a user with the provided ID
// @Tags         users
// @Produce      json
// @Param        id path integer true "ID of the user"
// @Success      200  {string}  string
// @Failure      400  {string}  string "Unknown ID"
// @Router       /users/{id} [delete]
func doDelete(c *gin.Context) {
	ID := getID(c)
	err := (&models.User{ID: ID}).DeleteUserByID()
	if err == models.ErrUnknownID {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("The user with the id '%d' is unknown.", ID)})
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("The user with the id '%d' is terminated.", ID))

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
