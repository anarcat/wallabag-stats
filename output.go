package main

import (
	"log"
	"time"
)

const pngFileSuffix = ".png"

func getWallabagStatsSubset(wbgStats *WallabagStats, duration time.Duration) WallabagStats {
	var subset WallabagStats
	since := time.Now().Add(duration)
	if *debug {
		log.Printf("getWallabagStatsSubset: since=%v", since)
	}
	var sinceDataSetNumber int
	for i := 1; i < len(wbgStats.Times); i++ {
		if wbgStats.Times[i].After(since) && wbgStats.Times[i-1].Before(since) {
			sinceDataSetNumber = i
			break
		}
	}
	if *debug {
		log.Printf("getWallabagStatsSubset: sinceDataSetNumber=%v", sinceDataSetNumber)
	}
	subset.Times = wbgStats.Times[sinceDataSetNumber:]
	subset.Total = wbgStats.Total[sinceDataSetNumber:]
	subset.Unread = wbgStats.Unread[sinceDataSetNumber:]
	subset.Starred = wbgStats.Starred[sinceDataSetNumber:]
	return subset
}

func generateOutputIfNewData(wbgStats *WallabagStats, total, archived, unread, starred float64) bool {
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
				log.Print("generating output")
			}
			generateOutput(wbgStats)
		} else {
			if *verbose {
				log.Print("not generating charts due to data-only flag")
			}
		}
		return true
	}
	return false
}

func generateOutput(wbgStats *WallabagStats) {
	wbgStatsLastDay := getWallabagStatsSubset(wbgStats, -24*time.Hour)
	wbgStatsLastWeek := getWallabagStatsSubset(wbgStats, -7*24*time.Hour)
	wbgStatsLastMonth := getWallabagStatsSubset(wbgStats, -30*24*time.Hour)
	wbgStatsLastYear := getWallabagStatsSubset(wbgStats, -365*24*time.Hour)
	if *debug {
		log.Printf("generateOutput: data sets in wbgStats=%v and wbgStatsLastDay=%v", len(wbgStats.Times), len(wbgStatsLastDay.Times))
		log.Printf("generateOutput: data sets in wbgStats=%v and wbgStatsLastWeek=%v", len(wbgStats.Times), len(wbgStatsLastWeek.Times))
		log.Printf("generateOutput: data sets in wbgStats=%v and wbgStatsLastMonth=%v", len(wbgStats.Times), len(wbgStatsLastMonth.Times))
		log.Printf("generateOutput: data sets in wbgStats=%v and wbgStatsLastYear=%v", len(wbgStats.Times), len(wbgStatsLastYear.Times))
	}

	generateChartPNG(wbgStats, *chartPNGPrefix+"overall"+pngFileSuffix)
}
