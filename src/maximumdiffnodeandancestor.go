package main

var maxDiff int

func maxAncestorDiff(root *TreeNode) int {
	if root == nil {
		return maxDiff
	}

	min := root.Val
	max := root.Val

	maxDiffHelper(root.Left, min, max)
	maxDiffHelper(root.Right, min, max)

	return maxDiff
}

func maxDiffHelper(node *TreeNode, minTillNode int, maxTillNode int) {
	if node == nil {
		return
	}

	diff := 0
	if node.Val < minTillNode {
		diff = maxTillNode - node.Val
		minTillNode = node.Val
	} else if node.Val > maxTillNode {
		diff = node.Val - minTillNode
		maxTillNode = node.Val
	} else {
		diff = maxTillNode - node.Val
		if diff < node.Val-minTillNode {
			diff = node.Val - minTillNode
		}
	}

	if diff > maxDiff {
		maxDiff = diff
	}

	maxDiffHelper(node.Left, minTillNode, maxTillNode)
	maxDiffHelper(node.Right, minTillNode, maxTillNode)
}
