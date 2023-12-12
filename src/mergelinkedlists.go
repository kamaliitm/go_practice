package main

// type Node struct {
// 	data int32
// 	next *Node
// }

// func mergeSortedLinkedLists(list1 *Node, list2 *Node) *Node {

// 	head := &Node{}
// 	tail := head

// 	for list1 != nil && list2 != nil {
// 		if list1.data <= list2.data {
// 			tail.next = list1
// 			list1 = list1.next
// 		} else {
// 			tail.next = list2
// 			list2 = list2.next
// 		}
// 		tail = tail.next
// 	}

// 	if list1 != nil {
// 		tail.next = list1
// 	} else if list2 != nil {
// 		tail.next = list2
// 	}

// 	return head.next
// }
