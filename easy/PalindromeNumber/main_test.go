package main

import (
	"testing"
)

var testCases = []struct {
	input    int
	expected bool
}{
	{101, true},
	{123, false},
	{10100101, true},
	{1, true},
}

func TestIsPalindrome(t *testing.T) {
	for num, test := range testCases {
		got := isPalindrome(test.input)
		if test.expected != got {
			t.Errorf("Case [%d] expected: %v, got: %v", num, test.expected, got)
		}
	}
}
