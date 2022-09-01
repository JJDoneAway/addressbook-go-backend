package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers(mux *http.ServeMux) {
	uc := newAddressController()

	mux.Handle("/addresses", uc)
	mux.Handle("/addresses/", uc)
}

func EncodeResponseAsJSON(data any, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
