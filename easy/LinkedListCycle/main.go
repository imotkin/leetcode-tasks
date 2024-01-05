package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	var adresses = make(map[*ListNode]struct{})

	for {
		if head != nil {
			if _, ok := adresses[head]; ok {
				return true
			} else {
				adresses[head] = struct{}{}
			}

			head = head.Next
		} else {
			return false
		}
	}
}
