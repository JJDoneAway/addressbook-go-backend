package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/JJDoneAway/addressbook/controllers"
	"github.com/JJDoneAway/addressbook/middleware"
	"github.com/JJDoneAway/addressbook/models"
	"github.com/gin-gonic/gin"
)

func TestAPIGetAllUsers(t *testing.T) {
	middleware.AddDummies()
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)

	r := gin.Default()
	controllers.AddAddressRouts(r)

	r.ServeHTTP(rec, req)

	res := rec.Result()

	if res.StatusCode != http.StatusOK {
		t.Fail()
	}

	var users []models.Address
	err := json.NewDecoder(res.Body).Decode(&users)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if len(users) < 1 {
		t.Error(res.Body)
		t.Error("Must be more than one")
		t.Fail()
	}

	if user := users[0]; user.ID == 0 {
		t.Error("ID must not be 0")
	}

}

func TestAPIGetUser(t *testing.T) {
	middleware.AddDummies()
	want := (&models.Address{}).GetAllAddresses()[0]

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/users/%d", want.ID), nil)

	r := gin.Default()
	controllers.AddAddressRouts(r)

	r.ServeHTTP(rec, req)

	res := rec.Result()

	if res.StatusCode != http.StatusOK {
		t.Fail()
	}

	var user models.Address
	err := json.NewDecoder(res.Body).Decode(&user)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if *want != user {
		t.Errorf("Want %v but got %v", want, user)
	}

}

func TestAPIGetUserErrors(t *testing.T) {
	middleware.AddDummies()

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/users/%d", -666), nil)

	r := gin.Default()
	controllers.AddAddressRouts(r)

	r.ServeHTTP(rec, req)

	res := rec.Result()

	if res.StatusCode != http.StatusBadRequest {
		t.Fail()
	}
}

func TestAPIPostUser(t *testing.T) {
	middleware.AddDummies()

	want := `{"FirstName": "Hans", "LastName": "Huber"}`

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader([]byte(want)))

	r := gin.Default()
	controllers.AddAddressRouts(r)

	r.ServeHTTP(rec, req)

	res := rec.Result()

	if res.StatusCode != http.StatusOK {
		t.Error(res)
	}

	var user models.Address
	err := json.NewDecoder(res.Body).Decode(&user)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if user.ID == 0 || user.FirstName != "Hans" || user.LastName != "Huber" {
		t.Errorf("Want %v but got %v", want, user)
	}
}
