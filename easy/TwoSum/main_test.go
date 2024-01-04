package main

import (
	"slices"
	"testing"
)

var testCases = []struct {
	input    []int
	target   int
	expected []int
}{
	{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	{[]int{3, 2, 4}, 6, []int{1, 2}},
	{[]int{3, 3}, 6, []int{0, 1}},
}

func TestTwoSum(t *testing.T) {
	for num, test := range testCases {
		got := twoSum(test.input, test.target)
		if !slices.Equal(test.expected, got) {
			t.Errorf("Case [%d] expected: %v, got: %v", num, test.expected, got)
		}
	}
}
