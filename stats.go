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

const DATA_JSON = "data.json"
const CONFIG_JSON = "config.json"
const LOCK_FILE = ".lock"

type WallabagStats struct {
	Times   []time.Time
	Total   []float64
	Unread  []float64
	Starred []float64
}

func readCurrentJson(curJson *WallabagStats) {
	if _, err := os.Stat(DATA_JSON); os.IsNotExist(err) {
		// in case file does not exist, we cannot prefill the WallabagStats
		log.Printf("file does not exist %s", DATA_JSON)
		return
	}
	b, err := ioutil.ReadFile(DATA_JSON)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, curJson)
	if err != nil {
		panic(err)
	}
}

func writeNewJson(newWbgStats WallabagStats) {
	b, err := json.Marshal(newWbgStats)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(DATA_JSON, b, 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	// start := time.Now()
	log.SetOutput(os.Stdout)

	// check if lock file exists and exit, so we do not run this process two times
	if _, err := os.Stat(LOCK_FILE); os.IsExist(err) {
		log.Fatalf("lock file exists %s", LOCK_FILE)
		os.Exit(1)
	}

	// create lock file and delete it on exit of main
	err := ioutil.WriteFile(LOCK_FILE, nil, 0644)
	if err != nil {
		panic(err)
	}
	defer os.Remove(LOCK_FILE)

	var wbgStats WallabagStats
	readCurrentJson(&wbgStats)

	var config wallabago.WallabagConfig
	raw, err := ioutil.ReadFile(CONFIG_JSON)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, &config)
	wallabago.Config = config

	total := float64(wallabago.GetNumberOfTotalArticles())
	archived := float64(wallabago.GetNumberOfArchivedArticles())
	starred := float64(wallabago.GetNumberOfStarredArticles())
	unread := float64(total - archived)
	/*log.Printf("total: %v\n", total)
	log.Printf("archived: %v\n", archived)
	log.Printf("unread: %v\n", unread)
	log.Printf("starred: %v\n", starred)
	log.Printf("time: %v\n", time.Now())
	log.Printf("wbgStats: %v\n", wbgStats)*/

	if wbgStats.Total[len(wbgStats.Total)-1] == total && wbgStats.Unread[len(wbgStats.Unread)-1] == unread && wbgStats.Starred[len(wbgStats.Starred)-1] == starred {
	} else {
		// log.Print("appending new values")
		wbgStats.Times = append(wbgStats.Times, time.Now())
		wbgStats.Total = append(wbgStats.Total, total)
		wbgStats.Unread = append(wbgStats.Unread, unread)
		wbgStats.Starred = append(wbgStats.Starred, starred)

		// log.Printf("wbgStats: %v\n", wbgStats)
		writeNewJson(wbgStats)

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
				ValueFormatter: func(v interface{}) string {
					if vf, isFloat := v.(float64); isFloat {
						return fmt.Sprintf("%0.0f", vf)
					}
					return ""
				},
			},
			YAxisSecondary: chart.YAxis{
				Name:      "Total",
				NameStyle: chart.StyleShow(),
				Style: chart.Style{
					Show: true, //enables / displays the secondary y-axis
				},
				ValueFormatter: func(v interface{}) string {
					if vf, isFloat := v.(float64); isFloat {
						return fmt.Sprintf("%0.0f", vf)
					}
					return ""
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
					Name:    "Unread",
					XValues: wbgStats.Times,
					YValues: wbgStats.Unread,
				},
				chart.TimeSeries{
					YAxis:   chart.YAxisSecondary,
					Name:    "Total",
					XValues: wbgStats.Times,
					YValues: wbgStats.Total,
				},
			},
		}

		graph.Elements = []chart.Renderable{
			chart.Legend(&graph),
		}

		buffer := bytes.NewBuffer([]byte{})
		err = graph.Render(chart.PNG, buffer)
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile("chart.png", buffer.Bytes(), 0644)
		if err != nil {
			panic(err)
		}
	}

	// log.Printf("time elapsed: %.2fs\n", time.Since(start).Seconds())
}
