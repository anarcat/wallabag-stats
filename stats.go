package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/Strubbl/wallabago"
	"github.com/wcharczuk/go-chart" //exposes "chart"
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

	// generate chart
	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: []float64{1.0, 2.0, 3.0, 4.0},
				YValues: []float64{1.0, 2.0, 3.0, 4.0},
			},
		},
	}

	buffer := bytes.NewBuffer([]byte{})
	err = graph.Render(chart.PNG, buffer)
	err = ioutil.WriteFile("chart.png", buffer.Bytes(), 0644)

	log.Printf("time elapsed: %.2fs\n", time.Since(start).Seconds())
}
