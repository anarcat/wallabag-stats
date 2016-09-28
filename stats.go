package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/Strubbl/wallabago"
)

func main() {
	start := time.Now()
	log.SetOutput(os.Stdout)

	var config wallabago.WallabagConfig
	raw, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, &config)
	wallabago.Config = config

	total := wallabago.GetNumberOfTotalArticles()
	archived := wallabago.GetNumberOfArchivedArticles()
	starred := wallabago.GetNumberOfStarredArticles()
	log.Printf("total: %v\n", total)
	log.Printf("archived: %v\n", archived)
	log.Printf("unread: %v\n", total-archived)
	log.Printf("starred: %v\n", starred)
	log.Printf("time: %v\n", time.Now())

	log.Printf("time elapsed: %.2fs\n", time.Since(start).Seconds())
}
