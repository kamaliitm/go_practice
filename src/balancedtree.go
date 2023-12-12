package main

import "math"

type Tree struct {
	val   int32
	left  *Tree
	right *Tree
}

func checkBalanced(tree *Tree) (bool, int) {
	if tree == nil {
		return true, -1
	}

	leftBalanced, leftHeight := checkBalanced(tree.left)
	if !leftBalanced {
		return false, leftHeight + 1
	}

	rightBalanced, rightHeight := checkBalanced(tree.right)
	if !rightBalanced {
		return false, rightHeight + 1
	}

	heightDiff := math.Abs(float64(leftHeight) - float64(rightHeight))
	treeHeight := int(math.Max(float64(leftHeight), float64(rightHeight))) + 1

	return heightDiff <= 1, treeHeight
}
