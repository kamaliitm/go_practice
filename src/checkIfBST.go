package main

// type QueueElement struct {
// 	node       *Tree
// 	upperBound int32
// 	lowerBound int32
// }

// type Queue struct {
// 	data []QueueElement
// }

// func (queue *Queue) dequeue() *QueueElement {
// 	if queue.data == nil || len(queue.data) == 0 {
// 		return nil
// 	}

// 	ret := &queue.data[0]
// 	queue.data = queue.data[1:]

// 	return ret
// }

// func (queue *Queue) enqueue(e QueueElement) {
// 	if queue.data == nil {
// 		queue.data = []QueueElement{}
// 	}

// 	queue.data = append(queue.data, e)
// }

// func checkIfBST(root *Tree) bool {
// 	lowerBound := int32(math.MinInt32)
// 	upperBound := int32(math.MaxInt32)

// 	queue := Queue{}
// 	queue.enqueue(QueueElement{root, lowerBound, upperBound})
// 	for len(queue.data) > 0 {
// 		qe := queue.dequeue()
// 		if qe.node.val < qe.lowerBound || qe.node.val > qe.upperBound {
// 			return false
// 		}

// 		if qe.node.left != nil {
// 			queue.enqueue(QueueElement{qe.node.left, qe.lowerBound, qe.node.val})
// 		}
// 		if qe.node.right != nil {
// 			queue.enqueue(QueueElement{qe.node.right, qe.node.val, qe.upperBound})
// 		}
// 	}

// 	return true
// }
