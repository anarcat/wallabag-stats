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
		XAxis: chart.XAxis{
			Name:           "Time",
			NameStyle:      chart.StyleShow(),
			Style:          chart.Style{Show: true},
			ValueFormatter: chart.TimeHourValueFormatter,
		},
		YAxis: chart.YAxis{
			Name:      "Unread",
			NameStyle: chart.StyleShow(),
			Style:     chart.Style{Show: true},
		},
		YAxisSecondary: chart.YAxis{
			Name:      "Total",
			NameStyle: chart.StyleShow(),
			Style: chart.Style{
				Show: false, //enables / displays the secondary y-axis
			},
		},
		Background: chart.Style{
			Padding: chart.Box{
				Top:  20,
				Left: 20,
			},
		},
		Series: []chart.Series{
			chart.TimeSeries{
				Name: "Unread",
				XValues: []time.Time{
					time.Now().AddDate(0, 0, -10),
					time.Now().AddDate(0, 0, -9),
					time.Now().AddDate(0, 0, -8),
					time.Now().AddDate(0, 0, -7),
					time.Now().AddDate(0, 0, -6),
					time.Now().AddDate(0, 0, -5),
					time.Now().AddDate(0, 0, -4),
					time.Now().AddDate(0, 0, -3),
					time.Now().AddDate(0, 0, -2),
					time.Now().AddDate(0, 0, -1),
					time.Now(),
				},
				YValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0},
			},
			chart.TimeSeries{
				YAxis: chart.YAxisSecondary,
				Name:  "Total",
				XValues: []time.Time{
					time.Now().AddDate(0, 0, -10),
					time.Now().AddDate(0, 0, -9),
					time.Now().AddDate(0, 0, -8),
					time.Now().AddDate(0, 0, -7),
					time.Now().AddDate(0, 0, -6),
					time.Now().AddDate(0, 0, -5),
					time.Now().AddDate(0, 0, -4),
					time.Now().AddDate(0, 0, -3),
					time.Now().AddDate(0, 0, -2),
					time.Now().AddDate(0, 0, -1),
					time.Now(),
				},
				YValues: []float64{3.0, 4.0, 5.0, 7.0, 8.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0},
			},
		},
	}

	graph.Elements = []chart.Renderable{
		chart.Legend(&graph),
	}

	buffer := bytes.NewBuffer([]byte{})
	err = graph.Render(chart.PNG, buffer)
	err = ioutil.WriteFile("chart.png", buffer.Bytes(), 0644)

	log.Printf("time elapsed: %.2fs\n", time.Since(start).Seconds())
}
