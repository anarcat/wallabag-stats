package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const defaultChartPNG = "html/chart-"
const defaultConfigJSON = "config.json"
const defaultDataJSON = "data.json"

var dataOnly = flag.Bool("data-only", false, "collect data only, do not generate any charts")
var printTable = flag.Bool("print-table", false, "prints all data as ascii table")
var prettyJSON = flag.Bool("pretty-json", false, "prints all data as formatted json")
var forceChart = flag.Bool("force", false, "force regeneration of chart even if no new data is present")
var debug = flag.Bool("d", false, "get debug output (implies verbose mode)")
var debugDebug = flag.Bool("dd", false, "get even more debug output like data (implies debug mode)")
var v = flag.Bool("v", false, "print version")
var verbose = flag.Bool("verbose", false, "verbose mode")
var chartPNGPrefix = flag.String("chart", defaultChartPNG, "file name prefix for the chart PNGs")
var configJSON = flag.String("config", defaultConfigJSON, "file name of config JSON file")
var dataJSON = flag.String("data", defaultDataJSON, "file name of data JSON file")

func handleFlags() {
	flag.Parse()
	if *debug && len(flag.Args()) > 0 {
		log.Printf("handleFlags: non-flag args=%v", strings.Join(flag.Args(), " "))
	}
	// version first, because it directly exits here
	if *v {
		fmt.Printf("version %v\n", version)
		os.Exit(0)
	}
	// test verbose before debug because debug implies verbose
	if *verbose && !*debug && !*debugDebug {
		log.Printf("verbose mode")
	}
	if *debug && !*debugDebug {
		log.Printf("handleFlags: debug mode")
		// debug implies verbose
		*verbose = true
	}
	if *debugDebug {
		log.Printf("handleFlags: debug² mode")
		// debugDebug implies debug
		*debug = true
		// and debug implies verbose
		*verbose = true
	}
}
