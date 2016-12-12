package main

import (
	"fmt"
	"log"
	"time"
)

const outputDirectory = "output"
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
		err := writeNewJSON(wbgStats)
		if err != nil {
			fmt.Println(err)
		}
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

func generateCharts(wbgStats *WallabagStats) {
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

	// generate only if at least two data rows are available
	if len(wbgStats.Times) > 1 {
		generateChartPNG(wbgStats, *chartPNGPrefix+"overall"+pngFileSuffix)
	}
	if len(wbgStatsLastDay.Times) > 1 && len(wbgStatsLastDay.Times) < len(wbgStats.Times) {
		generateChartPNG(&wbgStatsLastDay, *chartPNGPrefix+"day"+pngFileSuffix)
	}
	if len(wbgStatsLastWeek.Times) > 1 && len(wbgStatsLastWeek.Times) < len(wbgStats.Times) {
		generateChartPNG(&wbgStatsLastWeek, *chartPNGPrefix+"week"+pngFileSuffix)
	}
	if len(wbgStatsLastMonth.Times) > 1 && len(wbgStatsLastMonth.Times) < len(wbgStats.Times) {
		generateChartPNG(&wbgStatsLastMonth, *chartPNGPrefix+"month"+pngFileSuffix)
	}
	if len(wbgStatsLastYear.Times) > 1 && len(wbgStatsLastYear.Times) < len(wbgStats.Times) {
		generateChartPNG(&wbgStatsLastYear, *chartPNGPrefix+"year"+pngFileSuffix)
	}
}

func generateOutput(wbgStats *WallabagStats) {
	generateCharts(wbgStats)
	generateHTML(wbgStats)
}
