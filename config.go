package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/Strubbl/wallabago"
)

func getConfig() (wallabago.WallabagConfig, error) {
	if *debug {
		log.Printf("getConfig: file is %s", *configJSON)
	}
	var config wallabago.WallabagConfig
	raw, err := ioutil.ReadFile(*configJSON)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(raw, &config)
	return config, err
}
