package main

import (
	"log"
	"net/http"

	"./restapi"
)

func main() {
	petstoreAPI, err := restapi.NewConfigurator()

	if err != nil {
		log.Fatalln(err)
	}

	httpHandle := restapi.AddSwaggerUIToHandler(petstoreAPI)
	log.Println("Serving noodle poodle api on http://127.0.0.1:3000/swaggerui/")

	http.Handle("/", httpHandle)
	http.ListenAndServe(":3000", nil)
}
