package main

import (
	"testing"
	"time"
)

func TestIsDataAllEqualLengths(t *testing.T) {
	var tests = []struct {
		timesCount     int
		totalCount     int
		unreadCount    int
		starredCount   int
		expectedResult bool
	}{
		{0, 0, 0, 0, true},
		{10000, 10000, 10000, 10000, true},
		{0, 0, 0, 1, false},
		{0, 0, 1, 0, false},
		{0, 1, 0, 0, false},
		{1, 0, 0, 0, false},
		{1, 1, 0, 0, false},
		{1, 1, 1, 0, false},
		{0, 1, 0, 0, false},
		{0, 1, 1, 0, false},
		{0, 1, 1, 1, false},
		{0, 0, 1, 1, false},
		{1, 0, 1, 1, false},
		{1, 1, 0, 1, false},
	}
	for _, test := range tests {
		b := isDataAllEqualLengths(test.timesCount, test.totalCount, test.unreadCount, test.starredCount)
		if b != test.expectedResult {
			t.Errorf("isDataAllEqualLengths(%v, %v, %v, %v): expectedResult %v, got %v", test.timesCount, test.totalCount, test.unreadCount, test.starredCount, test.expectedResult, b)
		}
	}
}

func TestIsDataSetNew(t *testing.T) {
	var tests = []struct {
		w              WallabagStats
		total          float64
		archived       float64
		unread         float64
		starred        float64
		expectedResult bool
	}{
		{WallabagStats{Times: []time.Time{}, Total: []float64{}, Unread: []float64{}, Starred: []float64{}}, 1, 1, 0, 0, true},
		// simulate fresh new wallabag-stats installation and fresh new wallabag installation
		{WallabagStats{}, 0, 0, 0, 0, true},
		// simulate fresh new wallabag-stats installation
		{WallabagStats{}, 500, 100, 400, 20, true},
		// simulate fresh new wallabag-stats installation
		{WallabagStats{Times: []time.Time{time.Now()}, Total: []float64{0}, Unread: []float64{0}, Starred: []float64{0}}, 1, 1, 0, 0, true},
		// simulate corrupt data set in json file
		{WallabagStats{Times: []time.Time{time.Now()}, Total: []float64{0, 1}, Unread: []float64{0}, Starred: []float64{0}}, 1, 1, 0, 0, false},
		// simulate all API calls fail
		{WallabagStats{Times: []time.Time{time.Now()}, Total: []float64{1}, Unread: []float64{1}, Starred: []float64{1}}, 0, 0, 0, 0, false},
		// simulate deleting one archived item and failing to get starred number from API
		{WallabagStats{Times: []time.Time{time.Now()}, Total: []float64{100}, Unread: []float64{50}, Starred: []float64{10}}, 99, 49, 50, 0, false},
		// simulate deleting one archived starred item
		{WallabagStats{Times: []time.Time{time.Now()}, Total: []float64{100}, Unread: []float64{50}, Starred: []float64{10}}, 99, 49, 50, 9, true},
		// archived count is now zero although before we already had archived some articles
		{WallabagStats{Times: []time.Time{time.Now()}, Total: []float64{100}, Unread: []float64{50}, Starred: []float64{10}}, 110, 0, 110, 9, false},
		// archived count is now equal total count and unread is therefore zero, which is use case "all items read"
		{WallabagStats{Times: []time.Time{time.Now()}, Total: []float64{100}, Unread: []float64{50}, Starred: []float64{10}}, 110, 110, 0, 9, true},
	}
	for _, test := range tests {
		b := isDataSetNew(&test.w, test.total, test.archived, test.unread, test.starred)
		if b != test.expectedResult {
			t.Errorf("isDataSetNew(): expectedResult %v, got %v\ntotal=%v, unread=%v, starred=%v, archived=%v\nWallabagStats=%v", test.expectedResult, b, test.total, test.unread, test.starred, test.archived, test.w)
		}
	}
}
