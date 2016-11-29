package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// WallabagStats is a data set representing the number of total, unread and starred articles in wallabg at a given time
// This data format has been chosen to easily reuse the data for the axis in github.com/wcharczuk/go-chart. Otherwise
// it would have made more sense to save one data set and create an array of data sets
type WallabagStats struct {
	Times   []time.Time
	Total   []float64
	Unread  []float64
	Starred []float64
}

func readCurrentJSON(curJSON *WallabagStats) error {
	if *debug {
		log.Println("readCurrentJSON")
	}
	if _, err := os.Stat(*dataJSON); os.IsNotExist(err) {
		// in case file does not exist, we cannot prefill the WallabagStats
		if *verbose { // not fatal, just start with a new one
			log.Printf("file does not exist %s", *dataJSON)
		}
		return nil
	}
	b, err := ioutil.ReadFile(*dataJSON)
	if err != nil {
		if *debug {
			log.Println("readCurrentJSON: error while ioutil.ReadFile", err)
		}
		fmt.Println(err)
		return err
	}
	err = json.Unmarshal(b, curJSON)
	if err != nil {
		if *debug {
			log.Println("readCurrentJSON: error while json.Unmarshal", err)
		}
		return err
	}
	return nil
}

func writeNewJSON(newWbgStats *WallabagStats) error {
	if *debug {
		log.Println("writeNewJSON")
	}
	b, err := json.Marshal(newWbgStats)
	if err != nil {
		if *debug {
			log.Println("writeNewJSON: error while marshalling data json", err)
		}
		return err
	}
	err = ioutil.WriteFile(*dataJSON, b, 0644)
	if err != nil {
		if *debug {
			log.Println("writeNewJSON: error while writing data json", err)
		}
		return err
	}
	return nil
}
