package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

const dataJSON = "data.json"

// WallabagStats is a data set representing the number of total, unread and starred articles in wallabg at a given time
type WallabagStats struct {
	Times   []time.Time
	Total   []float64
	Unread  []float64
	Starred []float64
}

func readCurrentJSON(curJSON *WallabagStats) {
	if _, err := os.Stat(dataJSON); os.IsNotExist(err) {
		// in case file does not exist, we cannot prefill the WallabagStats
		log.Printf("file does not exist %s", dataJSON)
		return
	}
	b, err := ioutil.ReadFile(dataJSON)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, curJSON)
	if err != nil {
		panic(err)
	}
}

func writeNewJSON(newWbgStats WallabagStats) {
	b, err := json.Marshal(newWbgStats)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(dataJSON, b, 0644)
	if err != nil {
		panic(err)
	}
}
