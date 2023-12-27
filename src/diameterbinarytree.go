package main

var diameter int

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	diameter = 0

	diameterOfBinaryTreeHelper(root)

	return diameter
}

func diameterOfBinaryTreeHelper(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDiameter := diameterOfBinaryTreeHelper(root.Left)
	rightDiameter := diameterOfBinaryTreeHelper(root.Right)

	diameterIncludingRoot := leftDiameter + rightDiameter
	if diameterIncludingRoot > diameter {
		diameter = diameterIncludingRoot
	}

	maxChildDiameter := leftDiameter
	if rightDiameter > maxChildDiameter {
		maxChildDiameter = rightDiameter
	}

	return maxChildDiameter + 1
}
