package main

type LevelOrderQueueNode struct {
	node  *TreeNode
	level int
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := []LevelOrderQueueNode{{node: root, level: 0}}
	queueIter := 0
	currLevel := 0
	currLevelValues := []int{}
	levelOrderValues := [][]int{}

	for queueIter < len(queue) {
		queueNode := queue[queueIter]
		if queueNode.level != currLevel {
			currLevel = queueNode.level
			levelOrderValues = append(levelOrderValues, currLevelValues)
			currLevelValues = []int{queueNode.node.Val}
		} else {
			currLevelValues = append(currLevelValues, queueNode.node.Val)
		}

		if queueNode.node.Left != nil {
			queue = append(queue, LevelOrderQueueNode{node: queueNode.node.Left, level: currLevel + 1})
		}

		if queueNode.node.Right != nil {
			queue = append(queue, LevelOrderQueueNode{node: queueNode.node.Right, level: currLevel + 1})
		}

		queueIter += 1
	}

	levelOrderValues = append(levelOrderValues, currLevelValues)

	return levelOrderValues

}
