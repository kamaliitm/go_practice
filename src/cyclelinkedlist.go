package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slowIter := head
	fastIter := head

	for slowIter != nil && fastIter != nil && fastIter.Next != nil {
		slowIter = slowIter.Next
		fastIter = fastIter.Next.Next

		if slowIter == fastIter {
			return true
		}
	}

	return false
}
