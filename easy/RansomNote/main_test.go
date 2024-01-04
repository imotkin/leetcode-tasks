package main

import (
	"testing"
)

var testCases = []struct {
	ransomNote string
	magazine   string
	expected   bool
}{
	{"a", "b", false},
	{"aa", "aab", true},
}

func TestCanConstruct(t *testing.T) {
	for num, test := range testCases {
		got := canConstruct(test.ransomNote, test.magazine)
		if test.expected != got {
			t.Errorf("Case [%d] expected: %v, got: %v", num, test.expected, got)
		}
	}
}
