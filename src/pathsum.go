package main

import "math"

func buildTreeFromArray(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}

	nodeQueue := []*TreeNode{}
	root := &TreeNode{Val: arr[0]}
	nodeQueue = append(nodeQueue, root)
	nodeQueueIter := 0

	i := 1
	for i < len(arr) && nodeQueueIter < len(nodeQueue) {
		node := nodeQueue[nodeQueueIter]
		if arr[i] == math.MinInt {
			node.Left = nil
		} else {
			node.Left = &TreeNode{Val: arr[i]}
			nodeQueue = append(nodeQueue, node.Left)
		}

		i += 1
		if i < len(arr) {
			node.Right = &TreeNode{Val: arr[i]}
			nodeQueue = append(nodeQueue, node.Right)
			i += 1
		}

		nodeQueueIter += 1
	}

	return root
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	sumSoFarMap := map[int]int{}
	sumSoFarMap[0] = 1
	return pathSumFinder(root, targetSum, sumSoFarMap, 0)

}

func pathSumFinder(root *TreeNode, targetSum int, sumSoFarMap map[int]int, currSum int) int {
	if root == nil {
		return 0
	}

	res := 0
	currSum += root.Val
	remSum := currSum - targetSum
	if count, ok := sumSoFarMap[remSum]; ok {
		res += count
	}

	if count2, ok2 := sumSoFarMap[currSum]; ok2 {
		sumSoFarMap[currSum] = count2 + 1
	} else {
		sumSoFarMap[currSum] = 1
	}

	res += pathSumFinder(root.Left, targetSum, sumSoFarMap, currSum) +
		pathSumFinder(root.Right, targetSum, sumSoFarMap, currSum)

	// remove the current sum from the hashmap
	// in order not to use it during the parallel
	// processing of the subtree
	sumSoFarMap[currSum] = sumSoFarMap[currSum] - 1

	return res
}
