package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/JJDoneAway/addressbook/models"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/([a-zA-Z0-9]+)/?$`),
	}
}

// the method parseRequest satisfies the handler interface
func (uc *userController) parseRequest(r *http.Request) (models.User, error) {
	dec := json.NewDecoder(r.Body)
	var u models.User
	err := dec.Decode(&u)
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//as we always talk JSON we can set the header right here once‚
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Check GET all and POST path
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
		}
		return
	}

	// URL with potential user id
	urlParts := uc.userIDPattern.FindStringSubmatch(r.URL.Path)
	if len(urlParts) == 0 || len(urlParts) > 2 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	userId, err := strconv.ParseUint(urlParts[1], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		uc.doGet(userId, w, r)
	case http.MethodPut:
		EncodeResponseAsJSON(urlParts, w)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	//EncodeResponseAsJSON(urlParts, w)

}

func (uc *userController) doPOST(w http.ResponseWriter, r *http.Request) {
	user, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Could not parse User object: %v. Error was %v", r.Body, err)))
		return
	}
	err2 := user.InsertUser()
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	EncodeResponseAsJSON(user, w)
}

func (uc *userController) doGetAll(w http.ResponseWriter, r *http.Request) {
	EncodeResponseAsJSON((&models.User{}).GetAllUsers(), w)
}

func (uc *userController) doDeleteAll(w http.ResponseWriter, r *http.Request) {
	(&models.User{}).DeleteAllUsers()
	EncodeResponseAsJSON((&models.User{}).GetAllUsers(), w)
}

func (uc *userController) doGet(ID uint64, w http.ResponseWriter, r *http.Request) {
	user, err := (&models.User{ID: ID}).GetUserByID()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	EncodeResponseAsJSON(user, w)
}

func (uc *userController) doPut(ID uint64, w http.ResponseWriter, r *http.Request) {

}

func (uc *userController) doDelete(ID uint64, w http.ResponseWriter, r *http.Request) {

}
