package main

import (
	"log"
	"time"
)

func generateOutput(wbgStats *WallabagStats, total, archived, unread, starred float64) bool {
	if isDataSetNew(wbgStats, total, archived, unread, starred) == true {
		if *verbose {
			log.Println("found new stats data set")
		}
		wbgStats.Times = append(wbgStats.Times, time.Now())
		wbgStats.Total = append(wbgStats.Total, total)
		wbgStats.Unread = append(wbgStats.Unread, unread)
		wbgStats.Starred = append(wbgStats.Starred, starred)

		if *debugDebug {
			log.Printf("main: wbgStats: %v\n", wbgStats)
		}

		if *verbose {
			log.Print("writing data json file")
		}
		writeNewJSON(wbgStats)
		if !*dataOnly {
			if *verbose {
				log.Print("generating chart PNG")
			}
			generateChartPNG(wbgStats, *chartPNG)
		} else {
			if *verbose {
				log.Print("not generating charts due to data-only flag")
			}
		}
		return true
	}
	return false
}
