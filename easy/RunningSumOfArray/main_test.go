package main

import (
	"slices"
	"testing"
)

var testCases = []struct {
	input    []int
	expected []int
}{
	{[]int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
	{[]int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}},
	{[]int{3, 1, 2, 10, 1}, []int{3, 4, 6, 16, 17}},
}

func TestRunningSum(t *testing.T) {
	for num, test := range testCases {
		got := runningSum(test.input)
		if !slices.Equal(test.expected, got) {
			t.Errorf("Case [%d] expected: %v, got: %v", num, test.expected, got)
		}
	}
}
