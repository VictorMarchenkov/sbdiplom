package config

import (
	"io/ioutil"
	"log"
)

func GetConfig() []byte {

	rawConfig, err := ioutil.ReadFile("internal/config/config.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	return rawConfig
}
