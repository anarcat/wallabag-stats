package main

import "testing"

func TestReadJson(t *testing.T) {
	var raw []byte
	raw = []byte("{\"WallabagURL\": \"http://localhost\", \"ClientId\": \"555_puf29hbu4bnu2\", \"ClientSecret\": \"f2o9uhf32j8fj23fji2huo\", \"UserName\": \"john\", \"UserPassword\": \"passworddd\"}")
	c, e := readJson(raw)
	expectedWallabagURL := "http://localhost"
	if c.WallabagURL != expectedWallabagURL {
		t.Errorf("readJson: expected %v, got %v", expectedWallabagURL, c.WallabagURL)
	}
	if e != nil {
		t.Errorf("readJson: err!=nil = %v", e.Error())
	}
}
