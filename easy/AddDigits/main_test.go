package main

import (
	"testing"
)

var testCases = []struct {
	num      int
	expected int
}{
	{38, 2},
	{0, 0},
	{156, 3},
}

func TestAddDigits(t *testing.T) {
	for num, test := range testCases {
		got := addDigits(test.num)
		if test.expected != got {
			t.Errorf("Case [%d] expected: %v, got: %v", num, test.expected, got)
		}
	}
}
