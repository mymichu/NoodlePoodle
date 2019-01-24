package config

import (
	"log"

	"strconv"

	"github.com/micro/go-config"
)

// LoadConfiguration loads the configuration from file
func LoadConfiguration() {
	config.LoadFile("./temp/config.json")

	// Get address. Set default to localhost as fallback
	address := config.Get("hosts", "database", "address").String("localhost")

	// Get port. Set default to 3000 as fallback
	port := config.Get("hosts", "database", "port").Int(3000)

	log.Println("LOAD " + address + " - " + "PORT " + strconv.Itoa(port))
}

func LoadMqttConfig() HostMqtt {
	var mqttSetting HostMqtt
	config.Get("hosts", "mqtt").Scan(&mqttSetting)
	return mqttSetting
}

func StartWatchMqttChanges() <-chan HostMqtt {
	chanMQTT := make(chan HostMqtt)
	go func() {

		for {
			w, err := config.Watch("hosts", "mqtt")
			if err != nil {
				continue
			}

			// wait for next value
			v, err := w.Next()
			if err != nil {
				continue
			}

			var mqttSettigns HostMqtt

			erro := v.Scan(&mqttSettigns)
			if erro == nil {
				chanMQTT <- mqttSettigns
			}
		}
	}()
	return chanMQTT
}
