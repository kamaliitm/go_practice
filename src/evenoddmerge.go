package main

// Node defined in mergelinkedlists.go
// func evenOddMerge(list *Node) *Node {
// 	if list == nil {
// 		return nil
// 	}

// 	evenHead := list
// 	evenIter := evenHead
// 	var finalList *Node = nil
// 	oddHead := list.next
// 	oddIter := oddHead

// 	for evenIter != nil && oddIter != nil {
// 		evenIter.next = oddIter.next
// 		finalList = evenIter
// 		evenIter = evenIter.next
// 		if evenIter != nil {
// 			oddIter.next = evenIter.next
// 			oddIter = oddIter.next
// 		}
// 	}

// 	if evenIter != nil {
// 		evenIter.next = oddHead
// 	} else {
// 		finalList.next = oddHead
// 	}

// 	return evenHead
// }
