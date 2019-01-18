package main

import (
	"log"
	"net/http"

	"./config"
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

	writer := &config.WriterSettings{
		FilePath: "./temp/config.json",
	}
	clientKitchen := &config.Client{
		ID:    1,
		Place: "Test",
		URL:   "https://www.google.ch",
	}
	writer.ChangeOrAddClient(clientKitchen)
	config.LoadConfiguration()
	//http.ListenAndServe(":3000", nil)
}
