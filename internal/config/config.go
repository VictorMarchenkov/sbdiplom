package internal

import (
	"io/ioutil"
	"log"
)

func GetConfig() []byte {
	rawConfig, err := ioutil.ReadFile("../../config.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	return rawConfig
}
