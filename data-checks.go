package main

import (
	"fmt"
	"log"
)

func isDataAllEqualLengths(times, total, unread, starred int) bool {
	numbers := [4]int{times, total, unread, starred}
	if *debug {
		log.Printf("isDataAllEqualLengths: numbers=%v", numbers)
	}
	for i := 0; i < len(numbers); i++ {
		for m := 0; m < len(numbers); m++ {
			if i == m {
				continue
			}
			if *debugDebug {
				log.Printf("isDataAllEqualLengths: comparing numbers[%v]=%v with numbers[%v]=%v", i, numbers[i], m, numbers[m])
			}
			if numbers[i] != numbers[m] {
				return false
			}
		}
	}
	return true
}

func isDataSetNew(wbgStats *WallabagStats, total, archived, unread, starred float64) bool {
	if !isDataAllEqualLengths(len(wbgStats.Times), len(wbgStats.Total), len(wbgStats.Unread), len(wbgStats.Starred)) {
		fmt.Println("error: data set from JSON file is not valid, because array sizes have different length")
	}
	// comparing last data set with currently fetched data set
	if wbgStats.Total[len(wbgStats.Total)-1] == total && wbgStats.Unread[len(wbgStats.Unread)-1] == unread && wbgStats.Starred[len(wbgStats.Starred)-1] == starred {
		if *verbose {
			log.Println("no data change since last call --> nothing to do")
		}
		return false
		// also comparing each item of last data set for not being 0, but current being 0
	} else if (wbgStats.Total[len(wbgStats.Total)-1] != 0 && total == 0) || (wbgStats.Unread[len(wbgStats.Unread)-1] != 0 && unread == 0) || (wbgStats.Starred[len(wbgStats.Starred)-1] != 0 && starred == 0) {
		if *verbose {
			log.Println("found 0 instead of real value in total, unread or starred, aborting --> nothing to do")
		}
		return false
		// it is unlikely that we have zero archived items when we had more than zero archived items in our call before
	} else if unread == total && archived == 0 && wbgStats.Total[len(wbgStats.Total)-1]-wbgStats.Unread[len(wbgStats.Unread)-1] > 0 {
		if *verbose {
			log.Println("invalid unread count found, aborting --> nothing to do")
		}
		return false
	}
	return true
}
