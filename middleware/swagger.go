package middleware

import (
	"net/http"

	"github.com/JJDoneAway/addressbook/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func RegisterSwagger() {
	//register swagger UI
	docs.SwaggerInfo.BasePath = "/"
	http.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), //The url pointing to API definition
	))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/index.html", http.StatusMovedPermanently)
	})
}
