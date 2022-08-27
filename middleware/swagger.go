package middleware

import (
	"net/http"

	"github.com/JJDoneAway/addressbook/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func RegisterSwagger(mux *http.ServeMux) {
	//register swagger UI
	docs.SwaggerInfo.BasePath = "/"
	mux.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), //The url pointing to API definition
	))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/index.html", http.StatusMovedPermanently)
	})
}
