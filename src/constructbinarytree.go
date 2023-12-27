package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
	}

	leftSize := 0
	for leftSize < len(inorder) {
		if inorder[leftSize] == preorder[0] {
			break
		}

		leftSize += 1
	}

	arrSize := len(preorder)

	root.Left = buildTree(preorder[1:leftSize+1], inorder[0:leftSize])
	root.Right = buildTree(preorder[leftSize+1:arrSize], inorder[leftSize+1:arrSize])

	return root

}
