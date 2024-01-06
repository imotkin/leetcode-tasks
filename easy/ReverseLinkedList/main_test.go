package main

import (
	"slices"
	"testing"
)

func getNodes() (*ListNode, *ListNode) {
	node1 := ListNode{Val: 1}
	node2 := ListNode{Val: 2}
	node3 := ListNode{Val: 3}
	node1.Next = &node2
	node2.Next = &node3
	node3.Next = nil

	node4 := node3
	node4.Next = &node2
	node5 := node2
	node5.Next = &node1
	node6 := node1
	node6.Next = nil

	return &node1, &node4
}

var nodeDirect, nodeReverse = getNodes()

var testCases = []struct {
	node     *ListNode
	expected *ListNode
}{
	{nodeDirect, nodeReverse},
}

func getValues(node *ListNode) (values []int) {
	for node != nil {
		values = append(values, node.Val)
		node = node.Next
	}

	return
}

func reverse(slice []int) []int {
	var start, end = 0, len(slice) - 1

	for start < end {
		slice[start], slice[end] = slice[end], slice[start]

		start++
		end--
	}

	return slice
}

func TestReverseList(t *testing.T) {
	for num, test := range testCases {
		got := reverseList(test.node)
		if slices.Equal(getValues(test.expected), reverse(getValues(got))) {
			t.Errorf("Case [%d] expected: %p, got: %p", num, test.expected, got)
		}
	}
}
