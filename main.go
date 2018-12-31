package main

import (
	"log"
	"net/http"
	"strings"

	"local/swagger/configurator/api"
)

func fileServerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareTwo")
		if strings.HasPrefix(r.URL.Path, "/api") {
			log.Println("HELLO")
			next.ServeHTTP(w, r)
		} else {
			log.Println("HELLO 2")
			handler := http.FileServer(http.Dir("./ui"))
			handler = http.StripPrefix("/swaggerui/", handler)
			handler.ServeHTTP(w, r)
		}
	})
}
func main() {
	petstoreAPI, err := api.NewConfigurator()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Serving noodle poodle api on http://127.0.0.1:3000/swaggerui/")
	http.Handle("/", fileServerMiddleware(petstoreAPI))
	http.ListenAndServe(":3000", nil)
}
