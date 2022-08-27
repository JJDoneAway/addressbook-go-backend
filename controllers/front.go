package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers(mux *http.ServeMux) {
	uc := newUserController()

	mux.Handle("/users", uc)
	mux.Handle("/users/", uc)
}

func EncodeResponseAsJSON(data any, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
