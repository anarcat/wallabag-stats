package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const defaultChartPNG = "chart.png"
const defaultConfigJSON = "config.json"
const defaultDataJSON = "data.json"

var dataOnly = flag.Bool("data-only", false, "collect data only, do not generate any charts")
var debug = flag.Bool("d", false, "get debug output (implies verbose mode)")
var v = flag.Bool("v", false, "print version")
var verbose = flag.Bool("verbose", false, "verbose mode")
var chartPNG = flag.String("chart", defaultChartPNG, "file name to put the chart PNG")
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
	if *verbose && !*debug {
		log.Printf("verbose mode")
	}
	if *debug {
		log.Printf("handleFlags: debug mode")
		// debug implies verbose
		*verbose = true
	}
}
