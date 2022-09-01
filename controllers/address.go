package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/JJDoneAway/addressbook/models"
)

var ErrMissingMandatoryFields = errors.New("firstName and LastName are mandatory")

type addressController struct {
	addressIDPattern *regexp.Regexp
}

func newAddressController() *addressController {
	return &addressController{
		addressIDPattern: regexp.MustCompile(`^/addresses/([a-zA-Z0-9]+)/?$`),
	}
}

// extract the user out of the request body
func (uc *addressController) parseRequest(r *http.Request) (*models.Address, error) {
	dec := json.NewDecoder(r.Body)
	var u models.Address
	err := dec.Decode(&u)
	if err != nil {
		return nil, err
	}
	if u.FirstName == "" || u.LastName == "" {
		return nil, ErrMissingMandatoryFields
	}
	return &u, nil
}

/*
 * Entry point to the user controller.
 */
func (uc addressController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//as we always talk JSON we can set the header right here once‚
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// Check GET all and POST path.
	// e.g.: http://localhost:8080/addresses  and http://localhost:8080/addresses/
	// but not http://localhost:8080/addresses/something else
	if test, _ := regexp.MatchString(`^\/addresses(\/)?$`, r.URL.Path); test {
		switch r.Method {
		case http.MethodGet:
			uc.doGetAll(w, r)
		case http.MethodPost:
			uc.doPOST(w, r)
		case http.MethodDelete:
			uc.doDeleteAll(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			EncodeResponseAsJSON(fmt.Sprintf("The http method '%s' is not implemented on the main level", r.Method), w)
		}
		return
	}

	// URL with potential address id
	urlParts := uc.addressIDPattern.FindStringSubmatch(r.URL.Path)
	if len(urlParts) == 0 || len(urlParts) > 2 {
		w.WriteHeader(http.StatusNotFound)
		EncodeResponseAsJSON(fmt.Sprintf("The url path '%s' is not pointing to any resource", r.URL.Path), w)
		return
	}
	userId, err := strconv.ParseUint(urlParts[1], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		EncodeResponseAsJSON(fmt.Sprintf("The the resource id '%s' is not an integer. But it must be one", urlParts[1]), w)
		return
	}

	switch r.Method {
	case http.MethodGet:
		uc.doGet(userId, w, r)
	case http.MethodPut:
		user, err := uc.parseRequest(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			EncodeResponseAsJSON(fmt.Sprintf("The body of the request can not be parsed to an address entity '%s'.", r.Body), w)
			return
		}
		uc.doPut(userId, user, w, r)
	case http.MethodDelete:
		uc.doDelete(userId, w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		EncodeResponseAsJSON(fmt.Sprintf("The http method '%s' is not implemented on the id level", r.Method), w)
	}

}

// @Summary      Add a new address
// @Description  Will add a new address entity to the storage. The new created user will be returned. Don't add the Id to the user parameter
// @Tags         addresses
// @Accept       json
// @Produce      json
// @Param        user body models.Address true "The new User without ID"
// @Success      200
// @Failure      400  {string}  string "ID must be zero, Unparsable JSON body"
// @Router       /addresses [post]
func (uc *addressController) doPOST(w http.ResponseWriter, r *http.Request) {
	user, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		EncodeResponseAsJSON(fmt.Sprintf("Could not parse the body of the request. Error was '%v'", err), w)
		return
	}
	err = user.InsertAddress()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		EncodeResponseAsJSON(fmt.Sprintf("Could not insert User: %v. Error was '%v'", *user, err), w)
		return
	}
	EncodeResponseAsJSON(user, w)
}

// @Summary      List all addresses
// @Description  Provide a list of all currently known addresses
// @Tags         addresses
// @Produce      json
// @Success      200  {array}  models.Address
// @Router       /addresses [get]
func (uc *addressController) doGetAll(w http.ResponseWriter, r *http.Request) {
	EncodeResponseAsJSON((&models.Address{}).GetAllAddresses(), w)
}

// @Summary      Delete all addresses
// @Description  Will delete all addresses. an empty list will be returned
// @Tags         addresses
// @Produce      json
// @Success      200
// @Router       /addresses [delete]
func (uc *addressController) doDeleteAll(w http.ResponseWriter, r *http.Request) {
	(&models.Address{}).DeleteAllAddress()
	EncodeResponseAsJSON((&models.Address{}).GetAllAddresses(), w)
}

// @Summary      Get one addresses
// @Description  Get a address with the provided ID
// @Tags         addresses
// @Produce      json
// @Param        id path integer true "ID of the user"
// @Success      200  {object}  models.Address
// @Failure      400  {string}  string "Unknown ID"
// @Router       /addresses/{id} [get]
func (uc *addressController) doGet(ID uint64, w http.ResponseWriter, r *http.Request) {
	user, err := (&models.Address{ID: ID}).GetAddressByID()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		EncodeResponseAsJSON(fmt.Sprintf("There is no user entity with the id '%d'", ID), w)
		return
	}
	EncodeResponseAsJSON(user, w)
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
func (uc *addressController) doPut(ID uint64, user *models.Address, w http.ResponseWriter, r *http.Request) {
	if ID != user.ID {
		w.WriteHeader(http.StatusBadRequest)
		EncodeResponseAsJSON(fmt.Sprintf("The id out of the user entity '%d' must be equal to the id of the url path '%d', but wasn't", user.ID, ID), w)
		return
	}

	err := user.UpdateAddress()
	if err == models.ErrUnknownID {
		w.WriteHeader(http.StatusBadRequest)
		EncodeResponseAsJSON(fmt.Sprintf("The user with the id '%d' is unknown. Maybe you mean a POST request", user.ID), w)
		return
	}

	EncodeResponseAsJSON(user, w)

}

// @Summary      Delete one address
// @Description  Delete a address with the provided ID
// @Tags         addresses
// @Produce      json
// @Param        id path integer true "ID of the address"
// @Success      200  {string}  string
// @Failure      400  {string}  string "Unknown ID"
// @Router       /addresses/{id} [delete]
func (uc *addressController) doDelete(ID uint64, w http.ResponseWriter, r *http.Request) {
	err := (&models.Address{ID: ID}).DeleteAddressByID()
	if err == models.ErrUnknownID {
		w.WriteHeader(http.StatusBadRequest)
		EncodeResponseAsJSON(fmt.Sprintf("The user with the id '%d' is unknown.", ID), w)
		return
	}
	EncodeResponseAsJSON(fmt.Sprintf("The user with the id '%d' is deleted.", ID), w)

}
