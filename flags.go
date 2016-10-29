package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const defaultChartPNG = "chart.png"
const defaultConfigJSON = "config.json"
const defaultDataJSON = "data.json"

var debug = flag.Bool("d", false, "get debug output (implies verbose mode)")
var v = flag.Bool("v", false, "print version")
var verbose = flag.Bool("-verbose", false, "verbose mode")
var chartPNG = flag.String("-chart", defaultChartPNG, "file name to put the chart PNG")
var configJSON = flag.String("-config", defaultConfigJSON, "file name of config JSON file")
var dataJSON = flag.String("-data", defaultDataJSON, "file name of data JSON file")

func handleFlags() {
	flag.Parse()
	if *v {
		fmt.Printf("version %v\n", version)
		os.Exit(0)
	}
	if *debug {
		log.Printf("debug mode")
	}
	if *verbose {
		log.Printf("verbose mode")
	}
}
