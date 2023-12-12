package main

func findKLargestElementsInBST(root *Tree, k int) []int32 {
	res := []int32{}

	updateKLargestElements(root, k, res)

	return res
}

func updateKLargestElements(root *Tree, k int, res []int32) {
	if root != nil && len(res) < k {
		updateKLargestElements(root.right, k, res)
		if len(res) < k {
			res = append(res, root.val)
			updateKLargestElements(root.left, k, res)
		}
	}
}
