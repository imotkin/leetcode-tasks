package main

import (
	"slices"
	"testing"
)

type result struct {
	nums []int
	k    int
}

var testCases = []struct {
	input    []int
	val      int
	expected result
}{
	{[]int{3, 2, 2, 3}, 3, result{[]int{2, 2, 0, 0}, 2}},
	{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, result{[]int{1, 3, 4, 0, 0, 0, 0, 0}, 5}},
}

func TestRemoveElement(t *testing.T) {
	for num, test := range testCases {
		nums, k := removeElement(test.input, test.val)
		if !slices.Equal(test.expected.nums, nums) {
			t.Errorf("Case [%d] expected: %v, got: %v", num, test.expected.nums, nums)
		}
		if test.expected.k != k {
			t.Errorf("Case [%d] expected: %v, got: %v", num, test.expected.k, k)
		}
	}
}
