package main

import (
	"testing"
	"time"
)

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
		{WallabagStats{}, 0, 0, 0, 0, true},
		{WallabagStats{Times: []time.Time{time.Now()}, Total: []float64{0}, Unread: []float64{0}, Starred: []float64{0}}, 1, 1, 0, 0, true},
	}
	for _, test := range tests {
		b := isDataSetNew(&test.w, test.total, test.archived, test.unread, test.starred)
		if b != test.expectedResult {
			t.Errorf("isDataSetNew(): expectedResult %v, got %v", test.expectedResult, b)
		}
	}
}
