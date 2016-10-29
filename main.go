package main

import (
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/Strubbl/wallabago"
)

const lockFile = ".lock"

func printElapsedTime(start time.Time) {
	if *debug {
		log.Printf("time elapsed: %.2fs\n", time.Since(start).Seconds())
	}
}

func main() {
	start := time.Now()
	defer printElapsedTime(start)

	log.SetOutput(os.Stdout)

	handleFlags()

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

	// comparing last data set with currently fetched data set
	if wbgStats.Total[len(wbgStats.Total)-1] == total && wbgStats.Unread[len(wbgStats.Unread)-1] == unread && wbgStats.Starred[len(wbgStats.Starred)-1] == starred {
		// no data change since last call --> nothing to do
	} else {
		// log.Print("appending new values")
		wbgStats.Times = append(wbgStats.Times, time.Now())
		wbgStats.Total = append(wbgStats.Total, total)
		wbgStats.Unread = append(wbgStats.Unread, unread)
		wbgStats.Starred = append(wbgStats.Starred, starred)

		// log.Printf("wbgStats: %v\n", wbgStats)
		writeNewJSON(wbgStats)
		generateChartPNG(wbgStats)
	}
}
