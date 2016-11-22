package main

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

func generateChartAsciiTable(wbgStats WallabagStats) {
	if *debug {
		log.Println("generateChartAsciiTable start")
	}
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "No.", "Date", "Total", "Unread", "Starred")
	fmt.Fprintf(tw, format, "---", "----", "-----", "------", "-------")
	if *debug {
		log.Printf("generateChartAsciiTable len(wbgStats.Times)=%v", len(wbgStats.Times))
	}
	for i := 0; i < len(wbgStats.Times); i++ {
		fmt.Fprintf(tw, format, i+1, wbgStats.Times[i].Format("2006-01-15 15:04:05"), wbgStats.Total[i], wbgStats.Unread[i], wbgStats.Starred[i])
	}
	tw.Flush()
	if *debug {
		log.Println("generateChartAsciiTable end")
	}
}
