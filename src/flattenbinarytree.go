package main

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	curr := root
	temp := root.Right

	if root.Left != nil {
		root.Right = root.Left
		root.Left = nil

		for curr.Right != nil {
			curr = curr.Right
		}
	}

	curr.Right = temp

}

func flattenHelper(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	curr := root
	temp := root.Right
	leftFlattened := flattenHelper(root.Left)
	if leftFlattened != nil {
		curr.Left = nil
		curr.Right = leftFlattened
		for curr.Right != nil {
			curr = curr.Right
		}
	}

	rightFlattened := flattenHelper(temp)
	if rightFlattened != nil {
		curr.Right = rightFlattened
	}

	return root
}
