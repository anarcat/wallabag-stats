package main

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

func generateASCIITable(wbgStats WallabagStats, format string) {
	if *debug {
		log.Println("generateASCIITable start")
	}
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "No.", "Date", "Total", "Unread", "Starred")
	fmt.Fprintf(tw, format, "---", "----", "-----", "------", "-------")
	if *debug {
		log.Printf("generateChartAsciiTable len(wbgStats.Data)=%v", len(wbgStats.Data))
	}
	for i := 0; i < len(wbgStats.Data); i++ {
		fmt.Fprintf(tw, format, i+1, wbgStats.Data[i].Times.Format("2006-01-15 15:04:05"), wbgStats.Data[i].Total, wbgStats.Data[i].Unread, wbgStats.Data[i].Starred)
	}
	tw.Flush()
	if *debug {
		log.Println("generateASCIITable end")
	}
}
