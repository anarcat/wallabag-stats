package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func generatePrettyJSON(wbgStats WallabagStats) {
	if *debug {
		log.Println("generatePrettyJSON start")
	}
	j, err := json.MarshalIndent(wbgStats, "", "  ")
	if err != nil {
		if *debug {
			log.Println("generatePrettyJSON: error while marshalling wbgStats json")
		}
		fmt.Println(err)
	}
	fmt.Printf("%s\n", j)

	if *debug {
		log.Println("generatePrettyJSON end")
	}
}
