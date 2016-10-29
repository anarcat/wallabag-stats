package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Strubbl/wallabago"
)

func getConfig() wallabago.WallabagConfig {
	var config wallabago.WallabagConfig
	raw, err := ioutil.ReadFile(*configJSON)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, &config)
	return config
}
