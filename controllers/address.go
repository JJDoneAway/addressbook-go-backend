package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/JJDoneAway/addressbook/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddAddressRouts(router *gin.Engine) {
	router.GET("addresses", doGetAll)
	router.POST("addresses", doPOST)
	router.DELETE("addresses", doDeleteAll)
	router.GET("addresses/:id", doGet)
	router.PUT("addresses/:id", doPut)
	router.DELETE("addresses/:id", doDelete)
}

// @Summary      List all addresses
// @Description  Provide a list of all currently known addresses
// @Tags         addresses
// @Produce      json
// @Success      200  {array}  models.Address
// @Router       /addresses [get]
func doGetAll(c *gin.Context) {
	c.JSON(http.StatusOK, (&models.Address{}).GetAllAddresses())
}

// @Summary      Add a new addresses
// @Description  Will add a new addresses entity to the storage. The new created addresses will be returned. Don't add the Id to the addresses parameter
// @Tags         addresses
// @Accept       json
// @Produce      json
// @Param        addresses body models.Address true "The new addresses without ID"
// @Success      200
// @Failure      400  {string}  string "ID must be zero, Unparsable JSON body"
// @Router       /addresses [post]
func doPOST(c *gin.Context) {
	var user models.Address
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary      Delete all addresses
// @Description  Will delete all addresses. an empty list will be returned
// @Tags         addresses
// @Produce      json
// @Success      200
// @Router       /addresses [delete]
func doDeleteAll(c *gin.Context) {
	(&models.Address{}).DeleteAllAddresses()
	c.JSON(http.StatusOK, (&models.Address{}).GetAllAddresses())
}

// @Summary      Get one address
// @Description  Get a address with the provided ID
// @Tags         addresses
// @Produce      json
// @Param        id path integer true "ID of the user"
// @Success      200  {object}  models.Address
// @Failure      400  {string}  string "Unknown ID"
// @Router       /addresses/{id} [get]
func doGet(c *gin.Context) {
	ID := getID(c)
	user := (&models.Address{Model: gorm.Model{ID: ID}}).GetAddressByID()
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unknown ID"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// @Summary      Update an existing address
// @Description  Will update an existing address which is identified via its ID
// @Tags         addresses
// @Accept       json
// @Produce      json
// @Param        id path integer true "ID of the address"
// @Param        user body models.Address true "The new address without ID"
// @Success      200  {object}  models.Address
// @Failure      400  {string}  string "Unknown ID, ID miss match, Unparsable JSON body"
// @Router       /addresses/{id} [put]
func doPut(c *gin.Context) {
	var user models.Address
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ID := getID(c)
	if ID != user.Model.ID {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("The id out of the address entity '%d' must be equal to the id of the url path '%d', but wasn't", user.ID, ID)})
		return
	}

	// err := user.UpdateAddress()
	// if err == models.ErrUnknownID {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("The address with the id '%d' is unknown. Maybe you mean a POST request", user.ID)})
	// 	return
	// }

	c.JSON(http.StatusOK, user)
}

// @Summary      Delete one address
// @Description  Delete a address with the provided ID
// @Tags         addresses
// @Produce      json
// @Param        id path integer true "ID of the address"
// @Success      200  {string}  string
// @Failure      400  {string}  string "Unknown ID"
// @Router       /addresses/{id} [delete]
func doDelete(c *gin.Context) {
	ID := getID(c)

	err := (&models.Address{Model: gorm.Model{ID: ID}}).DeleteAddressByID()
	if err == models.ErrUnknownID {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("The address with the id '%d' is unknown.", ID)})
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("The address with the id '%d' is terminated.", ID))

}

////////////////
// Helpers    //
////////////////

// Helper function as I don't know how to do this in gin
func getID(c *gin.Context) uint {
	ret, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return 0
	}
	return uint(ret)
}
