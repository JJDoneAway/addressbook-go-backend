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

type userController struct {
	userIDPattern *regexp.Regexp
}

func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/([a-zA-Z0-9]+)/?$`),
	}
}

// extract the user out of the request body
func (uc *userController) parseRequest(r *http.Request) (*models.User, error) {
	dec := json.NewDecoder(r.Body)
	var u models.User
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
 * Entry point to the user controller. It implements the handler interface
 * https://pkg.go.dev/net/http#Handler
 *
 * From here we dispatch to the dedicated CRUD methods
 */
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//as we always talk JSON we can set the header right here once‚
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Check GET all and POST path.
	// e.g.: http://localhost:8080/users  and http://localhost:8080/users/
	// but not http://localhost:8080/users/something else
	if test, _ := regexp.MatchString(`^\/users(\/)?$`, r.URL.Path); test {
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

	// URL with potential user id
	urlParts := uc.userIDPattern.FindStringSubmatch(r.URL.Path)
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
			EncodeResponseAsJSON(fmt.Sprintf("The body of the request can not be parsed to an user entity '%s'.", r.Body), w)
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

// @Summary      Add a new user
// @Description  Will add a new user entity to the storage. The new created user will be returned. Don't add the Id to the user parameter
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user body models.User true "The new User without ID"
// @Success      200
// @Failure      400  {string}  string "ID must be zero, Unparsable JSON body"
// @Router       /users [post]
func (uc *userController) doPOST(w http.ResponseWriter, r *http.Request) {
	user, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		EncodeResponseAsJSON(fmt.Sprintf("Could not parse the body of the request. Error was '%v'", err), w)
		return
	}
	err = user.InsertUser()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		EncodeResponseAsJSON(fmt.Sprintf("Could not insert User: %v. Error was '%v'", *user, err), w)
		return
	}
	EncodeResponseAsJSON(user, w)
}

// @Summary      List all users
// @Description  Provide a list of all currently known user
// @Tags         users
// @Produce      json
// @Success      200  {array}  models.User
// @Router       /users [get]
func (uc *userController) doGetAll(w http.ResponseWriter, r *http.Request) {
	EncodeResponseAsJSON((&models.User{}).GetAllUsers(), w)
}

// @Summary      Delete all users
// @Description  Will delete all users. an empty list will be returned
// @Tags         users
// @Produce      json
// @Success      200
// @Router       /users [delete]
func (uc *userController) doDeleteAll(w http.ResponseWriter, r *http.Request) {
	(&models.User{}).DeleteAllUsers()
	EncodeResponseAsJSON((&models.User{}).GetAllUsers(), w)
}

// @Summary      Get one user
// @Description  Get a user with the provided ID
// @Tags         users
// @Produce      json
// @Param        id path integer true "ID of the user"
// @Success      200  {object}  models.User
// @Failure      400  {string}  string "Unknown ID"
// @Router       /users/{id} [get]
func (uc *userController) doGet(ID uint64, w http.ResponseWriter, r *http.Request) {
	user, err := (&models.User{ID: ID}).GetUserByID()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		EncodeResponseAsJSON(fmt.Sprintf("There is no user entity with the id '%d'", ID), w)
		return
	}
	EncodeResponseAsJSON(user, w)
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
func (uc *userController) doPut(ID uint64, user *models.User, w http.ResponseWriter, r *http.Request) {
	if ID != user.ID {
		w.WriteHeader(http.StatusBadRequest)
		EncodeResponseAsJSON(fmt.Sprintf("The id out of the user entity '%d' must be equal to the id of the url path '%d', but wasn't", user.ID, ID), w)
		return
	}

	err := user.UpdateUser()
	if err == models.ErrUnknownID {
		w.WriteHeader(http.StatusBadRequest)
		EncodeResponseAsJSON(fmt.Sprintf("The user with the id '%d' is unknown. Maybe you mean a POST request", user.ID), w)
		return
	}

	EncodeResponseAsJSON(user, w)

}

// @Summary      Delete one user
// @Description  Delete a user with the provided ID
// @Tags         users
// @Produce      json
// @Param        id path integer true "ID of the user"
// @Success      200  {string}  string
// @Failure      400  {string}  string "Unknown ID"
// @Router       /users/{id} [delete]
func (uc *userController) doDelete(ID uint64, w http.ResponseWriter, r *http.Request) {
	err := (&models.User{ID: ID}).DeleteUserByID()
	if err == models.ErrUnknownID {
		w.WriteHeader(http.StatusBadRequest)
		EncodeResponseAsJSON(fmt.Sprintf("The user with the id '%d' is unknown.", ID), w)
		return
	}
	EncodeResponseAsJSON(fmt.Sprintf("The user with the id '%d' is deleted.", ID), w)

}
