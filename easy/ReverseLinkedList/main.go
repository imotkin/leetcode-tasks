package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var nodes []*ListNode

	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}

	length := len(nodes)

	if length > 0 {
		for i := length - 1; i > 0; i-- {
			nodes[i].Next = nodes[i-1]
		}

		nodes[0].Next = nil
		return nodes[length-1]
	} else {
		return nil
	}
}
