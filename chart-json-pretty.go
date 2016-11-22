package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func generatePrettyJSON(wbgStats OldWallabagStats) {
	if *debug {
		log.Println("generatePrettyJSON start")
	}
	j, err := json.MarshalIndent(wbgStats, "", "  ")
	if err != nil {
		if *debug {
			log.Println("generatePrettyJSON: error while marshalling wbgStats json")
		}
		panic(err)
	}
	fmt.Printf("%s\n", j)

	if *debug {
		log.Println("generatePrettyJSON end")
	}
}
