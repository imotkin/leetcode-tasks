package main

import (
	"testing"
)

func initNodesWithoutCicle() *ListNode {
	var node1, node2, node3 *ListNode

	node3 = &ListNode{3, nil}
	node2 = &ListNode{2, node3}
	node1 = &ListNode{1, node2}

	return node1
}

func initNodesWithCicle() *ListNode {
	var node1, node2, node3 *ListNode

	node3 = &ListNode{3, nil}
	node2 = &ListNode{2, node3}
	node1 = &ListNode{1, node2}
	node3.Next = node1

	return node1
}

var testCases = []struct {
	node     *ListNode
	expected bool
}{
	{initNodesWithoutCicle(), false},
	{initNodesWithCicle(), true},
}

func TestHasCycle(t *testing.T) {
	for num, test := range testCases {
		got := hasCycle(test.node)
		if test.expected != got {
			t.Errorf("Case [%d] expected: %v, got: %v", num, test.expected, got)
		}
	}
}
