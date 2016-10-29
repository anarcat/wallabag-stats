package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/Strubbl/wallabago"
	"github.com/wcharczuk/go-chart" //exposes "chart"
)

const lockFile = ".lock"

func main() {
	// start := time.Now()
	log.SetOutput(os.Stdout)

	// check if lock file exists and exit, so we do not run this process two times
	if _, err := os.Stat(lockFile); os.IsExist(err) {
		log.Fatalf("lock file exists %s", lockFile)
		os.Exit(1)
	}

	// create lock file and delete it on exit of main
	err := ioutil.WriteFile(lockFile, nil, 0644)
	if err != nil {
		panic(err)
	}
	defer os.Remove(lockFile)

	var wbgStats WallabagStats
	readCurrentJSON(&wbgStats)

	wallabago.Config = getConfig()

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
		writeNewJSON(wbgStats)

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
