package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/Strubbl/wallabago"
)

const lockFile = ".lock"

func printElapsedTime(start time.Time) {
	if *debug {
		log.Printf("printElapsedTime: time elapsed %.2fs\n", time.Since(start).Seconds())
	}
}

func removeLockFile(lf string) {
	if *debug {
		log.Printf("removeLockFile: trying to delete %s\n", lf)
	}
	err := os.Remove(lf)
	if err != nil {
		log.Printf("removeLockFile: error while removing lock file %s\n", lf)
		log.Panic(err)
	}
}

func main() {
	start := time.Now()
	defer printElapsedTime(start)

	log.SetOutput(os.Stdout)

	handleFlags()

	// check if lock file exists and exit, so we do not run this process two times
	if _, err := os.Stat(lockFile); os.IsNotExist(err) {
		if *debug {
			log.Printf("main: no lockfile %s present", lockFile)
		}
	} else {
		fmt.Printf("abort: lock file exists %s\n", lockFile)
		os.Exit(1)
	}

	// check for config before writing lock file to use os.Exit in case no config found
	if *verbose {
		log.Println("reading config")
	}
	c, err := getConfig()
	if err == nil {
		if *debug {
			log.Println("main: setting wallabago.Config var")
		}
		wallabago.Config = c
	} else {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// create lock file and delete it on exit of main
	err = ioutil.WriteFile(lockFile, nil, 0644)
	if err != nil {
		if *debug {
			log.Println("main: error while writing lock file")
		}
		panic(err)
	}
	defer removeLockFile(lockFile)

	if *verbose {
		log.Println("reading data json file into memory")
	}
	var wbgStats WallabagStats
	readCurrentJSON(&wbgStats)

	if *verbose {
		log.Println("get current stats data set from Wallabag")
	}
	total := float64(wallabago.GetNumberOfTotalArticles())
	archived := float64(wallabago.GetNumberOfArchivedArticles())
	starred := float64(wallabago.GetNumberOfStarredArticles())
	unread := float64(total - archived)
	if *debug {
		log.Printf("main: total: %v\n", total)
		log.Printf("main: archived: %v\n", archived)
		log.Printf("main: unread: %v\n", unread)
		log.Printf("main: starred: %v\n", starred)
		log.Printf("main: time: %v\n", time.Now())
	}
	if *debugDebug {
		log.Printf("main: wbgStats: %v\n", wbgStats)
	}

	wasChartGenerated := generateOutput(&wbgStats, total, archived, unread, starred)

	if *forceChart && !wasChartGenerated {
		if *verbose {
			log.Print("generating chart PNG because of force flag")
		}
		generateChartPNG(&wbgStats, *chartPNG)
	}
	if *printTable {
		const format = "%v\t%v\t%v\t%v\t%v\t\n"
		generateASCIITable(wbgStats, format)
	}
	if *prettyJSON {
		generatePrettyJSON(wbgStats)
	}
	if *verbose {
		log.Print("main program finish")
	}
}
