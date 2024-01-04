package main

import (
	"slices"
	"testing"
)

var testCases = []struct {
	digits   []int
	expected []int
}{
	{[]int{1, 0, 0}, []int{1, 0, 1}},
	{[]int{1, 5, 9}, []int{1, 6, 0}},
	{[]int{9, 9}, []int{1, 0, 0}},
}

func TestPlusOne(t *testing.T) {
	for num, test := range testCases {
		got := plusOne(test.digits)
		if !slices.Equal(test.expected, got) {
			t.Errorf("Case [%d] expected: %v, got: %v", num, test.expected, got)
		}
	}
}
