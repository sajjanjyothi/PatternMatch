package main

import "testing"

var tests = []struct {
	haystack string
	needle   string
	wanted   int
}{
	{"West Bromwich Albion Football Club", "West Bromwich Albion", 1},
	{"West Bromwich Albion Football Club", "West Bromwich Albion FC", 1},
	{"West Bromwich Albion Football Club", "West Bromwich Albion F.C.", 1},
	{"West Bromwich Albion Football Club", "West Bromwich FC", 1},
	{"West Bromwich Albion Football Club", "West Bromwich", 1},
	{"West Bromwich Albion Football Club", "WBA FC", 1},
	{"West Bromwich Albion Football Club", "WB Albion", 1},
	{"West Bromwich Albion Football Club", "West Brom", 1},
	{"West Bromwich Albion Football Club", "West ham", 0},
	{"West Bromwich Albion Football Club", "West", 0},
	{"West Bromwich Albion Football Club", "East Albion", 0},
	{"West Bromwich Albion Football Club", "west bromwich albion", 1},
	{"West Bromwich Albion Football Club", "west brom", 1},
	{"West Bromwich Albion Football Club", "east brom", 0},
	{"West Bromwich Albion Football Club", "wafc", 0},
	{"West Bromwich Albion Football Club", "west bromwich eastham", 0},
	{"West Bromwich Albion Football Club", "", 0},
	{"West Bromwich Albion Football Club", " ", 0},
}

func TestMatchString(t *testing.T) {
	for _, val := range tests {
		obj := StringFinder{haystack: val.haystack, needle: val.needle}
		exp := obj.MatchString()

		if exp != val.wanted {
			t.Errorf("%s in %s expected %d got %d", val.needle, val.haystack, val.wanted, exp)
		}
	}
}
