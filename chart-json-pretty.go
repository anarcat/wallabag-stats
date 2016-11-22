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

func generatePrettyNewJSON(wbgStats WallabagStats) {
	if *debug {
		log.Println("generatePrettyNewJSON start")
	}
	j, err := json.MarshalIndent(wbgStats, "", "  ")
	if err != nil {
		if *debug {
			log.Println("generatePrettyNewJSON: error while marshalling wbgStats json")
		}
		panic(err)
	}
	fmt.Printf("%s\n", j)

	if *debug {
		log.Println("generatePrettyNewJSON end")
	}
}
