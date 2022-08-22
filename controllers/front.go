package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", uc)
	http.Handle("/users/", uc)
}

func EncodeResponseAsJSON(data any, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
