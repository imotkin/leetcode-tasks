package main

import (
	"slices"
	"testing"
)

var testCases = []struct {
	input    []byte
	expected []byte
}{
	{[]byte{}, []byte{}},
	{[]byte{1}, []byte{1}},
	{[]byte{1, 2, 3}, []byte{3, 2, 1}},
	{[]byte{1, 0, 1, 0}, []byte{0, 1, 0, 1}},
}

func TestReverseString(t *testing.T) {
	for num, test := range testCases {
		got := reverseString(test.input)
		if !slices.Equal(test.expected, got) {
			t.Errorf("Case [%d] expected: %v, got: %v", num, test.expected, got)
		}
	}
}
