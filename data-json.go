package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// WallabagStats is a data set representing the number of total, unread and starred articles in wallabg at a given time
type WallabagStats struct {
	Times   []time.Time
	Total   []float64
	Unread  []float64
	Starred []float64
}

func readCurrentJSON(curJSON *WallabagStats) {
	if *debug {
		log.Println("readCurrentJSON")
	}
	if _, err := os.Stat(*dataJSON); os.IsNotExist(err) {
		// in case file does not exist, we cannot prefill the WallabagStats
		if *verbose { // not fatal, just start with a new one
			log.Printf("file does not exist %s", *dataJSON)
		}
		return
	}
	b, err := ioutil.ReadFile(*dataJSON)
	if err != nil {
		if *debug {
			log.Println("readCurrentJSON: error while ioutil.ReadFile")
		}
		panic(err)
	}
	err = json.Unmarshal(b, curJSON)
	if err != nil {
		if *debug {
			log.Println("readCurrentJSON: error while json.Unmarshal")
		}
		panic(err)
	}
}

func writeNewJSON(newWbgStats WallabagStats) {
	if *debug {
		log.Println("writeNewJSON")
	}
	b, err := json.Marshal(newWbgStats)
	if err != nil {
		if *debug {
			log.Println("writeNewJSON: error while marshalling data json")
		}
		panic(err)
	}
	err = ioutil.WriteFile(*dataJSON, b, 0644)
	if err != nil {
		if *debug {
			log.Println("writeNewJSON: error while writing data json")
		}
		panic(err)
	}
}
