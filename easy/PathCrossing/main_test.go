package main

import (
	"testing"
)

var testCases = []struct {
	path     string
	expected bool
}{
	{"NES", false},
	{"NESWW", true},
	{"ENNNNNNNNNNNEEEEEEEEEESSSSSSSSSS", false},
}

func TestIsPathCrossing(t *testing.T) {
	for num, test := range testCases {
		got := isPathCrossing(test.path)
		if test.expected != got {
			t.Errorf("Case [%d] expected: %v, got: %v", num, test.expected, got)
		}
	}
}
