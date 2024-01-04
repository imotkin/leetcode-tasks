package main

import (
	"testing"
)

var testCases = []struct {
	input    []int
	expected bool
}{
	{[]int{1, 2, 2, 1}, true},
	{[]int{1, 2, 3, 4}, false},
	{[]int{1, 1, 1, 1}, true},
	{[]int{4, 3, 2, 1}, false},
}

func TestContainsDuplicate(t *testing.T) {
	for num, test := range testCases {
		got := containsDuplicate(test.input)
		if test.expected != got {
			t.Errorf("Case [%d] expected: %v, got: %v", num, test.expected, got)
		}
	}
}
