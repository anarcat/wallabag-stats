package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/wcharczuk/go-chart" //exposes "chart"
)

func generateChartPNG(wbgStats *WallabagStats, pictureName string) error {
	if *debug {
		log.Printf("generateChartPNG: start generating chart in file pictureName=%v", pictureName)
	}
	if *debugDebug {
		log.Printf("generateChartPNG: wbgStats: \n%v", wbgStats)
	}

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
	if *debug {
		log.Println("generateChartPNG: chart created")
	}
	graph.Elements = []chart.Renderable{
		chart.Legend(&graph),
	}
	if *debug {
		log.Println("generateChartPNG: add legend to chart")
	}

	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	if err != nil {
		if *debug {
			log.Println("generateChartPNG: error while rendering graph")
		}
		fmt.Println(err)
		return err
	}
	err = ioutil.WriteFile(pictureName, buffer.Bytes(), 0644)
	if err != nil {
		if *debug {
			log.Println("generateChartPNG: error while writing png")
		}
		fmt.Println(err)
		return err
	}
	return nil
}
