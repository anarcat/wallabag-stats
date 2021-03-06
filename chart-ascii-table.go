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
		log.Printf("generateChartAsciiTable len(wbgStats.Times)=%v", len(wbgStats.Times))
	}
	for i := 0; i < len(wbgStats.Times); i++ {
		if (i+1)%100 == 0 && i > 0 {
			fmt.Fprintf(tw, format, "---", "----", "-----", "------", "-------")
			fmt.Fprintf(tw, format, "No.", "Date", "Total", "Unread", "Starred")
			fmt.Fprintf(tw, format, "---", "----", "-----", "------", "-------")
		}
		// reference time: Mon Jan 2 15:04:05 -0700 MST 2006
		fmt.Fprintf(tw, format, i+1, wbgStats.Times[i].Format("2006-01-02 15:04:05"), wbgStats.Total[i], wbgStats.Unread[i], wbgStats.Starred[i])
	}
	tw.Flush()
	if *debug {
		log.Println("generateASCIITable end")
	}
}
