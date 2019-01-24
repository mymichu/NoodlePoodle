package main

import (
	"fmt"
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
	writer.ChangeClient(clientKitchen)
	config.LoadConfiguration()
	initMQTTSetting := config.LoadMqttConfig()
	fmt.Println("---- INIT ----")
	fmt.Println(initMQTTSetting)
	fmt.Println("--- CHANGES ----")
	settingsMQTT := config.StartWatchMqttChanges()
	for n := range settingsMQTT {
		fmt.Println(n)
	}
	//http.ListenAndServe(":3000", nil)
}
